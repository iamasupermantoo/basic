package admins

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
	"strings"
	"time"

	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

// ManageIndex 管理列表
func ManageIndex(adminId int, userId int, params *dtoadmins.ManageIndexParams) (interface{}, error) {
	data := &dto.IndexData{}
	data.Items = make([]*dtoadmins.AdminIndexData, 0)

	//添加过滤where条件
	filterParams := utils.NewFilterParams().
		InInt("parent_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE  ?", "%"+params.ParentName+"%")).
		Like("username LIKE ?", params.Username+"%").
		Like("nickname LIKE ?", params.Nickname+"%").
		Like("email LIKE ?", params.Email+"%").
		Like("domains LIKE ?", "%"+params.Domains+"%").
		InInt("id IN ?", models.FindTableColumnIntn("auth_assignment", "admin_id", "name LIKE ?", "%"+params.Role+"%")).
		EqInt("status = ?", params.Status).
		BetweenDate("expired_at BETWEEN ? AND ?", params.ExpiredAt)

	database.Db.Table("admin_user as au").
		Select("au.id", "au.parent_id", "au.username", "au.nickname", "au.email", "au.avatar", "au.money", "au.status", "au.domains", "au.data", "au.expired_at", "au.updated_at", "au.created_at", "au1.username as parentName").
		Joins("left join admin_user au1 on au.parent_id = au1.id").
		Order("au.parent_id ASC").
		Where("au.status > ?", models.AdminUserStatusDelete).
		Where("au.id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("au.id <> ?", adminId).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)

	roleListOptions := rds.GetRolesList(adminId)
	for _, v := range data.Items.([]*dtoadmins.AdminIndexData) {
		rolesList := rds.GetAdminRoles(v.Id)
		v.RolesStr = strings.Join(rolesList, " | ")

		v.Roles = map[string]bool{}
		for _, role := range roleListOptions {
			v.Roles[role.Value.(string)] = utils.ArrayStringIndexOf(rolesList, role.Value.(string)) > -1
		}

		//	配置管理详情
		if v.ParentId == models.SuperAdminId {
			v.DataJson = &dtoadmins.AdminInfoData{}
			_ = json.Unmarshal([]byte(v.Data), &v.DataJson)
		}
	}

	return data, nil
}

// ManageUpdate 管理员更新
func ManageUpdate(adminId int, userId int, params *dtoadmins.ManageUpdateParams) (interface{}, error) {
	nowTime := time.Now()

	//	更新过期时间
	if params.ExpiredTime != "" {
		expireTime, err := time.Parse("2006/01/02", params.ExpiredTime)
		if err != nil {
			return nil, err
		}
		params.ExpiredAt = int(expireTime.Unix())
	}

	//	更新管理密码
	if params.Password != "" {
		params.Password = utils.PasswordEncrypt(params.Password)
	}

	//	更新安全密钥
	if params.SecurityKey != "" {
		params.SecurityKey = utils.PasswordEncrypt(params.SecurityKey)
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		if result := database.Db.Model(&models.AdminUser{}).Where("id = ?", params.Id).
			Where("id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params); result.Error != nil {
			return result.Error
		}

		if len(params.Roles) > 0 {
			if result := tx.Where("admin_id=?", params.Id).Delete(&models.AuthAssignment{}); result.Error != nil {
				return result.Error
			}

			//更新角色信息
			for k, v := range params.Roles {
				if v {
					if result := tx.Create(&models.AuthAssignment{
						Name:      k,
						AdminId:   params.Id,
						CreatedAt: int(nowTime.Unix()),
					}); result.Error != nil {
						return result.Error
					}
				}
			}
		}
		return nil
	})

	// 清空所有域名对应的管理ID
	if params.Domains != "" {
		rdsConn := cache.Rds.Get()
		defer rdsConn.Close()
		rds.DelRedisDomainToAdminId(rdsConn)
	}
	return "ok", err
}

