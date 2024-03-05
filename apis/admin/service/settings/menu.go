package settings

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"github.com/goccy/go-json"
)

// AdminSettingMenuIndex 菜单列表
func AdminSettingMenuIndex(adminId int, userId int, params *dtoadmins.AdminSettingMenuIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtoadmins.AdminSettingMenuIndexData, 0)}
	//添加并过滤where条件
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", "%"+params.AdminName+"%")).
		InInt("parent_id IN ?", models.FindTableColumnIntn("admin_menu", "id", "name LIKE  ?", "%"+params.ParentName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("route LIKE ?", params.Route+"%").
		EqInt("status = ?", params.Status).
		EqInt("type=?", params.Type).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	if result := database.Db.Table("admin_menu as am").
		Select("am.id", "am.admin_id", "am.parent_id", "am.name", "am.route", "am.sort", "am.type", "am.status", "am.data", "am.created_at", "am1.name as parentName", "au.username as adminName").
		Joins("left join admin_menu am1 on am.parent_id = am1.id").
		Joins("left join admin_user au on am.admin_id = au.id").
		Where("am.status>?", models.AdminMenuStatusDelete).Where("am.type<>?", models.AdminMenuTypeAdminMenu).
		Where("am.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	for _, v := range data.Items.([]*dtoadmins.AdminSettingMenuIndexData) {
		v.DataJson = &dtoadmins.AdminMenuData{}
		_ = json.Unmarshal([]byte(v.Data), &v.DataJson)
	}
	return data, nil
}

// AdminSettingMenuDelete 菜单删除
func AdminSettingMenuDelete(adminId int, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.AdminMenu{}).
		Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.AdminMenuStatusDelete); result.Error != nil {
		return nil, result.Error
	}

	//	删除前台菜单缓存
	err := rds.DeleteHomeMenuList(params.Ids)
	return "ok", err
}

// AdminSettingMenuCreate 菜单新增
func AdminSettingMenuCreate(adminId int, userId int, params *dtoadmins.AdminSettingMenuCreateParams) (interface{}, error) {
	if result := database.Db.Model(&models.AdminMenu{}).Create(&models.AdminMenu{
		AdminId: adminId,
		Name:    params.Name,
		Type:    params.Type,
		Route:   params.Route,
	}); result.Error != nil {
		return nil, result.Error
	}

	//	删除前台菜单缓存
	err := rds.DeleteHomeMenuList([]int{params.Id})
	return "ok", err
}

// AdminSettingMenuUpdate 菜单更新
func AdminSettingMenuUpdate(adminId int, userId int, params *dtoadmins.AdminSettingMenuUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.AdminMenu{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params); result.Error != nil {
		return nil, result.Error
	}

	//	删除前台菜单缓存
	err := rds.DeleteHomeMenuList([]int{params.Id})
	return "ok", err
}

// AdminSettingMenuUpdateDesc 菜单更新详情
func AdminSettingMenuUpdateDesc(adminId int, userId int, params *dtoadmins.AdminSettingMenuUpdateDescParams) (interface{}, error) {
	dataJsonBytes, _ := json.Marshal(params.DataJson)
	if result := database.Db.Model(&models.AdminMenu{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Update("data", string(dataJsonBytes)); result.Error != nil {
		return nil, result.Error
	}

	//	删除前台菜单缓存
	err := rds.DeleteHomeMenuList([]int{params.Id})
	return "ok", err
}
