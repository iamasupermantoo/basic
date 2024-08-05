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

// InviteIndex 获取邀请列表
func InviteIndex(adminId, userId int, params *dtousers.InviteIndexParams) (interface{}, error) {
	data := &dto.IndexData{}
	data.Items = make([]*dtousers.InviteIndexData, 0)

	//	过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		Like("code LIKE ?", params.Code+"%").
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	// 查询数据库
	result := database.Db.Table("invite as i").
		Select("i.id", "i.admin_id", "i.user_id", "i.code", "i.status", "i.data", "i.updated_at", "i.created_at", "a.username as adminName", "u.username as username").
		Joins("left join admin_user as a on i.admin_id=a.id").
		Joins("left join user as u on i.user_id=u.id").
		Where("i.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("i.status > ?", models.UserInviteStatusDelete).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)
	return data, result.Error
}

// InviteDelete 删除邀请
func InviteDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.Invite{}).
		Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Updates(map[string]interface{}{"status": models.UserInviteStatusDelete, "updated_at": time.Now().Unix()}); result.Error != nil {
		return nil, result.Error
	}

	//	删除缓存
	err := rds.DeleteRedisInvite(params.Ids)
	return "ok", err
}

// InviteUpdate 更新邀请
func InviteUpdate(adminId, userId int, params *dtousers.InviteUpdateParams) (interface{}, error) {
	params.UpdatedAt = int(time.Now().Unix())
	result := database.Db.Model(&models.Invite{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params)
	if result.Error != nil {
		return nil, result.Error
	}

	//	删除缓存
	err := rds.DeleteRedisInvite([]int{params.Id})
	return "ok", err
}
