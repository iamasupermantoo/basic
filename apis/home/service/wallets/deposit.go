package wallets

import (
	"basic/apis/common/rds"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
	"github.com/goccy/go-json"
)

// DepositCreate 用户充值新增
func DepositCreate(adminId, userId int, lang string, params *dtowallets.DepositCreateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 判断充值范围
	amountBetween := &models.SettingAmountBetween{}
	amountBetweenStr := rds.RedisFindAdminSettingField(rdsConn, adminId, "walletDepositAmountBetween")
	_ = json.Unmarshal([]byte(amountBetweenStr), &amountBetween)
	if params.Money < amountBetween.Min || params.Money >= amountBetween.Max {
		return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "amountBetweenMismatch"))
	}

	// 判断用户是否存在
	userInfo := &models.User{}
	if result := database.Db.Model(userInfo).Where("id = ?", userId).Take(userInfo); result.Error != nil {
		return nil, result.Error
	}

	// 判断是否有充值方式
	walletPaymentInfo := &models.WalletPayment{}
	if result := database.Db.Model(walletPaymentInfo).Where("id = ?", params.PaymentId).
		Where("mode in ?", []int{models.WalletPaymentModeDeposit, models.WalletPaymentModeAssetsDeposit}).
		Take(walletPaymentInfo); result.Error != nil {
		return nil, result.Error
	}

	// 插入订单数据
	walletOrderType := models.WalletOrderTypeDeposit
	if walletPaymentInfo.Mode == models.WalletPaymentModeAssetsDeposit {
		walletOrderType = models.WalletOrderTypeAssetsDeposit
	}

	if result := database.Db.Create(&models.WalletOrder{
		OrderSn:  utils.NewRandom().OrderSn(),
		AdminId:  userInfo.AdminId,
		UserId:   userInfo.Id,
		Type:     walletOrderType,
		SourceId: walletPaymentInfo.Id,
		AssetsId: walletPaymentInfo.AssetsId,
		Voucher:  params.Voucher,
		Money:    params.Money,
	}); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}
