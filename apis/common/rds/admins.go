package rds

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

const (
	// RedisAdminChildrenIds 缓存管理下级Ids
	RedisAdminChildrenIds = "RedisAdminChildrenIds"

	// RedisAdminDomainToAdminId 缓存管理域名
	RedisAdminDomainToAdminId = "RedisAdminDomainToAdminId"

	// RedisAdminToSettingAdminId 缓存配置的管理ID
	RedisAdminToSettingAdminId = "RedisAdminToSettingAdminId"

	// RedisAdminDataInfo 缓存管理数据
	RedisAdminDataInfo = "RedisAdminDataInfo"
)

// RedisFindAdminChildrenIds 缓存查询管理下级Ids
func RedisFindAdminChildrenIds(adminId int) []int {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	data, err := RedisGetAdminChildren(rdsConn, adminId)
	if err != nil {
		data = GetAdminChildren(adminId)
		RedisSetAdminChildren(rdsConn, adminId, data)
	}
	return data
}

// RedisFindDomainToAdminId 通过域名获取管理ID
func RedisFindDomainToAdminId(rds redis.Conn, domain string) int {
	adminId, err := GetRedisDomainToAdminId(rds, domain)
	if err != nil {
		adminUserInfo := &models.AdminUser{}
		result := database.Db.Where("FIND_IN_SET(?, domains)", domain).Where("parent_id = ?", models.SuperAdminId).Find(&adminUserInfo)
		if result.Error != nil {
			panic(result.Error)
		}

		//	如果找不到域名对应的数据, 那么使用超级管理ID
		adminId = adminUserInfo.Id
		if result.RowsAffected == 0 {
			adminId = models.SuperAdminId
		}

		SetRedisDomainToAdminId(rds, domain, adminId)
	}

	return adminId
}

// RedisFindAdminIdToSettingAdminId 获取管理设置ID
func RedisFindAdminIdToSettingAdminId(rds redis.Conn, adminId int) int {
	adminSettingId, err := RedisAdminIdToSettingAdminId(rds, adminId)
	if err != nil {
		adminSettingId = GetAdminSettingId(adminId)
		RedisSetAdminIdToSettingAdminId(rds, adminId, adminSettingId)
	}

	return adminSettingId
}

// RedisFindAdminDataInfo 获取缓存管理信息
func RedisFindAdminDataInfo(rds redis.Conn, adminId int) *dtoadmins.AdminInfoData {
	dataInfo, err := RedisGetAdminDataInfo(rds, adminId)
	if err != nil {
		dataInfo = GetAdminDataInfo(adminId)
		RedisSetAdminDataInfo(rds, adminId, dataInfo)
	}
	return dataInfo
}

// SetRedisDomainToAdminId 设置域名对应的管理ID缓存
func SetRedisDomainToAdminId(rds redis.Conn, domain string, adminId int) {
	_, _ = rds.Do("HSET", RedisAdminDomainToAdminId, domain, adminId)
}

// GetRedisDomainToAdminId 获取缓存域名对应的管理ID
func GetRedisDomainToAdminId(rds redis.Conn, domain string) (int, error) {
	adminId, err := redis.Int(rds.Do("HGET", RedisAdminDomainToAdminId, domain))
	if err != nil {
		return 0, err
	}
	return adminId, nil
}

// DelRedisDomainToAdminId 删除缓存对应的管理ID
func DelRedisDomainToAdminId(rds redis.Conn) {
	_, _ = rds.Do("DEL", RedisAdminDomainToAdminId)
}

// GetAdminSettingId 获取管理设置ID
func GetAdminSettingId(adminId int) int {
	if adminId <= models.SuperAdminId {
		return models.SuperAdminId
	}

	adminInfo := &models.AdminUser{}
	result := database.Db.Model(adminInfo).Where("id = ?", adminId).Take(adminInfo)
	if result.Error != nil {
		panic(result.Error)
	}
	if adminInfo.ParentId != models.SuperAdminId {
		return GetAdminSettingId(adminInfo.ParentId)
	}
	return adminInfo.Id
}

