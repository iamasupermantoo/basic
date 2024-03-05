package rds

import (
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

const (
	RedisAdminSettingName = "_AdminSetting"
)

// RedisFindAdminSettingField 获取管理配置
func RedisFindAdminSettingField(rds redis.Conn, adminId int, field string) string {
	adminSettingId := RedisFindAdminIdToSettingAdminId(rds, adminId)
	value, err := RedisGetAdminSettingField(rds, adminSettingId, field)
	if err != nil {
		database.Db.Model(&models.AdminSetting{}).Select("value").Where("field = ?", field).Where("admin_id = ?", adminSettingId).Find(&value)
		RedisSetAdminSettingField(rds, adminSettingId, field, value)
		return value
	}
	return value
}

// RedisGetAdminSettingField 获取缓存配置信息
func RedisGetAdminSettingField(rds redis.Conn, adminSettingId int, field string) (string, error) {
	value, err := redis.String(rds.Do("HGET", RedisAdminSettingName+strconv.Itoa(adminSettingId), field))
	if err != nil {
		return "", err
	}
	return value, nil
}

// RedisSetAdminSettingField 设置缓存配置信息
func RedisSetAdminSettingField(rds redis.Conn, adminSettingId int, field, value string) {
	_, _ = rds.Do("HSET", RedisAdminSettingName+strconv.Itoa(adminSettingId), field, value)
}

// RedisDelAdminSettingField 删除缓存管理配置信息
func RedisDelAdminSettingField(rds redis.Conn, adminSettingId int, field string) {
	_, _ = rds.Do("HDEL", RedisAdminSettingName+strconv.Itoa(adminSettingId), field)
}

// RedisDelAllAdminSetting 删除管理缓存配置
func RedisDelAllAdminSetting(rds redis.Conn, adminSettingId int) {
	_, _ = rds.Do("DEL", RedisAdminSettingName+strconv.Itoa(adminSettingId))
}

// DeleteRedisAdminSetting 删除缓存管理设置
func DeleteRedisAdminSetting(ids []int) error {
	rds := cache.Rds.Get()
	defer rds.Close()

	for _, v := range ids {
		settingInfo := &models.AdminSetting{}
		result := database.Db.Model(settingInfo).Where("id = ?", v).Take(settingInfo)
		if result.Error == nil {
			RedisDelAdminSettingField(rds, settingInfo.AdminId, settingInfo.Field)
		}
	}
	return nil
}
