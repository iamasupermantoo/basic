package router

import (
	"basic/apis/home/controllers/index"
	controllersSocket "basic/apis/home/controllers/socket"
	"basic/apis/home/controllers/systems"
	"basic/config"
	"basic/middleware"
	"basic/models"
	"basic/module/socket"

	"github.com/gofiber/fiber/v2"
)

// InitHomeRouter 前台路由App
func InitHomeRouter(app *fiber.App, conf *config.Basic) *fiber.App {
	api := app.Group(models.HomeRoutePathname)

	//	启用websocket
	api.Get(models.HomeRouteVersion+"/ws", socket.NewSocketConn())

	//	不需要验证路由
	api.Route(models.HomeRouteVersion, func(router fiber.Router) {
		router.Get("/captcha/create", index.NewCaptcha).Name("创建验证码")
		router.Get("/captcha/:id/:width-:height", index.Captcha).Name("显示验证码")
		router.Get("/init", index.Init).Name("初始化数据")
		router.Post("/login", index.Login).Name("用户登录")
		router.Post("/register", index.Register).Name("用户注册")
		router.Post("/socket/subscribe", controllersSocket.Subscribe).Name("Socket订阅消息")

		router.Post("/footer", index.FooterInfo).Name("电脑端页脚信息")
		router.Post("/article/info", index.ArticleInfo).Name("文章信息")
		router.Post("/article/index", index.ArticleIndex).Name("文章信息")
		router.Post("/download", index.DownloadInfo).Name("下载信息")
		router.Post("/translate", index.TranslateList).Name("翻译信息列表")
		router.Post("/helpers", index.HelpersInfo).Name("帮助中心")
	})

	//	需要验证的路由
	privateKey, _ := models.GetServiceRsaPrivate(models.ServiceHomeName)
	api.Use(middleware.InitJwtMiddleware(privateKey))
	api.Route(models.HomeRouteVersion, func(router fiber.Router) {
		router.Post("/socket/bind", controllersSocket.BindUserId).Name("Socket绑定用户ID")
		router.Post("/upload", systems.Upload).Name("上传文件")
		//	用户模块
		UsersRouter(router)
		// 钱包路由
		WalletRouter(router)
	})

	//	加载swagger
	if conf.Debug {
		app.Use(middleware.InitSwaggerMiddleware())
	}
	return app
}
