package router

import (
	"basic/apis/home/controllers/users/auth"
	"basic/apis/home/controllers/users/invite"
	"basic/apis/home/controllers/users/level"
	"basic/apis/home/controllers/users/team"
	"basic/apis/home/controllers/users/user"

	"github.com/gofiber/fiber/v2"
)

// UsersRouter 前台用户模块
func UsersRouter(router fiber.Router) {
	usersRouter := router.Group("/user")

	//用户方法
	usersRouter.Post("/info", user.Info).Name("当前用户信息")
	usersRouter.Post("/update", user.Update).Name("更新当前用户信息")
	usersRouter.Post("/update/password", user.UpdatePassword).Name("更新当前用户密码")

	//用户等级方法
	usersRouter.Post("/level/index", level.Index).Name("等级列表")
	usersRouter.Post("/level/create", level.Create).Name("用户购买等级")

	//用户团队方法
	usersRouter.Post("/team/index", team.Index).Name("团队成员列表")
	usersRouter.Post("/team/details", team.Details).Name("团队详情")

	//	邀请码新增和分享
	usersRouter.Post("/invite/info", invite.Info).Name("邀请信息")

	//	认证管理
	usersRouter.Post("/auth/create", auth.Create).Name("申请｜重提认证")
	usersRouter.Post("/auth/info", auth.Info).Name("查询认证")
}
