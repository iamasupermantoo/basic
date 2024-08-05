package account

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

	config := views.NewTableViews("/wallets/account/index", "/wallets/account/update")

	// 获取支付方式
	adminId, _ := utils.GetContextClaims(c)
	settingAdminId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)
	settingAdminIds := rds.RedisFindAdminChildrenIds(settingAdminId)
	paymentOptions := wallets.GetWalletPaymentOptions(settingAdminIds, []int{models.WalletPaymentModeWithdraw, models.WalletPaymentModeAssetsWithdraw})
	if len(paymentOptions) == 0 {
		return c.JSON(utils.ErrorJson(errors.New("没有设置提现类型")))
	}

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名", "username").
		Text("备注", "name").
		Text("姓名", "realName").
		Text("卡号", "number").
		Text("银行代码", "code").
		Select("支付方式", "paymentId", paymentOptions).
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.WalletUserAccountStatusDisable}, {Label: "激活", Value: models.WalletUserAccountStatusActive}}).
		RangeDatePicker("创建时间", "createdAt").GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Select("提现方式", "paymentId", paymentOptions).
		Text("用户名", "username").
		Text("真实姓名", "realName").
		Text("卡号｜地址", "number").
		Text("银行代码", "code").
		SetValue("paymentId", paymentOptions[0].Value).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/wallets/account/delete", "用户提现账户删除").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增卡片", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/wallets/account/create", "用户提现账户新增").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理员", "adminName", false).
		Text("用户名", "username", false).
		EditText("备注", "name", true).
		Select("支付方式", "paymentId", true, paymentOptions).
		EditText("姓名", "realName", true).
		EditText("卡号|地址", "number", true).
		EditText("银行代码", "code", true).
		EditTextArea("数据", "data", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "激活", Value: models.WalletUserAccountStatusActive}, {Label: "禁用", Value: models.WalletUserAccountStatusDisable}}).
		DatePicker("更新时间", "updatedAt", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Text("备注", "name").
		Text("姓名", "realName").
		Text("卡号｜地址", "number").
		Text("银行代码", "code").
		Text("数据", "data").
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.WalletUserAccountStatusDisable}, {Label: "激活", Value: models.WalletUserAccountStatusActive}}).
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config:      views.NewDialogViews("update", "/wallets/account/update", "更新卡片信息").SetSizingSmall().SetParams(updateParams).SetInputList(updateInputList),
		},
	}

	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
