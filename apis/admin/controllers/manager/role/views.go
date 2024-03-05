package role

import (
	"basic/apis/admin/service/admins"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

// Views 角色Json模板
func Views(c *fiber.Ctx) error {
	authOptions := admins.GetRouterNameList()

	config := views.NewTableViews("/manage/role/index", "/manage/role/update").
		SetTableKey("name")

	//	查询参数设置
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("角色", "name").
		Text("描述", "desc").
		RangeDatePicker("创建时间", "createdAt").GetInputListInfo()

	//	新增角色参数设置
	createParams, createInputList := views.NewInputViews().
		Text("角色", "name").
		Text("描述", "desc").
		Checkbox("权限", "authority", authOptions). //权限列表数组
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "新增角色", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", "/manage/role/create", "新增角色").
				SetParams(createParams).
				SetInputList(createInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "删除角色", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", "/manage/role/delete", "删除角色").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "nameList", Scan: "name"}),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("角色", "name", true).
		Text("权限", "currentAuth", false).
		Text("描述", "desc", true).
		DatePicker("更新时间", "updatedAt", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	更新参数列表
	updateParams, updateInputList := views.NewInputViews().
		Text("描述", "desc").
		Checkbox("权限", "authority", authOptions). //权限数组
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/manage/role/update", "更新角色信息").
				SetParams(updateParams).
				SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)
	return c.JSON(utils.SuccessJson(config))
}
