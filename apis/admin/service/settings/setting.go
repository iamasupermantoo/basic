package settings

import (
	dtosettings "basic/apis/admin/dto/settings"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"github.com/goccy/go-json"
)

// SettingIndex 设置列表
func SettingIndex(adminId, userId int, params *dtosettings.SettingIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtosettings.SettingIndexData, 0)}
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username like ?", params.AdminName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("field LIKE ?", params.Field+"%").
		Like("value LIKE ?", "%"+params.Value+"%").
		Like("data LIKE ?", "%"+params.Data+"%").
		EqInt("type = ?", params.Type).
		EqInt("group_id = ?", params.GroupId).
		BetweenDate("updated_at BETWEEN ? AND ?", params.UpdatedAt)

	result := database.Db.Table("admin_setting AS s").
		Select("s.id", "s.admin_id", "au.username AS adminName", "s.group_id", "s.name", "s.type", "s.field", "s.value", "s.data", "s.updated_at", "s.created_at").
		Joins("LEFT JOIN admin_user AS au ON s.admin_id = au.id").
		Where("s.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)

	return data, result.Error
}

// SettingCreate 设置创建
func SettingCreate(adminId, userId int, params *dtosettings.SettingCreateParams) (interface{}, error) {
	settingInfo := &models.AdminSetting{
		AdminId: adminId,
		GroupId: params.GroupId,
		Name:    params.Name,
		Type:    params.Type,
		Field:   params.Field,
		Value:   params.Value,
		Data:    params.Data,
	}
	result := database.Db.Create(settingInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	//	更新缓存
	err := rds.DeleteRedisAdminSetting([]int{settingInfo.Id})
	return "ok", err
}

// SettingDelete 设置删除
func SettingDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	result := database.Db.Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Delete(&models.AdminSetting{}, params.Ids)
	if result.Error != nil {
		return nil, result.Error
	}

	//	更新缓存
	err := rds.DeleteRedisAdminSetting(params.Ids)
	return "ok", err
}

// SettingUpdate 设置修改
func SettingUpdate(adminId, userId int, params *dtosettings.SettingUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.AdminSetting{}).Where("id = ?", params.Id).Updates(params); result.Error != nil {
		return nil, result.Error
	}

	//	更新缓存
	err := rds.DeleteRedisAdminSetting([]int{params.Id})
	return "ok", err
}

// SettingUpdateDesc 更新设置详情
func SettingUpdateDesc(adminId, userId int, params *dtosettings.SettingUpdateDescParams) (interface{}, error) {
	value := params.Value

	//	序列化值
	if params.Type == models.AdminSettingTypeCheckbox ||
		params.Type == models.AdminSettingTypeChildren ||
		params.Type == models.AdminSettingTypeJson ||
		params.Type == models.AdminSettingTypeImages {
		valueBytes, _ := json.Marshal(params.Value)
		value = string(valueBytes)
	}

	//	更新值
	if result := database.Db.Model(&models.AdminSetting{}).Where("id = ?", params.Id).Update("value", value); result.Error != nil {
		return nil, result.Error
	}

	//	更新缓存
	err := rds.DeleteRedisAdminSetting([]int{params.Id})
	return "ok", err
}
