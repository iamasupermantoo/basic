package users

import (
	dtousers "basic/apis/admin/dto/users"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"time"

	"gorm.io/gorm"
)

// UserIndex 查询用户
func UserIndex(adminId, userId int, params *dtousers.UserIndexParams) (interface{}, error) {
	data := new(dto.IndexData)
	data.Items = make([]*dtousers.UserIndexData, 0)

	// 过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("parent_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.ParentName+"%")).
		Like("username LIKE ?", params.UserName+"%").
		Like("nickname LIKE ?", params.NickName+"%").
		Like("email LIKE ?", params.Email+"%").
		Like("telephone LIKE ?", params.Telephone+"%").
		EqInt("sex = ?", params.Sex).
		EqInt("type = ?", params.Type).
		EqInt("status = ?", params.Status).
		BetweenDate("birthday BETWEEN ? AND ?", params.Birthday).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	//	查询数据
	result := database.Db.Table("user u").
		Select("u.id", "u.parent_id", "u.avatar", "u.username", "u.nickname", "u.money", "u.score", "u.email", "u.telephone", "u.sex", "u.birthday", "u.updated_at", "u.created_at", "u.data", "u.desc", "a.username as adminName", "u.type", "u.status", "u1.username parentName").
		Joins("left join admin_user a on u.admin_id=a.id").
		Joins("left join user u1 on u.parent_id=u1.id").
		Where("u.status > ?", models.UserStatusDelete).
		Where("u.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)
	return data, result.Error
}

// UserUpdate 更新用户
func UserUpdate(adminId, userId int, params *dtousers.UserUpdateParams) (interface{}, error) {
	//	替换生日转为整型
	if params.Birthday != "" {
		birthdayTime, err := time.Parse("2006/01/02", params.Birthday)
		if err != nil {
			return nil, err
		}
		params.BirthdayInt = int(birthdayTime.Unix())
	}
	if result := database.Db.Model(&models.User{}).
		Where("id = ?", params.Id).
		Where("admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Updates(&params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// UserDelete 删除用户
func UserDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.User{}).
		Where("id IN ?", params.Ids).
		Where("admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.UserStatusDelete); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// UserCreate 创建用户
func UserCreate(adminId, userId int, params *dtousers.UserCreateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	获取邀请码信息
	var parentUserId int
	if params.Code != "" {
		inviteInfo, err := rds.RedisFindInviteCodeInfo(rdsConn, params.Code)
		if err != nil {
			return nil, err
		}

		//	是否覆盖上级ID， 管理ID
		parentUserId = inviteInfo.UserId
		adminId = inviteInfo.AdminId
	}

	//	插入用户数据
	if result := database.Db.Create(&models.User{
		ParentId: parentUserId,
		AdminId:  adminId,
		UserName: params.UserName,
		Password: utils.PasswordEncrypt(params.Password),
	}); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// UserBalance 用户余额变动
func UserBalance(adminId, userId int, params *dtousers.UserBalanceParam) (interface{}, error) {
	//	查询用户信息
	userInfo := models.User{}
	if result := database.Db.Model(&userInfo).
		Where("username = ?", params.UserName).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Take(&userInfo); result.Error != nil {
		return nil, result.Error
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		var err error
		switch params.Type {
		//	账户系统充值
		case models.WalletBillTypeSystemDeposit:
			err = users.UserDeposit(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userInfo.AdminId,
				ParentId: userInfo.ParentId,
				UserId:   userInfo.Id,
				BillType: models.WalletBillTypeSystemDeposit,
				Balance:  userInfo.Money,
				Money:    params.Money,
			})
		//	账户系统扣款
		case models.WalletBillTypeSystemWithdraw:
			err = users.UserSpending(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userInfo.AdminId,
				ParentId: userInfo.ParentId,
				UserId:   userInfo.Id,
				BillType: models.WalletBillTypeSystemWithdraw,
				Balance:  userInfo.Money,
				Money:    params.Money,
			})
		}
		return err
	})
	return "ok", err
}

// UserRelationTree 用户关系树
func UserRelationTree(adminId, userId int) (interface{}, error) {
	userTree := GetUserTree(adminId, 0)
	return userTree, nil
}

// GetUserTree 获取用户关系树
func GetUserTree(adminId, parentId int) []*dtousers.UserRelationTree {
	adminIds := rds.RedisFindAdminChildrenIds(adminId)

	//	获取当前层级用户数组
	userTree := make([]*dtousers.UserRelationTree, 0)
	if result := database.Db.Table("user u").
		Select("u.id", "u.avatar", "u.username", "u.nickname", "u.email", "u.telephone", "u.updated_at", "u.created_at", "COUNT(u.id) sumPeople", "IFNULL(SUM(d.money), 0) sumDeposit", "IFNULL(SUM(w.money), 0) sumWithdraw", "IFNULL(SUM(b.money), 0) sumEarnings").
		Joins("LEFT JOIN wallet_order d on u.id=d.user_id AND d.type = ? AND d.status = ?", models.WalletOrderTypeDeposit, models.WalletOrderStatusComplete).
		Joins("LEFT JOIN wallet_order w on u.id=w.user_id AND w.type = ? AND w.status = ?", models.WalletOrderTypeWithdraw, models.WalletOrderStatusComplete).
		Joins("LEFT JOIN wallet_bill b on u.id=b.user_id AND b.type = ?", models.WalletBillTypeEarnings).
		Where("u.admin_id IN ?", adminIds).
		Where("u.status = ?", models.UserStatusActive).
		Where("u.parent_id = ?", parentId).
		Group("u.id").
		Order("sumPeople DESC").
		Limit(50).
		Scan(&userTree); result.Error != nil {
		return nil
	}

	//	递归查询
	for _, v := range userTree {
		if parentId == 0 {
			v.Header = "root"
		} else {
			v.Header = "generic"
		}

		v.Children = make([]*dtousers.UserRelationTree, 0)
		v.Children = GetUserTree(adminId, v.Id)
	}

	return userTree
}
