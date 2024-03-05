package rds

import (
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils/views"

	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

const (
	// AuthAdminRouter 管理路由列表
	AuthAdminRouter = "AuthAdminRouter"
)

// RedisFindAdminRouter 获取管理路由
func RedisFindAdminRouter(adminId int) []string {
	rds := cache.Rds.Get()
	defer rds.Close()

	routerList, err := RedisGetAdminRouter(rds, adminId)
	if err != nil {
		rolesList := GetAdminRoles(adminId)
		for _, role := range rolesList {
			roleRouteList := GetRoleRouteList(role)
			routerList = append(routerList, roleRouteList...)
		}

		RedisSetAdminRouter(rds, adminId, routerList)
	}

	return routerList
}

// RedisGetAdminRouter 获取管理路由列表
func RedisGetAdminRouter(rds redis.Conn, adminId int) ([]string, error) {
	routerListBytes, err := redis.Bytes(rds.Do("HGET", AuthAdminRouter, adminId))
	if err != nil {
		return nil, err
	}

	routerList := make([]string, 0)
	_ = json.Unmarshal(routerListBytes, &routerList)
	return routerList, nil
}

// RedisSetAdminRouter 设置管理路由列表
func RedisSetAdminRouter(rds redis.Conn, adminId int, routerList []string) {
	routerListBytes, _ := json.Marshal(routerList)
	_, _ = rds.Do("HSET", AuthAdminRouter, adminId, routerListBytes)
}

// RedisDelAdminRouter 删除管理路由列表
func RedisDelAdminRouter(rds redis.Conn, adminId int) {
	_, _ = rds.Do("HDEL", AuthAdminRouter, adminId)
}

// GetRoleRouteNamePath 获取路由名称路径
func GetRoleRouteNamePath(routeName string) string {
	authChildInfo := &models.AuthChild{}
	database.Db.Where("parent = ?", routeName).Where("type = ?", models.AuthChildTypeRouteNameRoute).Find(&authChildInfo)
	return authChildInfo.Child
}

// GetRoleRouteName 获取角色底下所有路由名称
func GetRoleRouteName(role string) []string {
	var childs []*models.AuthChild
	database.Db.Where("parent = ?", role).Where("type = ?", models.AuthChildTypeRoleRouteName).Find(&childs)

	data := make([]string, 0)
	for _, v := range childs {
		data = append(data, v.Child)
	}
	return data
}

// GetRoleRouteList 获取角色路由列表
func GetRoleRouteList(role string) []string {
	data := make([]string, 0)
	routeNameList := GetRoleRouteName(role)
	for _, v := range routeNameList {
		routePath := GetRoleRouteNamePath(v)
		if routePath != "" {
			data = append(data, routePath)
		}
	}
	return data
}

// GetAdminRoles 获取管理角色
func GetAdminRoles(adminId int) []string {
	rolesList := []*models.AuthAssignment{}
	database.Db.Where("admin_id = ?", adminId).Find(&rolesList)

	data := make([]string, 0)
	for _, v := range rolesList {
		data = append(data, v.Name)
	}
	return data
}

// GetRolesList 获取全部角色
func GetRolesList(adminId int) []*views.InputOptions {
	data := make([]*views.InputOptions, 0)

	//	获取当前管理角色
	currentAdminRoleList := []string{}
	model := database.Db.Model(&models.AuthChild{}).Select("child").Where("type = ?", models.AuthChildTypeRoleParentRole)
	if adminId > -1 {
		roleList := GetAdminRoles(adminId)
		model.Where("parent IN ?", roleList)
	}
	model.Find(&currentAdminRoleList)

	for i := 0; i < len(currentAdminRoleList); i++ {
		data = append(data, &views.InputOptions{
			Label: currentAdminRoleList[i],
			Value: currentAdminRoleList[i],
		})
	}
	return data
}
