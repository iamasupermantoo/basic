package bill

import (
	"basic/apis/admin/service/wallets"
	"basic/apis/common/rds"
	"basic/module/cache"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

// Views 模版JSON信息
func Views(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	settingAdminId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)
	settingAdminIds := rds.RedisFindAdminChildrenIds(settingAdminId)
	walletBillTypeList := wallets.GetWalletBillTypeListOptions(rdsConn, adminId)
	walletAssetsOptions := wallets.GetWalletAssetsOptions(settingAdminIds, []*views.InputOptions{{Label: "账户资金", Value: 0}})
	config := views.NewTableViews("/wallets/bill/index", "/wallets/bill/update")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名", "username").
		Select("资产类型", "assetsId", walletAssetsOptions).
		Select("类型", "type", walletBillTypeList).
		RangeDatePicker("提交时间", "createdAt").GetInputListInfo()

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名", "username", false).
		Select("资产类型", "assetsId", false, walletAssetsOptions).
		Text("金额", "money", true).
		Text("余额", "balance", true).
		Text("数据", "data", true).
		Select("类型", "type", true, walletBillTypeList).
		DatePicker("提交时间", "createdAt", true).
		GetColumnsListInfo()

	return c.JSON(utils.SuccessJson(config))
}
