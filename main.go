package main

import (
	"basic/apis"
	"basic/config"
	"basic/module"
	"basic/module/cache"

	"github.com/goccy/go-json"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//	读取配置文件
	conf := config.ReadYamlConfigFile()

	//	项目APP
	app := fiber.New(fiber.Config{
		AppName:      conf.Basic.Name,
		ServerHeader: conf.Basic.ServerHeader,
		Prefork:      conf.Basic.Prefork,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})

	//	初始化模块
	module.InitModule(conf)

	//	初始化服务
	apis.InitService(app, conf.Basic)

	// 初始化订阅
	cache.InitSubscribe()

	//	启动监听端口
	_ = app.Listen(":" + conf.Basic.Port)
}
