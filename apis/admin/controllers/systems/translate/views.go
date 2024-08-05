package translate

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/system/translate/index", "/system/translate/update")
	//	查询一种翻译方式
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理员名称", "adminName").
		Text("名称", "name").
		Select("类型", "type", []*views.InputOptions{{Label: "系统翻译", Value: models.TranslateTypeDefault}, {Label: "前台翻译", Value: models.TranslateTypeHome}}).
		RangeDatePicker("创建时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Text("名称", "name").
		Select("类型", "type", []*views.InputOptions{{Label: "系统翻译", Value: models.TranslateTypeDefault}, {Label: "前台翻译", Value: models.TranslateTypeHome}}).
		Text("键名", "field").
		Text("键值", "value").
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/system/translate/delete", "批量删除").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/system/translate/create", "新增").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理员名称", "adminName", false).
		Text("名称", "name", false).
		EditText("语言标识", "lang", true).
		Select("类型", "type", true, []*views.InputOptions{{Label: "系统翻译", Value: models.TranslateTypeDefault}, {Label: "前台翻译", Value: models.TranslateTypeHome}}).
		EditText("键名", "field", false).
		EditTextArea("键值", "value", false).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	更新参数
	updateParams, updateInputList := views.NewInputViews().
		Text("名称", "name").
		Select("类型", "type", []*views.InputOptions{{Label: "系统翻译", Value: models.TranslateTypeDefault}, {Label: "前台翻译", Value: models.TranslateTypeHome}}).
		Text("键名", "field").
		Editor("键值", "value").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/system/translate/update", "更新信息").
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
