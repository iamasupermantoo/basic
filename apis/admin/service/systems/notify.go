package systems

import (
	dtosystems "basic/apis/admin/dto/systems"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
)

// NotifyIndex 获取通知列表
func NotifyIndex(adminId, userId int, params *dtosystems.NotifyIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtosystems.NotifyIndexData, 0)}
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		Like("name LIKE ?", params.Name+"%").
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	//	查询数据库
	result := database.Db.Table("notify as n").
		Select("n.id", "n.admin_id", "n.user_id", "n.name", "u.username as username", "a.username as adminName", "n.type", "n.content", "n.data", "n.status", "n.created_at", "n.updated_at").
		Joins("left join admin_user a on n.admin_id = a.id").
		Joins("left join user as u on n.user_id=u.id").
		Where("n.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("n.status > ?", models.SystemNotifyStatusDelete).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)
	return data, result.Error
}

// NotifyCreate 新增通知
func NotifyCreate(adminId, userId int, params *dtosystems.NotifyCreateParams) (interface{}, error) {
	// 获取用户信息
	userInfo := &models.User{}
	if result := database.Db.Where("username = ?", params.UserName).Take(userInfo); result.Error != nil {
		return nil, result.Error
	}

	params.AdminId = adminId
	params.UserId = userInfo.Id
	if result := database.Db.Table("notify").Create(params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// NotifyDelete 删除通知
func NotifyDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.Notify{}).
		Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.SystemNotifyStatusDelete); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// NotifyUpdate 更新通知
func NotifyUpdate(adminId, userId int, params *dtosystems.NotifyUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.Notify{}).
		Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Updates(params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}
