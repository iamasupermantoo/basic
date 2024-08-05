package rds

import (
	dtoadmins "basic/apis/admin/dto/admins"
	dtousers "basic/apis/admin/dto/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

const (
	RedisHomeMenu = "RedisHomeMenu"
)

// RedisFindHomeMenuList 缓存查询前台菜单列表
func RedisFindHomeMenuList(rds redis.Conn, adminId int) *dtousers.HomeMenuList {
	data, err := RedisGetHomeMenuList(rds, adminId)
	if err != nil {
		data = GetHomeMenuList(rds, adminId)
		RedisSetHomeMenuList(rds, adminId, data)
	}

	return data
}

// RedisGetHomeMenuList 获取缓存前台菜单列表
func RedisGetHomeMenuList(rds redis.Conn, adminId int) (*dtousers.HomeMenuList, error) {
	dataBytes, err := redis.Bytes(rds.Do("HGET", RedisHomeMenu, adminId))
	if err != nil {
		return nil, err
	}

	data := &dtousers.HomeMenuList{}
	_ = json.Unmarshal(dataBytes, &data)
	return data, nil
}

// RedisSetHomeMenuList 设置缓存前台菜单列表
func RedisSetHomeMenuList(rds redis.Conn, adminId int, data *dtousers.HomeMenuList) {
	dataBytes, _ := json.Marshal(data)
	_, _ = rds.Do("HSET", RedisHomeMenu, adminId, dataBytes)
}

// RedisDelHomeMenuList 删除缓存前台菜单列表
func RedisDelHomeMenuList(rds redis.Conn, adminId int) {
	_, _ = rds.Do("HDEL", RedisHomeMenu, adminId)
}

// DeleteHomeMenuList 删除前台缓存菜单列表
func DeleteHomeMenuList(ids []int) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	for _, v := range ids {
		menuInfo := &models.AdminMenu{}
		result := database.Db.Model(menuInfo).Where("id = ?", v).Take(menuInfo)
		if result.Error == nil {
			RedisDelHomeMenuList(rdsConn, menuInfo.AdminId)
		}
	}
	return nil
}

// GetHomeMenuList 获取管理前台菜单
func GetHomeMenuList(rds redis.Conn, adminId int) *dtousers.HomeMenuList {
	//	查询管理所有菜单
	menuList := make([]*models.AdminMenu, 0)
	database.Db.Model(&models.AdminMenu{}).Where("admin_id = ?", adminId).
		Where("type IN ?", []int{models.AdminMenuTypeTabularMenu, models.AdminMenuTypeUserMenu, models.AdminMenuTypeQuickMenu, models.AdminMenuTypeHomeMenu}).
		Where("parent_id = ?", 0).
		Where("status = ?", models.AdminMenuStatusActive).Order("sort asc").
		Scan(&menuList)

	//	TabBar菜单 快捷菜单 用户菜单
	data := &dtousers.HomeMenuList{}

	for i := 0; i < len(menuList); i++ {
		dataTmp := new(dtoadmins.AdminMenuData)
		_ = json.Unmarshal([]byte(menuList[i].Data), &dataTmp)

		switch menuList[i].Type {
		case models.AdminMenuTypeTabularMenu:
			//	tabBar菜单
			data.TabbarList = append(data.TabbarList, &dtousers.HomeMenuInfo{
				Id:         menuList[i].Id,
				Name:       menuList[i].Name,
				Route:      menuList[i].Route,
				Icon:       dataTmp.Icon,
				ActiveIcon: dataTmp.ActiveIcon,
				Data:       &dtousers.HomeMenuDataInfo{IsMobile: dataTmp.IsMobile, IsDesktop: dataTmp.IsPc},
				Children:   menuChildren(menuList[i].Id),
			})
		case models.AdminMenuTypeUserMenu:
			//	用户菜单
			data.MenuList = append(data.MenuList, &dtousers.HomeMenuInfo{
				Id:         menuList[i].Id,
				Name:       menuList[i].Name,
				Route:      menuList[i].Route,
				Icon:       dataTmp.Icon,
				ActiveIcon: dataTmp.ActiveIcon,
				Data:       &dtousers.HomeMenuDataInfo{IsMobile: dataTmp.IsMobile, IsDesktop: dataTmp.IsPc},
				Children:   menuChildren(menuList[i].Id),
			})

		case models.AdminMenuTypeQuickMenu:
			//	快捷菜单
			data.QuickList = append(data.QuickList, &dtousers.HomeMenuInfo{
				Id:         menuList[i].Id,
				Name:       menuList[i].Name,
				Route:      menuList[i].Route,
				Icon:       dataTmp.Icon,
				ActiveIcon: dataTmp.ActiveIcon,
				Data:       &dtousers.HomeMenuDataInfo{IsMobile: dataTmp.IsMobile, IsDesktop: dataTmp.IsPc},
				Children:   menuChildren(menuList[i].Id),
			})

		case models.AdminMenuTypeHomeMenu:
			//	前台菜单
			data.HomeList = append(data.HomeList, &dtousers.HomeMenuInfo{
				Id:         menuList[i].Id,
				Name:       menuList[i].Name,
				Route:      menuList[i].Route,
				Icon:       dataTmp.Icon,
				ActiveIcon: dataTmp.ActiveIcon,
				Data:       &dtousers.HomeMenuDataInfo{IsMobile: dataTmp.IsMobile, IsDesktop: dataTmp.IsPc},
				Children:   menuChildren(menuList[i].Id),
			})
		}
	}

	return data
}

func menuChildren(parentId int) []*dtousers.HomeMenuInfo {
	data := make([]*dtousers.HomeMenuInfo, 0)
	menuList := make([]*models.AdminMenu, 0)
	database.Db.Model(&models.AdminMenu{}).Where("parent_id = ?", parentId).Find(&menuList)
	for i := 0; i < len(menuList); i++ {
		if menuList[i].ParentId == parentId {
			dataTmp := new(dtoadmins.AdminMenuData)
			_ = json.Unmarshal([]byte(menuList[i].Data), &dataTmp)
			data = append(data, &dtousers.HomeMenuInfo{
				Id:         menuList[i].Id,
				Name:       menuList[i].Name,
				Route:      menuList[i].Route,
				Icon:       dataTmp.Icon,
				ActiveIcon: dataTmp.ActiveIcon,
				Data:       &dtousers.HomeMenuDataInfo{IsMobile: dataTmp.IsMobile, IsDesktop: dataTmp.IsPc},
				Children:   make([]*dtousers.HomeMenuInfo, 0),
			})
		}
	}
	return data
}
