package database

import (
	"basic/apis/admin/router"
	"basic/models"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// LoadAdminAuth 加载管理权限
func LoadAdminAuth() ([]*models.AuthItem, []*models.AuthChild, []*models.AuthAssignment) {
	//	代理管理员过滤掉的路由
	merchantManageFilterRouter := []string{
		"角色列表", "新增角色", "删除角色", "更新角色信息",
		"后台菜单列表", "后台菜单新增", "后台菜单更新", "后台菜单更新详情", "后台菜单删除",
	}
	agentManageOnlyRouter := []string{
		"查询字段语言", "更新字段语言", "Google翻译",
		"首页信息", "初始化数据", "管理信息", "管理提示音", "更新管理信息", "修改管理密码", "上传文件",
		"新增用户", "用户列表", "用户更新", "删除用户", "用户余额变动", "用户Json模版", "用户关系图",
		"邀请列表", "新增邀请", "更新邀请", "删除邀请", "邀请Json模版",
		"会员列表", "新增会员", "更新会员", "删除会员", "会员Json模版",
		"认证列表", "新增认证", "更新认证", "删除认证", "认证Json模版",
		"用户充值审核", "用户充值新增", "用户充值删除", "用户充值更新", "用户充值列表", "用户充值Json模版",
		"用户提现新增", "用户提现删除", "用户提现更新", "用户提现列表", "用户提现审核", "用户提现Json模版",
		"用户资产修改", "用户资产删除", "用户资产列表", "用户资产更新", "用户资产json模版",
		"用户账单列表", "用户账单Json模版",
		"用户提现账户新增", "用户提现账户删除", "用户提现账户更新", "用户提现账户列表", "用户提现账户Json模版",
	}
	authItem := []*models.AuthItem{
		{Name: models.AuthRoleSuperManage, Type: models.AuthItemTypeRole},
		{Name: models.AuthRoleMerchatManage, Type: models.AuthItemTypeRole},
		{Name: models.AuthRoleAgentManage, Type: models.AuthItemTypeRole},
		{Name: "*", Type: models.AuthItemTypeRoute},
		{Name: "所有权限", Type: models.AuthItemTypeName},
	}

	authChild := []*models.AuthChild{
		{Parent: "所有权限", Child: "*", Type: models.AuthChildTypeRouteNameRoute},
		{Parent: models.AuthRoleSuperManage, Child: "所有权限", Type: models.AuthChildTypeRoleRouteName},
		{Parent: models.AuthRoleSuperManage, Child: models.AuthRoleMerchatManage, Type: models.AuthChildTypeRoleParentRole},
		{Parent: models.AuthRoleMerchatManage, Child: models.AuthRoleAgentManage, Type: models.AuthChildTypeRoleParentRole},
	}

	app := fiber.New()
	router.InitAdminRouter(app, true)
	routerList := app.GetRoutes()
	methodList := []string{"GET", "POST"}
	for _, v := range routerList {
		if v.Name != "" && utils.ArrayStringIndexOf(methodList, v.Method) > -1 {
			//	基础路由载入
			authItem = append(authItem, &models.AuthItem{Name: v.Name, Type: models.AuthItemTypeName})
			authItem = append(authItem, &models.AuthItem{Name: v.Path, Type: models.AuthItemTypeRoute})
			authChild = append(authChild, &models.AuthChild{Parent: v.Name, Child: v.Path, Type: models.AuthChildTypeRouteNameRoute})

			//	代理管理是否过滤当前路由
			if utils.ArrayStringIndexOf(merchantManageFilterRouter, v.Name) == -1 {
				authChild = append(authChild, &models.AuthChild{Parent: models.AuthRoleMerchatManage, Child: v.Name, Type: models.AuthChildTypeRoleRouteName})
			}

			//	普通管理是否拥有当前路由
			if utils.ArrayStringIndexOf(agentManageOnlyRouter, v.Name) > -1 {
				authChild = append(authChild, &models.AuthChild{Parent: models.AuthRoleAgentManage, Child: v.Name, Type: models.AuthChildTypeRoleRouteName})
			}
		}
	}

	// 分配角色
	authAssignment := []*models.AuthAssignment{
		{Name: models.AuthRoleSuperManage, AdminId: models.SuperAdminId},
		{Name: models.AuthRoleMerchatManage, AdminId: models.AdminMerchat},
		{Name: models.AuthRoleAgentManage, AdminId: models.AdminAgent},
	}

	return authItem, authChild, authAssignment
}
