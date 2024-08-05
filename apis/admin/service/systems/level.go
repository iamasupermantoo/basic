package systems

import (
	dtosystems "basic/apis/admin/dto/systems"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"time"
)

// LevelIndex 获取等级列表
func LevelIndex(adminId, userId int, params *dtosystems.LevelIndexParams) (interface{}, error) {
	data := &dto.IndexData{}
	data.Items = make([]*dtosystems.LevelIndexData, 0)

	// 过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		Like("name LIKE ?", params.Name+"%").
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	// 查询数据库
	result := database.Db.Table("level l").
		Select("l.id", "l.admin_id", "l.level", "l.name", "l.money", "l.icon", "l.days", "l.data", "l.desc", "l.status", "l.updated_at", "l.created_at", "a.username as adminName").
		Joins("left join admin_user as a on l.admin_id=a.id").
		Where("l.status > ?", models.LevelStatusDelete).
		Where("l.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)

	return data, result.Error
}

// LevelCreate 新增等级配置
func LevelCreate(adminId, userId int, params *dtosystems.LevelCreateParams) (interface{}, error) {
	nowTime := time.Now()
	params.AdminId = adminId
	params.CreatedAt = int(nowTime.Unix())
	params.UpdatedAt = int(nowTime.Unix())
	if result := database.Db.Table("level").Create(params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// LevelDelete 删除等级列表
func LevelDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.Level{}).Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.LevelStatusDelete); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// LevelUpdate 更新等级列表
func LevelUpdate(adminId, userId int, params *dtosystems.LevelUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.Level{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}
