package lang

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/system/lang/index", "/system/lang/update")
	config.Pagination = &utils.Pagination{SortBy: "sort", Descending: false, Page: 1, RowsPerPage: 20}

	//	查询单种语言
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("语言名称", "name").
		Text("语言别名", "alias").
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.LangStatusActive}, {Label: "禁用", Value: models.LangStatusDisable}}).
		RangeDatePicker("创建时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Image("语言图标", "icon").
		Text("语言名称", "name").
		Text("语言别名", "alias").
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/system/lang/delete", "批量删除语言").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增语言", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/system/lang/create", "新增语言").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有语言数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		EditText("语言名称", "name", true).
		EditText("语言别名", "alias", true).
		Image("图标", "icon", true).
		EditNumber("排序", "sort", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "开启", Value: models.LangStatusActive}, {Label: "禁用", Value: models.LangStatusDisable}}).
		EditTextArea("数据", "data", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	更新语言数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Image("语言图标", "icon").
		Text("语言名称", "name").
		Text("语言别名", "alias").
		Number("排序", "sort").
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.LangStatusActive}, {Label: "禁用", Value: models.LangStatusDisable}}).
		TextArea("数据", "data").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/system/lang/update", "更新语言信息").
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
