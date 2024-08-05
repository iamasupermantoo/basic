package wallets

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
)

// AccountIndex 用户账户列表
func AccountIndex(adminId, userId int, params *dtowallets.AccountIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtowallets.AccountIndexData, 0)}

	// 过滤where 条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("real_name LIKE ?", params.RealName+"%").
		Like("number LIKE ?", params.Number+"%").
		Like("code LIKE ?", params.Code).
		EqInt("status = ?", params.Status).
		EqInt("payment_id = ?", params.PaymentId).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	result := database.Db.Table("wallet_user_account wua").
		Select("wua.id", "wua.admin_id", "au.username AS adminName", "wua.user_id", "u.username", "wua.payment_id", "wp.name AS paymentName", "wua.name", "wua.real_name", "wua.number", "wua.code", "wua.status", "wua.data", "wua.created_at", "wua.updated_at").
		Joins("LEFT JOIN admin_user AS au ON wua.admin_id = au.id").
		Joins("LEFT JOIN user AS u ON u.id = wua.user_id").
		Joins("LEFT JOIN wallet_payment AS wp ON wp.id = wua.payment_id").
		Where("wua.status > ?", models.WalletUserAccountStatusDelete).
		Where("wua.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Scopes(utils.Paginate(params.Pagination)).
		Count(&data.Count).
		Scan(&data.Items)

	return data, result.Error
}

// AccountCreate 用户账户新增
func AccountCreate(adminId, userId int, params *dtowallets.AccountCreateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	settingAdminId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)
	paymentInfo := &models.WalletPayment{}
	if result := database.Db.Model(paymentInfo).
		Where("id = ?", params.PaymentId).Where("mode IN ?", []int{models.WalletPaymentModeWithdraw, models.WalletPaymentModeAssetsWithdraw}).
		Where("admin_id = ?", settingAdminId).
		Take(paymentInfo); result.Error != nil {
		return nil, result.Error
	}

	// 判断用户是否存在
	userInfo := &models.User{}
	if result := database.Db.Model(userInfo).
		Where("username = ?", params.UserName).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Take(userInfo); result.Error != nil {
		return nil, result.Error
	}
	if rds.GetAdminSettingId(userInfo.AdminId) != paymentInfo.AdminId {
		return nil, errors.New("incorrectFormat")
	}

	if result := database.Db.Create(&models.WalletUserAccount{
		AdminId:   userInfo.AdminId,
		UserId:    userInfo.Id,
		PaymentId: params.PaymentId,
		RealName:  params.RealName,
		Number:    params.Number,
		Code:      params.Code,
	}); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// AccountUpdate 用户账户更新
func AccountUpdate(adminId, userId int, params *dtowallets.AccountUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletUserAccount{}).
		Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Updates(params); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// AccountDelete 用户账户删除
func AccountDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletUserAccount{}).
		Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.WalletAssetsStatusDelete); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}
