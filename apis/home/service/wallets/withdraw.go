package wallets

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
	"fmt"
	"time"

	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

// WithdrawCreate 用户提现创建
func WithdrawCreate(adminId, userId int, lang string, params *dtowallets.WithdrawCreateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 判断充值范围
	amountBetween := &models.SettingAmountBetween{}
	amountBetweenStr := rds.RedisFindAdminSettingField(rdsConn, adminId, "walletWithdrawAmountBetween")
	_ = json.Unmarshal([]byte(amountBetweenStr), &amountBetween)
	if params.Money < amountBetween.Min || params.Money >= amountBetween.Max {
		return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "amountBetweenMismatch"))
	}

	// 获取验证吗配置
	walletSettingMap := utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateWallet"))

	// 是否需要输入安全码
	if walletSettingMap["showSecurityPass"] {
		userInfo := &models.User{}
		if result := database.Db.Model(userInfo).Where("id = ?", userId).Take(userInfo); result.Error != nil {
			return nil, result.Error
		}

		// 如果没有设置安全密钥, 提示
		if userInfo.SecurityKey == "" {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "notSecurityKey"))
		}

		// 验证安全密钥是否正确
		if userInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "securityKeyIncorrect"))
		}
	}

	//	卡片信息
	accountInfo := &models.WalletUserAccount{}
	if result := database.Db.Model(accountInfo).
		Where("id = ?", params.AccountId).
		Where("user_id = ?", userId).
		Where("status = ?", models.WalletUserAccountStatusActive).
		Take(accountInfo); result.Error != nil {
		return nil, result.Error
	}

	//	支付信息
	paymentInfo := &models.WalletPayment{}
	if result := database.Db.Model(paymentInfo).
		Where("id = ?", accountInfo.PaymentId).
		Where("mode in ?", []int{models.WalletPaymentModeWithdraw, models.WalletPaymentModeAssetsWithdraw}).
		Take(paymentInfo); result.Error != nil {
		return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "accountBeEmpty"))
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		//	获取管理提现配置信息
		withdrawSetting := &dtoadmins.AdminSettingWithdrawSetting{}
		if err := json.Unmarshal([]byte(rds.RedisFindAdminSettingField(rdsConn, adminId, "walletWithdrawSetting")), &withdrawSetting); err != nil {
			return err
		}
		nowTime := time.Now()
		sevenDaysAgoTime := nowTime.AddDate(0, 0, withdrawSetting.Days)

		// 提现数量
		var withdrawCount int64
		if result := database.Db.Model(&models.WalletOrder{}).Where("user_id = ?", userId).Where("assets_id = ?", paymentInfo.Id).
			Where("status > ?", models.WalletOrderStatusRefuse).Where("updated_at > ?", sevenDaysAgoTime.Unix()).
			Count(&withdrawCount); result.Error != nil {
			return result.Error
		}

		if int(withdrawCount) >= withdrawSetting.Nums {
			return errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "exceededNumberOfWithdrawals") + fmt.Sprintf("(%d/%d)", withdrawSetting.Nums, withdrawCount))
		}

		// 查询用户是否存在
		userInfo := &models.User{}
		if result := database.Db.Model(userInfo).
			Where("id = ?", userId).
			Take(userInfo); result.Error != nil {
			return result.Error
		}

		walletOrderType := models.WalletOrderTypeWithdraw
		if paymentInfo.Mode == models.WalletPaymentModeAssetsWithdraw {
			walletOrderType = models.WalletOrderTypeAssetsWithdraw
		}

		// 创建钱包订单
		walletOrderInfo := &models.WalletOrder{
			OrderSn:  utils.NewRandom().OrderSn(),
			AdminId:  userInfo.AdminId,
			UserId:   userId,
			Type:     walletOrderType,
			SourceId: accountInfo.Id,
			AssetsId: paymentInfo.AssetsId,
			Money:    params.Money,
			Fee:      withdrawSetting.Fee * params.Money / 100,
		}
		if result := tx.Create(walletOrderInfo); result.Error != nil {
			return result.Error
		}

		switch paymentInfo.Mode {
		case models.WalletPaymentModeWithdraw:
			// 执行用户余额消费具体逻辑
			if err := users.UserSpending(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userInfo.AdminId,
				UserId:   userInfo.Id,
				ParentId: userInfo.ParentId,
				SourceId: walletOrderInfo.Id,
				BillType: models.WalletBillTypeWithdraw,
				Balance:  userInfo.Money,
				Money:    walletOrderInfo.Money,
			}); err != nil {
				return errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, err.Error()))
			}

		case models.WalletPaymentModeAssetsWithdraw:
			// 执行用户资金消费具体逻辑
			if err := users.UserAssetsSpending(tx, &dto.WalletOrderAgreeParams{
				AssetId:  paymentInfo.AssetsId,
				AdminId:  userInfo.AdminId,
				UserId:   userInfo.Id,
				ParentId: userInfo.ParentId,
				SourceId: walletOrderInfo.Id,
				BillType: models.WalletBillTypeAssetsWithdraw,
				Money:    walletOrderInfo.Money,
			}); err != nil {
				return errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, err.Error()))
			}
		}

		return nil
	})

	return "ok", err
}
