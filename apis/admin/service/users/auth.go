package users

import (
	dtousers "basic/apis/admin/dto/users"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"time"
)

// AuthIndex 获取认证列表
func AuthIndex(adminId, userId int, params *dtousers.AuthIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtousers.AuthIndexData, 0)}

	//	过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		Like("real_name LIKE ?", params.RealName+"%").
		Like("address LIKE ?", params.Address+"%").
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	// 查询数据库
	result := database.Db.Table("real_name_auth as r").
		Select("r.id", "r.admin_id", "r.user_id", "r.real_name", "r.number", "r.photo1", "r.photo2", "r.photo3", "r.telephone", "r.address", "r.type", "r.status", "r.data", "r.updated_at", "r.created_at", "a.username as adminName", "u.username as username").
		Joins("left join admin_user as a on r.admin_id=a.id").
		Joins("left join user as u on r.user_id=u.id").
		Where("r.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("r.status > ?", models.UserRealNameStatusDelete).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)
	return data, result.Error
}

// AuthCreate 新增认证
func AuthCreate(adminId, userId int, params *dtousers.AuthCreateParams) (interface{}, error) {
	//	用户信息
	userInfo := &models.User{}
	if result := database.Db.Model(userInfo).
		Where("username = ?", params.UserName).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("status = ?", models.UserStatusActive).Take(userInfo); result.Error != nil {
		return nil, result.Error
	}

	nowTime := time.Now()
	currentAuthInfo := &models.RealNameAuth{}
	params.Status = models.UserRealNameStatusActive
	params.UpdatedAt = int(nowTime.Unix())
	params.AdminId = userInfo.AdminId
	params.UserId = userInfo.Id
	result := database.Db.Model(currentAuthInfo).Where("user_id = ?", userInfo.Id).Find(currentAuthInfo)
	if result.RowsAffected == 0 {
		params.CreatedAt = int(nowTime.Unix())
		params.Telephone = ""
		result = database.Db.Table("real_name_auth").Create(params)
	} else {
		result = database.Db.Model(&models.RealNameAuth{}).Where("id = ?", currentAuthInfo.Id).Updates(params)
	}
	return "ok", result.Error
}

// AuthDelete 删除认证
func AuthDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	result := database.Db.Model(&models.RealNameAuth{}).
		Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.UserRealNameStatusDelete)
	return "ok", result.Error
}

// AuthUpdate 更新认证
func AuthUpdate(adminId, userId int, params *dtousers.AuthUpdateParams) (interface{}, error) {
	result := database.Db.Model(&models.RealNameAuth{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params)
	return "ok", result.Error
}
