package payment

import (
	"basic/apis/admin/service/wallets"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)
	options := []*views.InputOptions{{Label: "账户资金", Value: 0}}
	if adminId == models.SuperAdminId {
		options = append(options, &views.InputOptions{Label: "提现类型", Value: -1})
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	settingAdminId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)
	settingAdminIds := rds.RedisFindAdminChildrenIds(settingAdminId)
	walletAssetsOptions := wallets.GetWalletAssetsOptions(settingAdminIds, options)
	config := views.NewTableViews("wallets/payment/index", "wallets/payment/update")

	// 查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("名称", "name").
		Select("资产", "assetsId", walletAssetsOptions).
		Select("模式", "mode", wallets.WalletPaymentModeOptions).
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.WalletPaymentStatusDisable}, {Label: "激活", Value: models.WalletPaymentStatusActive}}).
		Select("类型", "type", wallets.WalletPaymentTypeOptions).
		Text("数据", "data").
		Text("详情", "desc").
		RangeDatePicker("提交时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Image("图标", "icon").
		Text("名称", "name").
		Select("资产", "assetsId", walletAssetsOptions).
		Select("模式", "mode", wallets.WalletPaymentModeOptions).
		Select("类型", "type", wallets.WalletPaymentTypeOptions).
		SetValue("mode", models.WalletPaymentModeWithdraw).
		SetValue("type", models.WalletPaymentTypeDigital).
		SetValue("assetsId", walletAssetsOptions[0].Value).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/wallets/payment/delete", "删除支付").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增支付", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/wallets/payment/create", "新增支付").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理", "adminName", false).
		Image("图标", "icon", true).
		Translate("名称", "name", true).
		Select("资产", "assetsId", true, walletAssetsOptions).
		Select("模式", "mode", true, wallets.WalletPaymentModeOptions).
		Select("类型", "type", true, wallets.WalletPaymentTypeOptions).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "激活", Value: models.WalletPaymentStatusActive}, {Label: "禁用", Value: models.WalletPaymentStatusDisable}}).
		DatePicker("操作时间", "updatedAt", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Image("图标", "icon").
		Text("名称", "name").
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.WalletPaymentStatusDisable}, {Label: "激活", Value: models.WalletPaymentStatusActive}}).
		Text("数据", "data").
		TextArea("详情", "desc").
		GetInputListInfo()

	dataJsonParam, dataJsonInputList := views.NewInputViews().
		Text("名称", "name").
		Text("姓名|网络", "realName").
		Text("卡号|地址", "number").
		Text("识别号", "code").GetInputListInfo()
	optionsParam, optionsInputList := views.NewInputViews().
		Json("", "dataJson", dataJsonInputList).SetValue("dataJson", dataJsonParam).
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config:      views.NewDialogViews("update", "/wallets/payment/update", "更新支付信息").SetSizingSmall().SetParams(updateParams).SetInputList(updateInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "信息", Color: views.ColorSecondary, Size: "xs", Eval: "scope.row.adminId > 0 && (scope.row.mode==1 || scope.row.mode==2)"},
			Config: views.NewDialogViews("setting", "/wallets/payment/desc", "配置管理").
				SetParams(optionsParam).SetInputList(optionsInputList).
				SetSizingSmall(),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
