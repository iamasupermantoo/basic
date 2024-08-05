package home

import (
	"basic/apis/home/router"
	"basic/config"
	"basic/middleware"
	"basic/models"

	"github.com/gofiber/fiber/v2"
)

// InitHomeRouterFiberApp 启动前台App
func InitHomeRouterFiberApp(conf *config.Basic) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: models.ServiceHomeName,
	})

	//	初始化跨域中间件
	app.Use(middleware.InitCorsMiddleware())

	//	初始化速率限制中间件
	app.Use(middleware.InitLimiterMiddleware())

	//	初始化异常错误中间件
	app.Use(middleware.InitRecoverMiddleware())

	//	初始化日志中间件
	app.Use(middleware.InitLoggerMiddleware())

	//	初始化路由文件
	router.InitHomeRouter(app, conf)

	return app
}
