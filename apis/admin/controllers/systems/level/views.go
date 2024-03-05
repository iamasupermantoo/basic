package level

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/system/level/index", "/system/level/update")

	//	查询单个等级
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("名称", "name").
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.LevelStatusActive}, {Label: "禁用", Value: models.LevelStatusDisable}}).
		RangeDatePicker("创建时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Image("图标", "icon").
		Number("等级标识", "level").
		Text("名称", "name").
		Number("金额", "money").
		Number("天数", "days").
		Editor("详情", "desc").
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/system/level/delete", "批量删除等级").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增等级", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/system/level/create", "新增等级").SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有等级数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		EditNumber("等级标识", "level", true).
		Translate("名称", "name", true).
		Image("图标", "icon", true).
		Text("金额", "money", true).
		EditNumber("天数", "days", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "开启", Value: models.LevelStatusActive}, {Label: "禁用", Value: models.LevelStatusDisable}}).
		EditTextArea("数据", "data", true).
		Translate("详情", "desc", false).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	更新等级数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Image("图标", "icon").
		Number("等级标识", "level").
		Text("名称", "name").
		Number("天数", "days").
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.LevelStatusActive}, {Label: "禁用", Value: models.LevelStatusDisable}}).
		TextArea("数据", "data").
		Editor("详情", "desc").
		GetInputListInfo()
	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/system/level/update", "更新等级信息").
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)
	return c.JSON(utils.SuccessJson(config))
}
