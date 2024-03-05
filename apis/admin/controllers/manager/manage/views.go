package manage

import (
	"basic/apis/common/rds"
	"basic/models"
	"basic/utils"
	"basic/utils/views"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Views 管理Json模板
func Views(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)
	roleListOptions := rds.GetRolesList(adminId)
	config := views.NewTableViews("/manage/index", "/manage/update")

	//	查询参数设置
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("上级名称", "parentName").
		Text("管理名称", "username").
		Text("昵称", "nickname").
		Text("邮箱", "email").
		Text("域名", "domains").
		Text("角色", "role").
		Select("状态", "status", []*views.InputOptions{{Label: "冻结", Value: models.AdminUserStatusDisable}, {Label: "激活", Value: models.LangStatusActive}}).
		RangeDatePicker("过期时间", "expiredAt").GetInputListInfo()

	//	新增管理参数设置
	createViews := views.NewInputViews().
		Checkbox("角色", "roles", roleListOptions).
		SetValue("roles", map[string]bool{roleListOptions[0].Value.(string): true}).
		Text("管理名称", "username").
		Password("密码", "password")
	if adminId == models.SuperAdminId {
		createViews.Text("域名", "domains").DatePicker("过期时间", "expiredAt").SetValue("expiredAt", time.Now().AddDate(0, 1, 0).Unix())
	}
	createParams, createInputList := createViews.GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "新增管理", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", "/manage/create", "新增管理").SetSizingSmall().
				SetParams(createParams).
				SetInputList(createInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "删除管理", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", "/manage/delete", "删除管理").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "username", Field: "ids", Scan: "id"}),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Image("头像", "avatar", false).
		Text("上级用户", "parentName", true).
		EditText("域名", "domains", true).
		Text("管理名称", "username", true).
		EditText("昵称", "nickname", true).
		Text("角色", "rolesStr", false).
		EditText("邮箱", "email", true).
		Text("金额", "money", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "激活", Value: models.AdminUserStatusActive}, {Label: "冻结", Value: models.AdminUserStatusDisable}}).
		DatePicker("创建时间", "createdAt", true).
		DatePicker("过期时间", "expiredAt", true).
		GetColumnsListInfo()

	updateParams, updateInputList := views.NewInputViews().
		Image("头像", "avatar").
		Text("邮箱", "email").
		Text("昵称", "nickname").
		Password("密码", "password").
		Password("安全密码", "securityKey").
		DatePicker("过期时间", "expiredTime").
		Select("状态", "status", []*views.InputOptions{{Label: "冻结", Value: models.AdminUserStatusDisable}, {Label: "激活", Value: models.AdminUserStatusActive}}).
		Checkbox("角色", "roles", roleListOptions).
		GetInputListInfo()

	dataJsonParam, dataJsonInputList := views.NewInputViews().
		Select("前端模版", "template", []*views.InputOptions{{Label: "默认模版", Value: "default"}}).
		Number("代理数量", "nums").GetInputListInfo()
	optionsParam, optionsInputList := views.NewInputViews().
		Json("", "dataJson", dataJsonInputList).SetValue("dataJson", dataJsonParam).
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/manage/update", "更新管理信息").
				SetParams(updateParams).SetSizingSmall().
				SetInputList(updateInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "配置管理", Color: views.ColorSecondary, Size: "xs", Eval: "scope.row.parentId==1"},
			Config: views.NewDialogViews("setting", "/manage/update/desc", "配置管理").
				SetParams(optionsParam).SetInputList(optionsInputList).
				SetSizingSmall(),
		},
		{
			ButtonViews: views.ButtonViews{Label: "同步管理", Color: views.ColorPositive, Size: "xs", Eval: "scope.row.parentId==1"},
			Config: views.NewDialogViews("synchronous", "/manage/sync", "同步管理").
				SetContent("你确定要同步管理员配置吗？").
				SetSizingSmall(),
		},
		{
			ButtonViews: views.ButtonViews{Label: "重置管理", Color: views.ColorNegative, Size: "xs", Eval: "scope.row.parentId==1"},
			Config: views.NewDialogViews("reset", "/manage/reset", "重置管理").
				SetContent("你确定重置该用户吗？").
				SetSizingSmall(),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)
	return c.JSON(utils.SuccessJson(config))
}
