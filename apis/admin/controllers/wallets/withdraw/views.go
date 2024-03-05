package withdraw

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
	walletPaymentOptions := wallets.GetWalletPaymentOptions(settingAdminIds, []int{models.WalletOrderTypeWithdraw, models.WalletOrderTypeAssetsWithdraw})
	if len(walletPaymentOptions) == 0 {
		return c.JSON(utils.ErrorJson(errors.New("没有提现方式")))
	}

	config := views.NewTableViews("/wallets/withdraw/index", "/wallets/withdraw/update")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名称", "username").
		Text("拒绝理由", "data").
		Text("真实姓名", "realName").
		Text("银行卡号｜地址", "number").
		Text("银行代码", "code").
		Select("类型", "type", []*views.InputOptions{{Label: "余额提现", Value: models.WalletOrderTypeWithdraw}, {Label: "资产提现", Value: models.WalletOrderTypeAssetsWithdraw}}).
		Select("状态", "status", []*views.InputOptions{{Label: "拒绝", Value: models.WalletOrderStatusRefuse}, {Label: "审核", Value: models.WalletOrderStatusActive}, {Label: "同意", Value: models.WalletOrderStatusComplete}}).
		RangeDatePicker("提交时间", "createdAt").GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Text("用户名称", "username").
		Number("金额", "money").
		Select("提现方式", "paymentId", walletPaymentOptions).
		SetValue("paymentId", walletPaymentOptions[0].Value).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/wallets/withdraw/delete", "用户提现删除").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "orderSn", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增提现", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/wallets/withdraw/create", "新增用户提现").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名称", "username", false).
		Select("提现类型", "type", true, []*views.InputOptions{{Label: "余额提现", Value: models.WalletOrderTypeWithdraw}, {Label: "资产提现", Value: models.WalletOrderTypeAssetsWithdraw}}).
		Text("订单编号", "orderSn", true).
		Text("金额", "money", true).
		Text("真实姓名", "reaName", true).
		Text("卡号", "number", true).
		Text("银行代码", "code", true).
		EditTextArea("拒绝理由", "data", true).
		Select("状态", "status", true, []*views.InputOptions{{Label: "同意", Value: models.WalletOrderStatusComplete}, {Label: "拒绝", Value: models.WalletOrderStatusRefuse}, {Label: "审核", Value: models.WalletOrderStatusActive}}).
		DatePicker("更新时间", "updatedAt", true).
		DatePicker("申请时间", "createdAt", true).
		GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Select("状态", "status", []*views.InputOptions{{Label: "拒绝", Value: models.WalletOrderStatusRefuse}, {Label: "同意", Value: models.WalletOrderStatusComplete}}).
		TextArea("拒绝理由", "data").
		SetValue("status", models.WalletOrderStatusComplete).
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "审核", Color: views.ColorPrimary, Size: "xs", Eval: "scope.row.status == 10"},
			Config:      views.NewDialogViews("status", "/wallets/withdraw/status", "审核订单").SetSizingSmall().SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
