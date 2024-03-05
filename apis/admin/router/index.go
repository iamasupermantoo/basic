package router

import (
	"basic/apis/admin/controllers/index"
	"basic/middleware"
	"basic/models"

	"github.com/gofiber/fiber/v2"
)

// InitAdminRouter 初始化后台路由A
func InitAdminRouter(app *fiber.App, isInit bool) *fiber.App {
	api := app.Group(models.AdminRoutePathname)
	if !isInit {
		//	不需要验证路由
		api.Route(models.AdminRouteVersion, func(router fiber.Router) {
			router.Get("/captcha/create", index.NewCaptcha).Name("创建验证码")
			router.Get("/captcha/:id/:width-:height", index.Captcha).Name("显示验证码")
			router.Post("/login", index.Login).Name("管理登录")
		})
	}

	//	需要验证的方法
	privateKey, _ := models.GetServiceRsaPrivate(models.ServiceAdminName)
	api.Use(middleware.InitJwtMiddleware(privateKey))
	api.Route(models.AdminRouteVersion, func(router fiber.Router) {
		router.Get("/index", index.Index).Name("首页信息")
		router.Get("/init", index.Init).Name("初始化数据")
		router.Get("/info", index.Info).Name("管理信息")
		router.Get("/audio", index.Audio).Name("管理提示音")
		router.Post("/update", index.Update).Name("更新管理信息")
		router.Post("/update/password", index.UpdatePassword).Name("修改管理密码")
		router.Post("/upload", index.Upload).Name("上传文件")

		//	用户模块
		UsersRouter(router)

		//	系统国家模块
		SystemRouter(router)

		// 钱包模块
		WalletRouter(router)

		// 设置管理
		SettingsRouter(router)

		// 管理模块
		ManagerRouter(router)
	})

	//	保存后台全部路, 权限验证使用
	models.InitAdminRouter(app)
	return app
}
