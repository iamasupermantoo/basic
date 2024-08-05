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

// LangIndex 获取语言列表
func LangIndex(adminId, userId int, params *dtosystems.LangIndexParams) (interface{}, error) {
	data := &dto.IndexData{}
	data.Items = make([]*dtosystems.LangIndexData, 0)

	// 过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("alias LIKE ?", params.Alias+"%").
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	// 查询数据库
	result := database.Db.Table("lang as l").
		Select("l.id", "l.admin_id", "l.name", "l.alias", "l.icon", "l.data", "l.sort", "l.status", "l.created_at", "a.username as adminName").
		Joins("left join admin_user as a on l.admin_id=a.id").
		Where("l.status > ?", models.LangStatusDelete).
		Where("l.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)

	return data, result.Error
}

// LangCreate 新增语言配置
func LangCreate(adminId, userId int, params *dtosystems.LangCreateParams) (interface{}, error) {
	nowTime := time.Now()
	params.AdminId = adminId
	params.CreatedAt = int(nowTime.Unix())
	params.UpdatedAt = int(nowTime.Unix())
	if result := database.Db.Table("lang").Create(params); result.Error != nil {
		return nil, result.Error
	}

	//	删除管理语言列表
	err := rds.DeleteLangList([]int{params.Id})
	return "ok", err
}

// LangDelete 删除语言配置
func LangDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.Lang{}).Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.LangStatusDelete); result.Error != nil {
		return nil, result.Error
	}

	//	删除管理语言列表
	err := rds.DeleteLangList(params.Ids)
	return "ok", err
}

// LangUpdate 更新语言配置
func LangUpdate(adminId, userId int, params *dtosystems.LangUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.Lang{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params); result.Error != nil {
		return nil, result.Error
	}

	//	删除管理语言列表
	err := rds.DeleteLangList([]int{params.Id})
	return "ok", err
}
