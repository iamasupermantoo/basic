package router

import (
	"basic/apis/admin/controllers/manager/logs"
	"basic/apis/admin/controllers/manager/manage"
	"basic/apis/admin/controllers/manager/menu"
	"basic/apis/admin/controllers/manager/role"

	"github.com/gofiber/fiber/v2"
)

// ManagerRouter 管理模块
func ManagerRouter(router fiber.Router) {
	manageRouter := router.Group("/manage")

	//操作日志管理
	manageRouter.Post("/logs/index", logs.Index).Name("操作日志列表")
	manageRouter.Post("/logs/views", logs.Views).Name("操作日志Json模板")

	//角色管理
	manageRouter.Post("/role/create", role.Create).Name("新增角色")
	manageRouter.Post("/role/delete", role.Delete).Name("删除角色")
	manageRouter.Post("/role/index", role.Index).Name("角色列表")
	manageRouter.Post("/role/update", role.Update).Name("更新角色信息")
	manageRouter.Post("/role/views", role.Views).Name("角色Json模版")

	//管理
	manageRouter.Post("/index", manage.Index).Name("管理列表")
	manageRouter.Post("/create", manage.Create).Name("新增管理")
	manageRouter.Post("/delete", manage.Delete).Name("删除管理")
	manageRouter.Post("/update", manage.Update).Name("更新管理")
	manageRouter.Post("/update/desc", manage.UpdateDesc).Name("更新管理配置")
	manageRouter.Post("/reset", manage.MerchantReset).Name("重置商户数据")
	manageRouter.Post("/sync", manage.SynchronousMerchantData).Name("同步商户数据")
	manageRouter.Post("/views", manage.Views).Name("管理Json模板")

	//菜单管理
	manageRouter.Post("/menu/index", menu.Index).Name("后台菜单列表")
	manageRouter.Post("/menu/create", menu.Create).Name("后台菜单新增")
	manageRouter.Post("/menu/update", menu.Update).Name("后台菜单更新")
	manageRouter.Post("/menu/update/desc", menu.UpdateDesc).Name("后台菜单更新详情")
	manageRouter.Post("/menu/delete", menu.Delete).Name("后台菜单删除")
	manageRouter.Post("/menu/views", menu.Views).Name("后台菜单Json模板")
}
