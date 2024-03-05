package admins

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"

	"github.com/goccy/go-json"
)

// GetAdminRouterMenuList 获取管理路由菜单
func GetAdminRouterMenuList(adminId int) ([]string, []*dtoadmins.MenuInfo) {
	routerList := rds.RedisFindAdminRouter(adminId)
	menuList := make([]*models.AdminMenu, 0)
	if result := database.Db.Model(&models.AdminMenu{}).Where("status = ?", models.AdminMenuStatusActive).Where("type = ?", models.AdminMenuTypeAdminMenu).Order("parent_id ASC").Order("sort ASC").Scan(&menuList); result.Error != nil {
		return nil, nil
	}

	return routerList, MenuChildren(0, menuList, routerList)
}

// MenuChildren 菜单子级
func MenuChildren(menuId int, menuList []*models.AdminMenu, routerList []string) []*dtoadmins.MenuInfo {
	var data []*dtoadmins.MenuInfo
	for _, menu := range menuList {
		if menu.ParentId == menuId {
			if menu.Route == "" || utils.ArrayStringIndexOf(routerList, "*") > -1 || utils.ArrayStringIndexOf(routerList, models.AdminRoutePathname+models.AdminRouteVersion+menu.Route) > -1 {
				childrenList := MenuChildren(menu.Id, menuList, routerList)
				if menu.Route != "" || len(childrenList) > 0 {
					dataTmp := new(dtoadmins.AdminMenuData)
					_ = json.Unmarshal([]byte(menu.Data), &dataTmp)
					data = append(data, &dtoadmins.MenuInfo{
						Id:       menu.Id,
						Name:     menu.Name,
						Route:    menu.Route,
						Children: childrenList,
						Data:     dataTmp,
					})
				}
			}
		}
	}
	return data
}

// AdminMenuIndex 后台菜单列表
func AdminMenuIndex(adminId int, userId int, params *dtoadmins.AdminMenuIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtoadmins.AdminMenuIndexData, 0)}
	//添加并过滤where条件
	filterParams := utils.NewFilterParams().
		InInt("parent_id IN ?", models.FindTableColumnIntn("admin_menu", "id", "name LIKE  ?", "%"+params.ParentName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("route LIKE ?", params.Route+"%").
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	if result := database.Db.Table("admin_menu as am").
		Select("am.id,am.parent_id,am.name,am.route,am.sort,am.status,am.data,am.created_at,am1.name as parentName,au.username as adminName").
		Joins("left join admin_menu am1 on am.parent_id = am1.id").
		Joins("left join admin_user au on am.admin_id = au.id").
		Where("am.status>?", models.AdminMenuStatusDelete).Where("am.type=?", models.AdminMenuTypeAdminMenu).
		Where("am.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	for _, v := range data.Items.([]*dtoadmins.AdminMenuIndexData) {
		v.DataJson = &dtoadmins.AdminMenuData{}
		_ = json.Unmarshal([]byte(v.Data), &v.DataJson)
	}

	return data, nil
}

// AdminMenuUpdate 更新菜单
func AdminMenuUpdate(adminId int, userId int, params *dtoadmins.AdminMenuUpdateParams) (interface{}, error) {
	result := database.Db.Model(&models.AdminMenu{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params)
	return "ok", result.Error
}

// AdminMenuUpdateDesc 更新菜单详情
func AdminMenuUpdateDesc(adminId, userid int, params *dtoadmins.AdminMenuUpdateDescParams) (interface{}, error) {
	dataJsonBytes, _ := json.Marshal(params.DataJson)
	result := database.Db.Model(&models.AdminMenu{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Update("data", string(dataJsonBytes))
	return "ok", result.Error
}

// AdminMenuCreate 新增后台菜单
func AdminMenuCreate(adminId int, userId int, params *dtoadmins.AdminMenuCreateParams) (interface{}, error) {
	result := database.Db.Model(&models.AdminMenu{}).Create(&models.AdminMenu{
		AdminId: models.SuperAdminId,
		Name:    params.Name,
		Type:    models.AdminMenuTypeAdminMenu,
		Route:   params.Route,
	})
	return "ok", result.Error
}

// AdminMenuDelete 后台菜单删除
func AdminMenuDelete(adminId int, userId int, params *dto.DeleteParams) (interface{}, error) {
	result := database.Db.Model(&models.AdminMenu{}).
		Where("id IN ?", params.Ids).Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(&models.AdminMenu{
		Status: models.AdminMenuStatusDelete,
	})
	return "ok", result.Error
}
