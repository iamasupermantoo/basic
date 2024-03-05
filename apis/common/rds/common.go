package rds

import "basic/module/cache"

// InitAdminRedis 初始化管理缓存数据
func InitAdminRedis(adminId int) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	管理国家列表缓存
	RedisDelCountryList(rdsConn, adminId)

	//	删除所有缓存
	DelRedisInvite(rdsConn)

	//	删除管理语言列表
	RedisDelLangList(rdsConn, adminId)

	//	删除管理前台菜单
	RedisDelHomeMenuList(rdsConn, adminId)

	//	删除管理缓存配置
	RedisDelAllAdminSetting(rdsConn, adminId)

	//	删除所有查询的翻译缓存
	RedisDelAllTranslateAll(rdsConn, adminId)

	//	删除前台翻译缓存
	RedisDelAllHomeTranslateLang(rdsConn, adminId)
}
