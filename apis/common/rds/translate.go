package rds

import (
	"basic/apis/common/dto"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/gomodule/redigo/redis"
)

const (
	RedisTranslateName = "_Translate"
)

// RedisFindTranslateField 获取翻译值
func RedisFindTranslateField(rds redis.Conn, adminId int, lang, field string) string {
	adminSettingId, err := RedisAdminIdToSettingAdminId(rds, adminId)
	value, err := RedisGetTranslateField(rds, adminSettingId, lang, field)
	if err != nil {
		database.Db.Model(&models.Translate{}).Select("value").Where("admin_id = ?", adminSettingId).Where("lang = ?", lang).Where("field = ?", field).Find(&value)
		RedisSetTranslateField(rds, adminSettingId, lang, field, value)
	}

	if value == "" {
		value = field
	}
	return value
}

// RedisFindHomeTranslateLang 获取管理翻译列表
func RedisFindHomeTranslateLang(rds redis.Conn, adminId int, lang string) []*dto.TranslateOptions {
	adminSettingId, err := RedisAdminIdToSettingAdminId(rds, adminId)
	data, err := RedisGetHomeTranslateLang(rds, adminSettingId, lang)
	if err != nil {
		data = make([]*dto.TranslateOptions, 0)
		database.Db.Model(&models.Translate{}).Select("field as label", "value").Where("admin_id = ?", adminSettingId).Where("type = ?", models.TranslateTypeHome).Where("lang = ?", lang).Find(&data)
		RedisSetHomeTranslateLang(rds, adminSettingId, lang, data)
	}
	return data
}

// RedisGetTranslateField 获取管理缓存翻译
func RedisGetTranslateField(rds redis.Conn, adminId int, lang, field string) (string, error) {
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId) + "_" + lang
	value, err := redis.String(rds.Do("HGET", redisName, field))
	if err != nil {
		return "", err
	}
	return value, nil
}

// RedisSetTranslateField 管理缓存翻译
func RedisSetTranslateField(rds redis.Conn, adminId int, lang, field, value string) {
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId) + "_" + lang
	_, _ = rds.Do("HSET", redisName, field, value)
}

// RedisDelTranslateField 管理缓存翻译删除
func RedisDelTranslateField(rds redis.Conn, adminId int, lang, field string) {
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId) + "_" + lang
	_, _ = rds.Do("HDEL", redisName, field)
}

// RedisDelTranslateAllField 删除翻译所有Field
func RedisDelTranslateAllField(rds redis.Conn, adminId int, lang string) {
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId) + "_" + lang
	_, _ = rds.Do("DEL", redisName)
}

// RedisDelAllTranslateAll 删除管理所有翻译缓存
func RedisDelAllTranslateAll(rds redis.Conn, adminId int) {
	langList := []*models.Lang{}
	database.Db.Model(models.Lang{}).Where("admin_id = ?", adminId).Find(&langList)

	for _, v := range langList {
		RedisDelTranslateAllField(rds, adminId, v.Alias)
	}
}

// RedisGetHomeTranslateLang 获取管理语言翻译列表
func RedisGetHomeTranslateLang(rds redis.Conn, adminId int, lang string) ([]*dto.TranslateOptions, error) {
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId)
	valueBytes, err := redis.Bytes(rds.Do("HGET", redisName, lang))
	if err != nil {
		return nil, err
	}

	data := make([]*dto.TranslateOptions, 0)
	_ = json.Unmarshal(valueBytes, &data)
	return data, nil
}

// RedisSetHomeTranslateLang 缓存管理语言翻译列表
func RedisSetHomeTranslateLang(rds redis.Conn, adminId int, lang string, translateList []*dto.TranslateOptions) {
	valueBytes, _ := json.Marshal(translateList)
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId)
	_, _ = rds.Do("HSET", redisName, lang, valueBytes)
}

// RedisDelHomeTranslateLang 删除管理语言翻译列表
func RedisDelHomeTranslateLang(rds redis.Conn, adminId int, lang string) {
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId)
	_, _ = rds.Do("HDEL", redisName, lang)
}

// RedisDelAllHomeTranslateLang 删除前台翻译列表
func RedisDelAllHomeTranslateLang(rds redis.Conn, adminId int) {
	redisName := RedisTranslateName + "_" + strconv.Itoa(adminId)
	_, _ = rds.Do("DEL", redisName)
}

// DeleteRedisTranslate 删除所有语言翻译
func DeleteRedisTranslate(ids []int) error {
	rds := cache.Rds.Get()
	defer rds.Close()

	for _, v := range ids {
		translateInfo := &models.Translate{}
		result := database.Db.Model(translateInfo).Where("id = ?", v).Take(translateInfo)

		//	删除缓存中字段
		if result.Error == nil {
			RedisDelTranslateField(rds, translateInfo.AdminId, translateInfo.Lang, translateInfo.Field)

			if translateInfo.Type == models.TranslateTypeHome {
				RedisDelHomeTranslateLang(rds, translateInfo.AdminId, translateInfo.Lang)
			}
		}
	}

	return nil
}