// RedisAdminIdToSettingAdminId 获取缓存管理设置ID
func RedisAdminIdToSettingAdminId(rds redis.Conn, adminId int) (int, error) {
	adminSettingId, err := redis.Int(rds.Do("HGET", RedisAdminToSettingAdminId, adminId))
	if err != nil {
		return 0, err
	}
	return adminSettingId, nil
}

// RedisSetAdminIdToSettingAdminId 设置管理设置ID
func RedisSetAdminIdToSettingAdminId(rds redis.Conn, adminId, adminSettingId int) {
	_, _ = rds.Do("HSET", RedisAdminToSettingAdminId, adminId, adminSettingId)
}

// GetAdminDataInfo 获取管理配置信息
func GetAdminDataInfo(adminId int) *dtoadmins.AdminInfoData {
	dataInfo := &dtoadmins.AdminInfoData{}
	adminInfo := &models.AdminUser{}
	database.Db.Model(adminInfo).Where("id = ?", adminId).Find(adminInfo)

	//	如果当前管理信息是超级管理或不是商户, 那么使用默认值
	if adminInfo.Id == models.SuperAdminId || adminInfo.ParentId != models.SuperAdminId {
		dataInfo = &dtoadmins.AdminInfoData{Nums: 0, Template: "default"}
	}

	_ = json.Unmarshal([]byte(adminInfo.Data), &dataInfo)
	return dataInfo
}

// RedisGetAdminDataInfo 缓存获取管理配置信息
func RedisGetAdminDataInfo(rds redis.Conn, adminId int) (*dtoadmins.AdminInfoData, error) {
	dataBytes, err := redis.Bytes(rds.Do("HGET", RedisAdminDataInfo, adminId))
	if err != nil {
		return nil, err
	}

	data := &dtoadmins.AdminInfoData{}
	_ = json.Unmarshal(dataBytes, &data)
	return data, nil
}

// RedisSetAdminDataInfo 设置管理配置信息
func RedisSetAdminDataInfo(rds redis.Conn, adminId int, data *dtoadmins.AdminInfoData) {
	dataBytes, _ := json.Marshal(data)
	_, _ = rds.Do("HSET", RedisAdminDataInfo, adminId, dataBytes)
}

// RedisDelAdminDataInfo 删除管理信息
func RedisDelAdminDataInfo(rds redis.Conn, adminId int) {
	_, _ = rds.Do("HDEL", RedisAdminDataInfo, adminId)
}

// GetAdminChildren 递归获取管理子级
func GetAdminChildren(adminId int) []int {
	adminIds := []int{adminId}

	//	为了取admin_id = 0 的数据
	if adminId == models.SuperAdminId && len(adminIds) == 1 {
		adminIds = append(adminIds, 0)
	}

	ids := make([]int, 0)
	database.Db.Model(&models.AdminUser{}).Select("id").Where("parent_id = ?", adminId).Where("status > ?", models.AdminUserStatusDelete).Find(&ids)
	for _, v := range ids {
		adminIds = append(adminIds, GetAdminChildren(v)...)
	}

	return adminIds
}

// RedisGetAdminChildren 获取管理下级IDs缓存
func RedisGetAdminChildren(rds redis.Conn, adminId int) ([]int, error) {
	dataBytes, err := redis.Bytes(rds.Do("HGET", RedisAdminChildrenIds, adminId))
	if err != nil {
		return nil, err
	}

	data := make([]int, 0)
	_ = json.Unmarshal(dataBytes, &data)
	return data, nil
}

// RedisSetAdminChildren 设置管理下级IDs缓存
func RedisSetAdminChildren(rds redis.Conn, adminId int, data []int) {
	dataBytes, _ := json.Marshal(data)
	_, _ = rds.Do("HSET", RedisAdminChildrenIds, adminId, dataBytes)
}

// RedisDelAdminChildren 删除管理下级IDs缓存
func RedisDelAdminChildren(rds redis.Conn, adminId int) {
	_, _ = rds.Do("HDEL", RedisAdminChildrenIds, adminId)
}

// RedisDelAminAllChildren 清除所有管理下级IDS缓存
func RedisDelAminAllChildren(rds redis.Conn) {
	adminList := make([]*models.AdminUser, 0)
	database.Db.Model(&models.AdminUser{}).Find(&adminList)

	for _, adminInfo := range adminList {
		RedisDelAdminChildren(rds, adminInfo.Id)
	}
}
