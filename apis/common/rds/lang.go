package rds

import (
	dtousers "basic/apis/admin/dto/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

const (
	RedisLanguage = "RedisLanguage"
)

// RedisFindLangList 缓存查询管理语言列表
func RedisFindLangList(rds redis.Conn, adminId int) []*dtousers.UserInitLanguage {
	data, err := RedisGetLangList(rds, adminId)
	if err != nil {
		data = GetLangList(adminId)
		RedisSetLangList(rds, adminId, data)
	}

	return data
}

// GetLangList 获取管理语言列表
func GetLangList(adminId int) []*dtousers.UserInitLanguage {
	data := make([]*dtousers.UserInitLanguage, 0)
	database.Db.Model(&models.Lang{}).Where("admin_id = ?", adminId).Where("status = ?", models.LangStatusActive).Order("sort asc").Find(&data)
	return data
}

// RedisGetLangList 缓存获取管理语言列表
func RedisGetLangList(rds redis.Conn, adminId int) ([]*dtousers.UserInitLanguage, error) {
	dataBytes, err := redis.Bytes(rds.Do("HGET", RedisLanguage, adminId))
	if err != nil {
		return nil, err
	}

	data := make([]*dtousers.UserInitLanguage, 0)
	_ = json.Unmarshal(dataBytes, &data)
	return data, nil
}

// RedisSetLangList 设置缓存管理语言列表
func RedisSetLangList(rds redis.Conn, adminId int, langList []*dtousers.UserInitLanguage) {
	dataBytes, _ := json.Marshal(langList)
	_, _ = rds.Do("HSET", RedisLanguage, adminId, dataBytes)
}

// RedisDelLangList 删除缓存管理语言列表
func RedisDelLangList(rds redis.Conn, adminId int) {
	_, _ = rds.Do("HDEL", RedisLanguage, adminId)
}

// DeleteLangList 删除管理语言列表
func DeleteLangList(ids []int) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	for _, v := range ids {
		langInfo := &models.Lang{}
		result := database.Db.Model(langInfo).Where("id = ?", v).Take(langInfo)
		if result.Error == nil {
			RedisDelLangList(rdsConn, langInfo.AdminId)
		}
	}

	return nil
}
