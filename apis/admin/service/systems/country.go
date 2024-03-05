package systems

import (
	dtosystems "basic/apis/admin/dto/systems"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	_ "basic/utils"
	"time"
)

// CountryIndex 获取国家列表
func CountryIndex(adminId, userId int, params *dtosystems.CountryIndexParams) (interface{}, error) {
	data := &dto.IndexData{}
	data.Items = make([]*dtosystems.CountryIndexData, 0)

	//	过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("alias LIKE ?", params.Alias+"%").
		Like("code = ?", params.Code).
		Like("iso1 = ?", params.Iso1).
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	//	查询数据库
	result := database.Db.Table("country c").
		Select("c.id", "c.admin_id", "c.name", "a.username as adminName", "c.alias", "c.icon", "c.iso1", "c.sort", "c.code", "c.status", "c.data", "c.created_at").
		Joins("left join admin_user as a on c.admin_id=a.id").
		Where("c.status > ?", models.CountryStatusDelete).
		Where("c.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)

	return data, result.Error
}

// CountryCreate 新增国家配置
func CountryCreate(adminId, userId int, params *dtosystems.CountryCreateParams) (interface{}, error) {
	nowTime := time.Now()
	params.AdminId = adminId
	params.CreatedAt = int(nowTime.Unix())
	params.UpdatedAt = int(nowTime.Unix())

	if result := database.Db.Table("country").Create(params); result.Error != nil {
		return nil, result.Error
	}

	//	删除国家列表缓存
	err := rds.DeleteCountryList([]int{params.Id})
	return "ok", err
}

// CountryDelete 删除国家配置
func CountryDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	result := database.Db.Model(&models.Country{}).Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Updates(map[string]interface{}{"status": models.CountryStatusDelete, "updated_at": time.Now().Unix()})
	if result.Error != nil {
		return nil, result.Error
	}

	//	删除国家列表缓存
	err := rds.DeleteCountryList(params.Ids)
	return "ok", err
}

// CountryUpdate 更新国家配置
func CountryUpdate(adminId, userId int, params *dtosystems.CountryUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.Country{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(&params); result.Error != nil {
		return nil, result.Error
	}

	//	删除国家列表缓存
	err := rds.DeleteCountryList([]int{params.Id})
	return "ok", err
}
