package deposit

import (
	"basic/apis/admin/service/wallets"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/utils"
	"basic/utils/views"
	"errors"

	"github.com/gofiber/fiber/v2"
)

// Views 模版JSON信息
func Views(c *fiber.Ctx) error {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	adminId, _ := utils.GetContextClaims(c)
	settingAdminId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)
	settingAdminIds := rds.RedisFindAdminChildrenIds(settingAdminId)
	walletPaymentOptions := wallets.GetWalletPaymentOptions(settingAdminIds, []int{models.WalletPaymentModeDeposit, models.WalletPaymentModeAssetsDeposit})
	if len(walletPaymentOptions) == 0 {
		return c.JSON(utils.ErrorJson(errors.New("没有充值设置")))
	}

	config := views.NewTableViews("/wallets/deposit/index", "/wallets/deposit/update")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名", "username").
		Text("拒绝理由", "data").
		Select("充值方式", "sourceId", walletPaymentOptions).
		Select("类型", "type", []*views.InputOptions{{Label: "余额充值", Value: models.WalletOrderTypeDeposit}, {Label: "资产充值", Value: models.WalletOrderTypeAssetsDeposit}}).
		Select("状态", "status", []*views.InputOptions{{Label: "拒绝", Value: models.WalletOrderStatusRefuse}, {Label: "审核", Value: models.WalletOrderStatusActive}, {Label: "同意", Value: models.WalletOrderStatusComplete}}).
		RangeDatePicker("提交时间", "createdAt").GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Text("用户名", "username").
		Number("金额", "money").
		Select("充值方式", "paymentId", walletPaymentOptions).
		SetValue("paymentId", walletPaymentOptions[0].Value).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/wallets/deposit/delete", "用户充值删除").SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "orderSn", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增订单", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/wallets/deposit/create", "新增平台订单").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名", "username", false).
		Select("充值方式", "sourceId", false, walletPaymentOptions).
		Text("订单编号", "orderSn", true).
		Text("金额", "money", true).
		Select("类型", "type", true, []*views.InputOptions{{Label: "余额充值", Value: models.WalletOrderTypeDeposit}, {Label: "资产充值", Value: models.WalletOrderTypeAssetsDeposit}}).
		Select("状态", "status", true, []*views.InputOptions{{Label: "同意", Value: models.WalletOrderStatusComplete}, {Label: "拒绝", Value: models.WalletOrderStatusRefuse}, {Label: "审核", Value: models.WalletOrderStatusActive}}).
		Image("支付凭证", "voucher", true).
		EditTextArea("拒绝理由", "data", true).
		DatePicker("操作时间", "updatedAt", true).
		DatePicker("提交时间", "createdAt", true).
		GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Select("操作", "status", []*views.InputOptions{{Label: "拒绝", Value: models.WalletOrderStatusRefuse}, {Label: "同意", Value: models.WalletOrderStatusComplete}}).
		TextArea("拒绝理由", "data").
		SetValue("status", models.WalletOrderStatusComplete).
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "审核", Color: views.ColorPrimary, Size: "xs", Eval: "scope.row.status == 10"},
			Config:      views.NewDialogViews("status", "/wallets/deposit/status", "审核订单").SetSizingSmall().SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
