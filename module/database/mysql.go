package database

import (
	"basic/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

// InitGorm 初始化Gorm
func InitGorm(conf *config.Config) {
	databaseConf := conf.Database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=Local", databaseConf.User, databaseConf.Pass, databaseConf.Server, databaseConf.Port, databaseConf.DbName)

	//	设置logger等级
	loggerLevel := logger.Info
	if !conf.Basic.Debug {
		loggerLevel = logger.Error
	}

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(loggerLevel),
	})

	Db = db
}
