package apis

import (
	"basic/apis/admin"
	"basic/apis/home"
	"basic/config"
	"basic/middleware"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// InitService 初始化服务
func InitService(app *fiber.App, conf *config.Basic) *fiber.App {
	//	开启静态文件
	app.Static("/", "./public")
	app.Use(middleware.InitFaviconMiddleware())

	//	初始化前端路由App
	homeFiber := home.InitHomeRouterFiberApp(conf)

	//	初始化后端路由App
	adminFiber := admin.InitAdminRouterFiberApp()

	//	初始化域名路由中间件
	hosts := map[string]*middleware.Host{middleware.OtherDomain: {Fiber: homeFiber}}
	adminDomainsList := strings.Split(conf.AdmDomain, ",")
	for _, v := range adminDomainsList {
		hosts[v] = &middleware.Host{Fiber: adminFiber}
	}
	app.Use(middleware.InitDomainRoutingMiddleware(hosts))

	return app
}
