package country

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/system/country/index", "/system/country/update")

	//	查询单个国家
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("国家名称", "name").
		Text("国家别名", "alias").
		Text("二位代码", "iso1").
		Text("区号", "code").
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.CountryStatusActive}, {Label: "禁用", Value: models.CountryStatusDisable}}).
		RangeDatePicker("创建时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Image("国家图标", "icon").
		Text("国家名称", "name").
		Text("国家别名", "alias").
		Text("二位代码", "iso1").
		Text("国家区号", "code").
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/system/country/delete", "批量删除国家").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增国家", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/system/country/create", "新增国家").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有国家数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		EditText("国家名字", "name", true).
		EditText("国家别名", "alias", true).
		Image("图标", "icon", true).
		EditText("二位代码", "iso1", true).
		EditNumber("排序", "sort", true).
		EditText("区号", "code", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "开启", Value: models.CountryStatusActive}, {Label: "禁用", Value: models.CountryStatusDisable}}).
		EditTextArea("数据", "data", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	更新国家数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Image("国家图标", "icon").
		Text("国家名字", "name").
		Text("国家别名", "alias").
		Text("二位代码", "iso1").
		Number("排序", "sort").
		Text("区号", "code").
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.CountryStatusActive}, {Label: "禁用", Value: models.CountryStatusDisable}}).
		TextArea("数据", "data").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/system/country/update", "更新国家信息").
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
