package wallets

import (
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils/views"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// WalletPaymentTypeOptions 支付类型Options
var WalletPaymentTypeOptions = []*views.InputOptions{
	{Label: "银行卡", Value: models.WalletPaymentTypeBank},
	{Label: "数字货币", Value: models.WalletPaymentTypeDigital},
	{Label: "第三方支付", Value: models.WalletPaymentTypeThree},
}

// WalletPaymentModeOptions 支付模式Options
var WalletPaymentModeOptions = []*views.InputOptions{
	{Label: "余额充值", Value: models.WalletPaymentModeDeposit},
	{Label: "资产充值", Value: models.WalletPaymentModeAssetsDeposit},
	{Label: "余额提现", Value: models.WalletPaymentModeWithdraw},
	{Label: "资产提现", Value: models.WalletPaymentModeAssetsWithdraw},
}

// WalletAssetsTypeOptions 支付模式Options
var WalletAssetsTypeOptions = []*views.InputOptions{
	{Label: "法币资产", Value: models.WalletAssetsTypeFiatCurrency},
	{Label: "数字货币资产", Value: models.WalletAssetsTypeDigitalCurrency},
	{Label: "虚拟货币资产", Value: models.WalletAssetsTypeVirtualCurrency},
}

// GetWalletAssetsTypeName 获取资产类型中文名
func GetWalletAssetsTypeName(assetsTypeNumber int) string {
	for _, assetsType := range WalletAssetsTypeOptions {
		if assetsType.Value == assetsTypeNumber {
			return assetsType.Label
		}
	}
	return ""
}

// GetWalletPaymentOptions 获取支付方式Options
func GetWalletPaymentOptions(settingAdminIds []int, walletOrderModeList []int) []*views.InputOptions {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	data := make([]*views.InputOptions, 0)
	walletPaymentList := make([]*models.WalletPayment, 0)
	result := database.Db.Model(&models.WalletPayment{}).Where("admin_id IN ?", settingAdminIds).Where("admin_id > 0")
	if len(walletOrderModeList) != 0 {
		result.Where("mode IN ?", walletOrderModeList)
	}
	if result.Find(&walletPaymentList); result.Error != nil {
		panic(result.Error)
	}

	for i := 0; i < len(walletPaymentList); i++ {
		data = append(data, &views.InputOptions{
			Label: rds.RedisFindTranslateField(rdsConn, settingAdminIds[0], "zh-CN", walletPaymentList[i].Name) + "【管理ID: " + strconv.Itoa(walletPaymentList[i].AdminId) + " | " + models.WalletPaymentTypeList[walletPaymentList[i].Type] + "." + models.WalletPaymentModeList[walletPaymentList[i].Mode] + "】",
			Value: walletPaymentList[i].Id,
		})
	}

	return data
}

// GetWalletAssetsOptions 获取钱包资产Options
func GetWalletAssetsOptions(adminIds []int, options []*views.InputOptions) []*views.InputOptions {
	data := make([]*views.InputOptions, 0)
	if options != nil {
		data = append(data, options...)
	}

	var walletAssetsList []*models.WalletAssets
	if result := database.Db.Model(&models.WalletAssets{}).Where("admin_id IN ?", adminIds).Find(&walletAssetsList); result.Error != nil {
		return nil
	}

	for i := 0; i < len(walletAssetsList); i++ {
		data = append(data, &views.InputOptions{
			Label: walletAssetsList[i].Name + "" + GetWalletAssetsTypeName(walletAssetsList[i].Type) + "【 管理ID: " + strconv.Itoa(walletAssetsList[i].AdminId) + " 】",
			Value: walletAssetsList[i].Id,
		})
	}

	return data
}

// GetWalletBillTypeListOptions 获取账单列表
func GetWalletBillTypeListOptions(rdsConn redis.Conn, adminId int) []*views.InputOptions {
	billList := models.GetBillTypeList()
	optionsList := make([]*views.InputOptions, 0)
	for k, v := range billList {
		optionsList = append(optionsList, &views.InputOptions{
			Label: rds.RedisFindTranslateField(rdsConn, adminId, "zh-CN", v),
			Value: k,
		})
	}
	return optionsList
}
