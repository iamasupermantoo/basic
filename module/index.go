package module

import (
	"basic/config"
	"basic/module/cache"
	"basic/module/crontab"
	"basic/module/database"
)

// InitModule 初始化模块
func InitModule(conf *config.Config) {
	//	初始化定时任务
	crontab.InitCrontab()

	//	初始化数据库
	database.InitGorm(conf)

	//	初始化Redis
	cache.InitRedis(conf.Redis)
}
