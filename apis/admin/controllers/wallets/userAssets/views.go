package userassets

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
	adminId, _ := utils.GetContextClaims(c)
	config := views.NewTableViews("/user/assets/index", "/user/assets/update")

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	settingAdminId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)
	settingAdminIds := rds.RedisFindAdminChildrenIds(settingAdminId)
	assetsList := wallets.GetWalletAssetsOptions(settingAdminIds, nil)
	if len(assetsList) == 0 {
		return c.JSON(utils.ErrorJson(errors.New("没有用户资产列表")))
	}

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名", "username").
		Select("资产类型", "walletAssetsId", assetsList).
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.WalletUserAssetsStatusDisable}, {Label: "激活", Value: models.WalletUserAssetsStatusActive}}).
		RangeDatePicker("激活时间", "createdAt").GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Text("用户名", "username").
		Select("资产类型", "walletAssetsId", assetsList).
		Select("类型", "type", []*views.InputOptions{{Label: "加款", Value: models.WalletBillTypeSystemAssetsDeposit}, {Label: "减款", Value: models.WalletBillTypeSystemAssetsWithdraw}}).
		Number("金额", "money").
		SetValue("type", models.WalletBillTypeSystemAssetsDeposit).
		SetValue("walletAssetsId", assetsList[0].Value).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/user/assets/delete", "用户充值删除").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "id", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "资产加减款", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("amount", "/user/assets/balance", "用户资产加减款").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名", "username", false).
		Select("资产类型", "walletAssetsId", true, assetsList).
		Text("金额", "money", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "启用", Value: models.WalletUserAssetsStatusActive}, {Label: "禁用", Value: models.WalletUserAssetsStatusDisable}}).
		DatePicker("最近时间", "updatedAt", true).
		DatePicker("激活时间", "createdAt", true).
		GetColumnsListInfo()
	return c.JSON(utils.SuccessJson(config))
}
