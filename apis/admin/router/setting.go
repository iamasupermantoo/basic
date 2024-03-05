package router

import (
	"basic/apis/admin/controllers/settings/article"
	"basic/apis/admin/controllers/settings/menu"
	"basic/apis/admin/controllers/settings/setting"

	"github.com/gofiber/fiber/v2"
)

// SettingsRouter 设置路由
func SettingsRouter(router fiber.Router) {
	router.Route("/setting", func(router fiber.Router) {
		// 文章管理
		router.Post("/article/create", article.Create).Name("文章创建")
		router.Post("/article/delete", article.Delete).Name("文章删除")
		router.Post("/article/update", article.Update).Name("文章修改")
		router.Post("/article/index", article.Index).Name("文章列表")
		router.Post("/article/views", article.Views).Name("文章json模版")

		// 设置管理
		router.Post("/create", setting.Create).Name("设置创建")
		router.Post("/delete", setting.Delete).Name("设置删除")
		router.Post("/update", setting.Update).Name("设置修改")
		router.Post("/update/desc", setting.UpdateDesc).Name("设置修改详情")
		router.Post("/index", setting.Index).Name("设置列表")
		router.Post("/views", setting.Views).Name("设置json模版")

		//菜单管理
		router.Post("/menu/index", menu.Index).Name("菜单列表")
		router.Post("/menu/create", menu.Create).Name("菜单新增")
		router.Post("/menu/update", menu.Update).Name("菜单更新")
		router.Post("/menu/update/desc", menu.UpdateDesc).Name("菜单更新详情")
		router.Post("/menu/delete", menu.Delete).Name("菜单删除")
		router.Post("/menu/views", menu.Views).Name("菜单Json模板")
	})
}
