package assets

import (
	"basic/apis/admin/service/wallets"
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/wallets/assets/index", "/wallets/assets/update")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("名称", "name").
		Select("类型", "type", wallets.WalletAssetsTypeOptions).
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.WalletAssetsStatusDisable}, {Label: "激活", Value: models.WalletAssetsStatusActive}}).
		RangeDatePicker("激活时间", "createdAt").GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Image("图标", "icon").
		Text("名称", "name").
		Number("汇率", "rate").
		Select("类型", "type", wallets.WalletAssetsTypeOptions).
		SetValue("type", models.WalletAssetsTypeDigitalCurrency).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/wallets/assets/delete", "资产删除").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增资产", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/wallets/assets/create", "新增资产").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Translate("名称", "name", true).
		Image("图标", "icon", true).
		Text("汇率", "rate", true).
		Select("类型", "type", true, wallets.WalletAssetsTypeOptions).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "激活", Value: models.WalletAssetsStatusActive}, {Label: "禁用", Value: models.WalletAssetsStatusDisable}}).
		EditTextArea("数据", "data", true).
		DatePicker("更新时间", "updatedAt", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Image("图标", "icon").
		Text("名称", "name").
		Number("汇率", "rate").
		Select("类型", "type", wallets.WalletAssetsTypeOptions).
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.WalletAssetsStatusDisable}, {Label: "激活", Value: models.WalletAssetsStatusActive}}).
		Text("数据", "data").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config:      views.NewDialogViews("update", "/wallets/assets/update", "更新资产信息").SetSizingSmall().SetParams(updateParams).SetInputList(updateInputList),
		},
	}

	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)
	return c.JSON(utils.SuccessJson(config))
}
