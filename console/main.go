package main

import (
	"basic/config"
	"basic/console/database"
	"basic/module"
	"fmt"
)

func main() {
	//	读取配置文件
	conf := config.ReadYamlConfigFile()

	//	初始化模块
	module.InitModule(conf)
	fmt.Println("1) 载入初始化数据库")

	var opt int
	_, _ = fmt.Scan(&opt)

	switch opt {
	//	初始化数据库
	case 1:
		database.InitDatabase()
	case 2:
		database.LoadAdminAuth()
	}
}
