package router

import (
	"basic/apis/admin/controllers/users/auth"
	"basic/apis/admin/controllers/users/invite"
	"basic/apis/admin/controllers/users/level"
	"basic/apis/admin/controllers/users/user"
	"github.com/gofiber/fiber/v2"
)

// UsersRouter 用户模块路由
func UsersRouter(router fiber.Router) {
	//	用户管理
	router.Post("/user/create", user.Create).Name("新增用户")
	router.Post("/user/index", user.Index).Name("用户列表")
	router.Post("/user/update", user.Update).Name("用户更新")
	router.Post("/user/delete", user.Delete).Name("删除用户")
	router.Post("/user/balance", user.Balance).Name("用户余额变动")
	router.Post("/user/views", user.Views).Name("用户Json模版")
	router.Post("/user/relation", user.Relation).Name("用户关系图")

	//	邀请管理
	router.Post("/user/invite/index", invite.Index).Name("邀请列表")
	router.Post("/user/invite/create", invite.Create).Name("新增邀请")
	router.Post("/user/invite/update", invite.Update).Name("更新邀请")
	router.Post("/user/invite/delete", invite.Delete).Name("删除邀请")
	router.Post("/user/invite/views", invite.Views).Name("邀请Json模版")

	//	会员管理
	router.Post("/user/level/index", level.Index).Name("会员列表")
	router.Post("/user/level/create", level.Create).Name("新增会员")
	router.Post("/user/level/update", level.Update).Name("更新会员")
	router.Post("/user/level/delete", level.Delete).Name("删除会员")
	router.Post("/user/level/views", level.Views).Name("会员Json模版")

	//	认证管理
	router.Post("/user/auth/index", auth.Index).Name("认证列表")
	router.Post("/user/auth/create", auth.Create).Name("新增认证")
	router.Post("/user/auth/update", auth.Update).Name("更新认证")
	router.Post("/user/auth/delete", auth.Delete).Name("删除认证")
	router.Post("/user/auth/views", auth.Views).Name("认证Json模版")
}
