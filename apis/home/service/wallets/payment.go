package wallets

import (
	"basic/apis/common/rds"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"github.com/goccy/go-json"
)

// PaymentIndex 支付列表
func PaymentIndex(adminId int, lang string, params *dtowallets.PaymentIndexParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	获取商户管理ID
	if len(params.Modes) == 0 {
		params.Modes = []int{models.WalletPaymentModeWithdraw, models.WalletPaymentModeAssetsWithdraw}
	}
	adminSettingId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)
	filterParams := utils.NewFilterParams().
		InInt("mode IN ?", params.Modes)

	walletPaymentList := make([]*dtowallets.PaymentInfo, 0)
	if result := database.Db.
		Table("wallet_payment").
		Where("admin_id IN ?", []int{0, adminSettingId}).
		Scopes(filterParams.Scopes()).
		Order("admin_id ASC").
		Find(&walletPaymentList); result.Error != nil {
		return nil, result.Error
	}

	data := make([]*dtowallets.HomePaymentIndexData, 0)
	for _, v := range walletPaymentList {
		//	翻译名称
		v.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, v.Name)

		//	如果是管理信息, 并且是充值方式
		if v.AdminId > 0 && (v.Mode == models.WalletPaymentModeDeposit || v.Mode == models.WalletPaymentModeAssetsDeposit) {
			v.DataJson = &models.WalletPaymentData{}
			_ = json.Unmarshal([]byte(v.Data), &v.DataJson)
		}

		//	顶级列表
		if v.AdminId == 0 {
			data = append(data, &dtowallets.HomePaymentIndexData{
				PaymentInfo: *v,
				Items:       make([]*dtowallets.HomePaymentIndexData, 0),
			})
			continue
		}

		for _, item := range data {
			if item.Type == v.Type && item.Mode == v.Mode {
				item.Items = append(item.Items, &dtowallets.HomePaymentIndexData{
					PaymentInfo: *v,
					Items:       make([]*dtowallets.HomePaymentIndexData, 0),
				})
			}
		}
	}

	return data, nil
}
