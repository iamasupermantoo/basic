package admins

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"gorm.io/gorm"
	"strings"
)

// RoleIndex 角色列表
func RoleIndex(params *dtoadmins.RoleIndexParams) (interface{}, error) {
	data := dto.IndexData{Items: make([]*dtoadmins.RoleIndexData, 0)}

	//添加过滤where条件
	filterParams := utils.NewFilterParams().
		Like("name LIKE ?", params.Name+"%").
		Like("desc LIKE ?", params.Desc+"%").
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	params.Pagination.SortBy = "name"
	if result := database.Db.Table("auth_item").
		Where("type=?", models.AuthItemTypeRole).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	//	获取所有路由名称
	routerNameList := GetRouterNameList()
	for _, item := range data.Items.([]*dtoadmins.RoleIndexData) {
		//	获取管理角色 对应的 路由名称
		currentAuthList := []string{}
		roleRouteNameList := rds.GetRoleRouteName(item.RoleInfo.Name)
		item.Authority = map[string]bool{}
		for _, v := range routerNameList {
			if utils.ArrayStringIndexOf(roleRouteNameList, v.Value.(string)) > -1 {
				currentAuthList = append(currentAuthList, v.Label)
				item.Authority[v.Value.(string)] = true
			}
		}
		item.CurrentyAuth = strings.Join(currentAuthList, " | ")
	}

	return data, nil
}

// RoleDelete 删除角色
func RoleDelete(params *dtoadmins.RoleDeleteParams) (interface{}, error) {
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		for _, name := range params.NameList {
			// 删除角色
			result := tx.Where("name = ?", name).Where("type=?", models.AuthItemTypeRole).
				Delete(&models.AuthItem{})
			if result.Error != nil {
				return result.Error
			}

			// 删除对应的权限
			if result = tx.Where("parent = ? ", name).
				Where("type=?", models.AuthChildTypeRoleRouteName).
				Delete(&models.AuthChild{}); result.Error != nil {
				return result.Error
			}
		}
		return nil
	})
	return "ok", err
}

// RoleCreate 新增角色
func RoleCreate(params *dtoadmins.RoleCreateParams) (interface{}, error) {
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		//	新增角色

		if result := tx.Create(&models.AuthItem{
			Type: models.AuthItemTypeRole,
			Name: params.Name,
			Desc: params.Desc,
		}); result.Error != nil {
			return result.Error
		}

		//	加载角色权限
		for k, v := range params.Authority {
			if v {
				if result := tx.Create(&models.AuthChild{
					Parent: params.Name,
					Child:  k,
					Type:   models.AuthChildTypeRoleRouteName,
				}); result.Error != nil {
					return result.Error
				}
			}
		}
		return nil
	})
	return "ok", err
}

// RoleUpdate 更新角色
func RoleUpdate(params *dtoadmins.RoleUpdateParams) (interface{}, error) {
	//	查询角色信息
	roleInfo := &models.AuthItem{}
	if result := database.Db.Model(roleInfo).Where("name = ?", params.Name).Take(roleInfo); result.Error != nil {
		return nil, result.Error
	}

	//事务方法
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		if result := database.Db.Model(&models.AuthItem{}).Where("name = ?", roleInfo.Name).
			Update("desc", params.Desc); result.Error != nil {
			return result.Error
		}

		//	删除原先权限
		if result := database.Db.Where("parent = ?", params.Name).Where("type=?", models.AuthChildTypeRoleRouteName).Delete(&models.AuthChild{}); result.Error != nil {
			return result.Error
		}

		//	加载角色权限
		for k, v := range params.Authority {
			if v {
				if result := tx.Create(&models.AuthChild{
					Parent: params.Name,
					Child:  k,
					Type:   models.AuthChildTypeRoleRouteName,
				}); result.Error != nil {
					return result.Error
				}
			}
		}
		return nil
	})
	return "ok", err
}
