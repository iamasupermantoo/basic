package setting

import (
	dtosettings "basic/apis/admin/dto/settings"
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

// Views 模版JSON信息
func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/setting/index", "/setting/update")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("名称", "name").
		Select("分组", "groupId", dtosettings.AdminSettingGroupOptions).
		Select("类型", "type", dtosettings.AdminSettingTypeOptions).
		Text("键名", "field").Text("键值", "value").Text("配置", "data").
		RangeDatePicker("更新时间", "updatedAt").
		GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Select("分组", "groupId", dtosettings.AdminSettingGroupOptions).
		Select("类型", "type", dtosettings.AdminSettingTypeOptions).
		Text("名称", "name").
		Text("键名", "field").
		TextArea("键值", "value").
		TextArea("配置", "data").
		SetValue("groupId", models.AdminSettingGroupBasic).
		SetValue("type", models.AdminSettingTypeText).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", "/setting/delete", "配置批量删除").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增设置", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/setting/create", "新增设置").SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Select("分组", "groupId", true, dtosettings.AdminSettingGroupOptions).
		Select("类型", "type", true, dtosettings.AdminSettingTypeOptions).
		EditText("名称", "name", true).
		EditText("键名", "field", true).
		EditTextArea("键值", "value", true).
		EditTextArea("配置", "data", true).
		DatePicker("更新时间", "updatedAt", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Select("分组", "groupId", dtosettings.AdminSettingGroupOptions).
		Select("类型", "type", dtosettings.AdminSettingTypeOptions).
		Text("名称", "name").
		Text("键名", "field").
		TextArea("键值", "value").
		TextArea("配置", "data").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/setting/update", "更新设置").
				SetParams(updateParams).SetInputList(updateInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "详情", Color: views.ColorSecondary, Size: "xs"},
			Config: views.NewDialogViews("delete", "/setting/update/desc", "更新配置详情").
				SetParams(&views.SettingOptions{Operate: "setting", Name: "name", Type: "type", Value: "value", Params: "value", Input: "data"}),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
