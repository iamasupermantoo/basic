package admin

import (
	"basic/apis/admin/router"
	"basic/middleware"
	"basic/models"
	"github.com/gofiber/fiber/v2"
)

// InitAdminRouterFiberApp 启动后台APP
func InitAdminRouterFiberApp() *fiber.App {
	app := fiber.New(fiber.Config{
		AppName: models.ServiceAdminName,
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
	router.InitAdminRouter(app, false)

	return app
}
