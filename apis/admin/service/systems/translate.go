package systems

import (
	dtosystems "basic/apis/admin/dto/systems"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/module/translate"
	"basic/utils"
)

// TranslateIndex 获取字典列表
func TranslateIndex(adminId, userId int, params *dtosystems.TranslateIndexParams) (interface{}, error) {
	data := new(dto.IndexData)
	data.Items = make([]*dtosystems.TranslateIndexData, 0)

	// 过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		Like("name LIKE ?", params.Name+"%").
		EqInt("type = ?", params.Type).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	// 查询数据库
	result := database.Db.Table("translate t").
		Select("t.id", "t.admin_id", "t.type", "t.name", "t.field", "t.value", "t.lang", "t.created_at", "a.username as adminName").
		Joins("left join admin_user as a on t.admin_id=a.id").
		Where("t.admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)

	return data, result.Error
}

// TranslateCreate 新增字典配置
func TranslateCreate(adminId, userId int, params *dtosystems.TranslateCreateParams) (interface{}, error) {
	params.AdminId = adminId
	//	判断是否插入失败
	if result := database.Db.Table("translate").Create(&params); result.Error != nil {
		return nil, result.Error
	}

	//	更新缓存配置
	err := rds.DeleteRedisTranslate([]int{params.Id})
	return "ok", err
}

// TranslateDelete 删除字典配置
func TranslateDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Delete(&models.Translate{}); result.Error != nil {
		return nil, result.Error
	}

	//	更新缓存配置
	err := rds.DeleteRedisTranslate(params.Ids)
	return "ok", err
}

// TranslateUpdate 更新字典配置
func TranslateUpdate(adminId, userId int, params *dtosystems.TranslateUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.Translate{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params); result.Error != nil {
		return nil, result.Error
	}

	//	更新缓存配置
	err := rds.DeleteRedisTranslate([]int{params.Id})
	return "ok", err
}

// TranslateLangField 翻译语言字段数据
func TranslateLangField(adminId, userId int, params *dtosystems.TranslateFiedParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	adminId = rds.RedisFindAdminIdToSettingAdminId(rdsConn, params.AdminId)

	data := make([]*dtosystems.TranslateFieldData, 0)
	result := database.Db.Table("lang as l").
		Select("l.id", "l.name", "l.alias", "l.icon", "IFNULL(t.id, 0) as translateId", "IFNULL(t.value, '') as value").
		Joins("left join translate as t on l.alias=t.lang and l.admin_id=t.admin_id and t.field=?", params.Field).
		Where("l.status=?", models.LangStatusActive).
		Where("l.admin_id = ?", adminId).
		Find(&data)
	return data, result.Error
}

// TranslateUpdateFields 更新字段语言
func TranslateUpdateFields(adminId, userId int, params *dtosystems.TranslateUpdateField) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	adminId = rds.RedisFindAdminIdToSettingAdminId(rdsConn, params.AdminId)
	currentTranslateInfo := &models.Translate{}
	if result := database.Db.Model(currentTranslateInfo).
		Where("admin_id = ?", adminId).
		Where("field = ?", params.Field).
		Where("lang = ?", "zh-CN").
		Take(currentTranslateInfo); result.Error != nil {
		return nil, result.Error
	}

	for _, v := range params.Data {
		translateInfo := &models.Translate{}

		//	如果id = 0 那么新增, 否则修改
		if v.TranslateId > 0 {
			if result := database.Db.Model(translateInfo).Where("id = ?", v.TranslateId).Update("value", v.Value); result.Error != nil {
				return nil, result.Error
			}
		} else {
			if result := database.Db.Create(&models.Translate{
				AdminId: adminId,
				Lang:    v.Alias,
				Name:    currentTranslateInfo.Name,
				Type:    models.TranslateTypeHome,
				Field:   currentTranslateInfo.Field,
				Value:   v.Value,
			}); result.Error != nil {
				return nil, result.Error
			}
		}

	}

	//	删除语言缓存
	rds.RedisDelAllTranslateAll(rdsConn, adminId)
	rds.RedisDelAllHomeTranslateLang(rdsConn, adminId)
	return "ok", nil
}

// TranslateGoogleField google 翻译
func TranslateGoogleField(adminId, userId int, params *dtosystems.TranslateGoogleParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	adminId = rds.RedisFindAdminIdToSettingAdminId(rdsConn, params.AdminId)
	fieldText := rds.RedisFindTranslateField(rdsConn, adminId, "zh-CN", params.Field)

	return translate.Translate(fieldText, "zh-CN", params.Alias)
}