// ManageUpdateDesc 更新商户配置
func ManageUpdateDesc(adminId int, userId int, params *dtoadmins.AdminUpdateDesc) (interface{}, error) {
	dataJsonBytes, _ := json.Marshal(params.DataJson)
	if result := database.Db.Model(&models.AdminUser{}).Where("id = ?", params.Id).
		Where("id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Update("data", string(dataJsonBytes)); result.Error != nil {
		return nil, result.Error
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	rds.RedisDelAdminDataInfo(rdsConn, params.Id)
	return "ok", nil
}

// ManageCreate 新增管理
func ManageCreate(adminId int, userId int, params *dtoadmins.ManageCreateParams) (interface{}, error) {
	nowTime := time.Now()
	// 当前管理信息
	currentAdminInfo := &models.AdminUser{}
	if result := database.Db.Model(currentAdminInfo).Where("id = ?", adminId).Take(currentAdminInfo); result.Error != nil {
		return nil, result.Error
	}

	//	如果是商户, 那么根据配置创建代理个数
	if currentAdminInfo.ParentId == models.SuperAdminId {
		currentAdminData := &dtoadmins.AdminInfoData{}
		_ = json.Unmarshal([]byte(currentAdminInfo.Data), &currentAdminData)

		var currentAdminChildrenNums int64
		if result := database.Db.Model(&models.AdminUser{}).Where("parent_id = ?", currentAdminInfo.Id).Where("status > ?", models.AdminUserStatusDelete).Count(&currentAdminChildrenNums); result.Error != nil {
			return nil, result.Error
		}
		if currentAdminChildrenNums > int64(currentAdminData.Nums) {
			return nil, errors.New("超过限制代理数量,需要升级服务器～")
		}
	}

	//	过期时间, 如果没有设置, 那么无限时间
	expiredAt := nowTime.AddDate(10, 0, 0).Unix()
	if params.ExpiredAt != "" {
		expireTime, err := time.Parse("2006/01/02", params.ExpiredAt)
		if err != nil {
			return nil, err
		}
		expiredAt = expireTime.Unix()
	}

	//	事务开始
	data := ""
	if adminId == models.SuperAdminId {
		adminDataInfo := &dtoadmins.AdminInfoData{Nums: 5, Template: "default"}
		adminDataInfoBytes, _ := json.Marshal(adminDataInfo)
		data = string(adminDataInfoBytes)
	}

	adminInfo := &models.AdminUser{
		ParentId:  adminId,
		UserName:  params.UserName,
		Password:  utils.PasswordEncrypt(params.Password),
		Domains:   params.Domains,
		Data:      data,
		ExpiredAt: int(expiredAt),
		CreatedAt: int(nowTime.Unix()),
	}
	if err := database.Db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Create(adminInfo); result.Error != nil {
			return result.Error
		}

		//添加选中角色
		for k, v := range params.Roles {
			if v {
				result := tx.Create(&models.AuthAssignment{
					Name:      k,
					AdminId:   adminInfo.Id,
					CreatedAt: int(nowTime.Unix()),
				})
				if result.Error != nil {
					return result.Error
				}
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	如果是添加商户, 那么初始化数据
	if adminId == models.SuperAdminId {
		if _, err := MerchantReset(&dtoadmins.MerchantParams{Id: adminInfo.Id}); err != nil {
			return nil, err
		}
	}

	//	删除下级Ids 缓存
	rds.RedisDelAminAllChildren(rdsConn)
	return "ok", nil
}

// ManageDelete 管理删除
func ManageDelete(adminId int, userId int, params *dto.DeleteParams) (interface{}, error) {
	for _, v := range params.Ids {
		if result := database.Db.Model(models.AdminUser{}).
			Where("id IN ?", rds.RedisFindAdminChildrenIds(v)).
			Where("id <> ?", adminId).
			Where("id <> ?", models.SuperAdminId).
			Update("status", models.AdminUserStatusDelete); result.Error != nil {
			return nil, result.Error
		}
	}
	return "ok", nil
}
