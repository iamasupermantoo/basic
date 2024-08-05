package router

import (
	"basic/apis/admin/controllers/systems/country"
	"basic/apis/admin/controllers/systems/lang"
	"basic/apis/admin/controllers/systems/level"
	"basic/apis/admin/controllers/systems/notify"
	"basic/apis/admin/controllers/systems/translate"

	"github.com/gofiber/fiber/v2"
)

// SystemRouter 系统模块路由
func SystemRouter(router fiber.Router) {
	system := router.Group("/system")
	//	国家配置
	system.Route("/country", func(router fiber.Router) {
		router.Post("/create", country.Create).Name("新增国家配置")
		router.Post("/index", country.Index).Name("查询国家配置")
		router.Post("/delete", country.Delete).Name("删除国家配置")
		router.Post("/update", country.Update).Name("更新国家配置")
		router.Post("/views", country.Views).Name("国家配置Json模版")
	})
	//	等级配置
	system.Route("/level", func(router fiber.Router) {
		router.Post("/create", level.Create).Name("新增等级配置")
		router.Post("/index", level.Index).Name("查询等级配置")
		router.Post("/delete", level.Delete).Name("删除等级配置")
		router.Post("/update", level.Update).Name("更新等级配置")
		router.Post("/views", level.Views).Name("等级配置Json模版")
	})
	//	语言配置
	system.Route("/lang", func(router fiber.Router) {
		router.Post("/create", lang.Create).Name("新增语言配置")
		router.Post("/index", lang.Index).Name("查询语言配置")
		router.Post("/delete", lang.Delete).Name("删除语言配置")
		router.Post("/update", lang.Update).Name("更新语言配置")
		router.Post("/views", lang.Views).Name("语言配置Json模版")
	})
	//	翻译配置
	system.Route("/translate", func(router fiber.Router) {
		router.Post("/create", translate.Create).Name("新增翻译配置")
		router.Post("/index", translate.Index).Name("查询翻译配置")
		router.Post("/delete", translate.Delete).Name("删除翻译配置")
		router.Post("/update", translate.Update).Name("更新翻译配置")
		router.Post("/fields", translate.Field).Name("查询字段语言")
		router.Post("/google", translate.GoogleTranslateField).Name("Google翻译")
		router.Post("/update/fields", translate.UpdateLangField).Name("更新字段语言")
		router.Post("/views", translate.Views).Name("翻译配置Json模版")
	})
	//	通知配置
	system.Route("/notify", func(router fiber.Router) {
		router.Post("/create", notify.Create).Name("新增通知配置")
		router.Post("/index", notify.Index).Name("查询通知配置")
		router.Post("/delete", notify.Delete).Name("删除通知配置")
		router.Post("/update", notify.Update).Name("更新通知配置")
		router.Post("/views", notify.Views).Name("通知配置Json模版")
	})

}
