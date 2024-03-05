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
	RedisCountry = "RedisCountry"
)

// RedisFindCountryList 缓存查询国家列表
func RedisFindCountryList(rds redis.Conn, adminId int) []*dtousers.UserInitCountry {
	data, err := RedisGetCountryList(rds, adminId)
	if err != nil {
		data = GetCountryList(adminId)
		RedisSetCountryList(rds, adminId, data)
	}
	return data
}

// GetCountryList 获取国家列表
func GetCountryList(adminId int) []*dtousers.UserInitCountry {
	data := make([]*dtousers.UserInitCountry, 0)
	database.Db.Model(&models.Country{}).Where("admin_id = ?", adminId).Where("status = ?", models.CountryStatusActive).Order("sort asc").Find(&data)
	return data
}

// RedisGetCountryList 获取缓存国家列表
func RedisGetCountryList(rds redis.Conn, adminId int) ([]*dtousers.UserInitCountry, error) {
	dataBytes, err := redis.Bytes(rds.Do("HGET", RedisCountry, adminId))
	if err != nil {
		return nil, err
	}

	data := make([]*dtousers.UserInitCountry, 0)
	_ = json.Unmarshal(dataBytes, &data)
	return data, nil
}

// RedisSetCountryList 设置国家缓存
func RedisSetCountryList(rds redis.Conn, adminId int, countryList []*dtousers.UserInitCountry) {
	dataBytes, _ := json.Marshal(countryList)
	_, _ = rds.Do("HSET", RedisCountry, adminId, dataBytes)
}

// RedisDelCountryList 删除国家列表缓存
func RedisDelCountryList(rds redis.Conn, adminId int) {
	_, _ = rds.Do("HDEL", RedisCountry, adminId)
}

// DeleteCountryList 删除国家列表缓存
func DeleteCountryList(ids []int) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	for _, v := range ids {
		countryInfo := &models.Country{}
		result := database.Db.Model(countryInfo).Where("id = ?", v).Take(v)
		if result.Error == nil {
			RedisDelCountryList(rdsConn, countryInfo.AdminId)
		}
	}

	return nil
}
