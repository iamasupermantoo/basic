package wallets

import (
	"basic/apis/common/rds"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
	"strconv"
)

// AccountIndex 卡片列表
func AccountIndex(adminId, userId int, acceptLanguage string, params *dtowallets.HomeAccountIndexParams) (interface{}, error) {
	accountList := make([]*dtowallets.HomeAccountInfo, 0)
	if result := database.Db.Table("wallet_user_account as wua").
		Select("wua.id", "wp1.name", "wp.name AS PaymentName", "wp.type AS PaymentType", "wp.id AS PaymentId", "wp.icon", "wua.real_name", "wua.number", "wua.code", "wua.status", "wua.created_at").
		Joins("left join wallet_payment as wp on wp.id = wua.payment_id AND wp.mode IN ?", params.Modes).
		Joins("inner join wallet_payment as wp1 on wp.type = wp1.type and wp.mode = wp1.mode and wp1.admin_id = 0").
		Where("wua.user_id = ?", userId).Where("").
		Where("wua.status > ?", models.WalletUserAccountStatusDelete).
		Find(&accountList); result.Error != nil {
		return nil, result.Error
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	for _, v := range accountList {
		v.Name = rds.RedisFindTranslateField(rdsConn, adminId, acceptLanguage, v.Name)
	}
	return accountList, nil
}

// AccountUpdate 编辑卡片
func AccountUpdate(adminId, userId int, acceptLanguage string, params *dtowallets.AccountUpdateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

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
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, acceptLanguage, "notSecurityKey"))
		}

		// 验证安全密钥是否正确
		if params.SecurityKey == "" || userInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, acceptLanguage, "securityKeyIncorrect"))
		}
	}

	if result := database.Db.Model(&models.WalletUserAccount{}).
		Where("user_id = ?", userId).
		Where("id = ?", params.Id).
		Where("status > ?", models.WalletUserAccountStatusDelete).
		Updates(params); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// AccountCreate 新增卡片
func AccountCreate(adminId, userId int, params *dtowallets.AccountCreateParams) (interface{}, error) {
	// 获取卡片数量设置，判断卡片是否超出添加限制
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	walletAccountNums := rds.RedisFindAdminSettingField(rdsConn, adminId, "walletAccountNums")
	cardLimit, err := strconv.ParseInt(walletAccountNums, 10, 64)
	if err != nil {
		return nil, err
	}

	var currentWalletAccountNums int64
	if result := database.Db.Model(&models.WalletUserAccount{}).
		Where("user_id = ?", userId).
		Where("payment_id = ?", params.PaymentId).
		Where("status > ?", models.WalletUserAccountStatusDelete).
		Count(&currentWalletAccountNums); result.Error != nil {
		return nil, result.Error
	}

	// 判断卡片是否超出添加限制
	if currentWalletAccountNums >= cardLimit {
		return nil, errors.New("cardAdditionLimitExceeded")
	}

	// 创建卡片
	if result := database.Db.Where("user_id = ?", userId).Create(&models.WalletUserAccount{
		AdminId:   adminId,
		UserId:    userId,
		PaymentId: params.PaymentId,
		RealName:  params.RealName,
		Number:    params.Number,
		Code:      params.Code,
	}); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// AccountDelete 删除卡片
func AccountDelete(adminId, userId int, lang string, params *dtowallets.HomeAccountDeleteParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 获取验证吗配置
	walletSettingMap := utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateWallet"))

	// 是否需要输入安全码
	if walletSettingMap["showSecurityPass"] {
		userInfo := &models.User{}
		if result := database.Db.Model(userInfo).Where("id = ?", userId).Take(userInfo); result.Error != nil {
			return nil, result.Error
		}

		// 验证安全密钥是否正确
		if userInfo.SecurityKey != utils.PasswordEncrypt(params.SecurityKey) {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "securityKeyIncorrect"))
		}
	}

	if result := database.Db.Model(&models.WalletUserAccount{}).
		Where("user_id = ?", userId).
		Where("id = ?", params.Id).
		Update("status", models.WalletUserAccountStatusDelete); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// AccountInfo 卡片详情
func AccountInfo(adminId, userId int, lang string, params *dtowallets.AccountInfoParams) (interface{}, error) {
	accountInfo := &dtowallets.HomeAccountInfo{}
	if result := database.Db.Table("wallet_user_account as wua").
		Select("wua.id", "wp1.name", "wp.name as PaymentName", "wp.id AS PaymentId", "wp.icon", "wua.real_name", "wua.number", "wua.code", "wua.status", "wua.created_at").
		Joins("left join wallet_payment as wp on wp.id = wua.payment_id").
		Joins("inner join wallet_payment as wp1 on wp.type = wp1.type and wp.mode = wp1.mode and wp1.admin_id = 0").
		Where("wua.id = ?", params.Id).
		Where("wua.user_id = ?", userId).
		Where("wua.status > ?", models.WalletUserAccountStatusDelete).
		Find(accountInfo); result.Error != nil {
		return nil, result.Error
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()
	accountInfo.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, accountInfo.Name)

	return accountInfo, nil
}
