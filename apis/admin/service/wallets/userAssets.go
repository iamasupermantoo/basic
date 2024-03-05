package wallets

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"errors"
	"gorm.io/gorm"
)

// UserAssetsIndex 用户资产查询
func UserAssetsIndex(adminId, userId int, params *dtowallets.UserAssetsIndexParams) (interface{}, error) {
	// 过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		EqInt("wallet_assets_id = ?", params.WalletAssetsId).
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	// 获取查询数据
	data := dto.IndexData{Items: make([]*dtowallets.UserAssetsIndexData, 0)}
	if result := database.Db.Table("wallet_user_assets AS wua").
		Select("wua.id", "wua.admin_id", "wua.user_id", "wua.wallet_assets_id", "wa.name AS walletAssetsName", "au.username AS adminName", "u.username AS userName", "wua.status", "wua.money", "wua.data", "wua.updated_at", "wua.created_at").
		Joins("left join user AS u ON wua.user_id = u.id").
		Joins("left join admin_user AS au ON wua.admin_id = au.id").
		Joins("left join wallet_assets AS wa ON wua.wallet_assets_id = wa.id").
		Where("wua.status > ?", models.WalletUserAssetsStatusDelete).
		Where("wua.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Scopes(utils.Paginate(params.Pagination)).
		Count(&data.Count).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

// UserAssetsDelete 用户资产删除
func UserAssetsDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletUserAssets{}).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id IN ?", params.Ids).
		Update("status", models.WalletUserAssetsStatusDelete); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// UserAssetsUpdate 用户资产更新
func UserAssetsUpdate(adminId, userId int, params *dtowallets.UserAssetsUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletUserAssets{}).
		Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("status > ?", models.WalletUserAccountStatusDelete).
		Update("status", params.Status); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// UserAssetsBalance 用户资产余额变动
func UserAssetsBalance(adminId, userId int, params *dtowallets.UserAssetsBalanceParams) (interface{}, error) {
	//	查询用户信息
	userInfo := models.User{}
	if result := database.Db.Model(&userInfo).
		Where("username = ?", params.UserName).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Take(&userInfo); result.Error != nil {
		return nil, result.Error
	}

	//	资产信息
	assetsInfo := &models.WalletAssets{}
	if result := database.Db.Model(assetsInfo).Where("id = ?", params.WalletAssetId).Take(assetsInfo); result.Error != nil {
		return nil, result.Error
	}

	//	如果当前资产管理不属于当前用户的资产, 那么提示错误
	if assetsInfo.AdminId != rds.GetAdminSettingId(userInfo.AdminId) {
		return nil, errors.New("incorrectFormat")
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		switch params.Type {
		//	资产系统充值
		case models.WalletBillTypeSystemAssetsDeposit:
			if err := users.UserAssetsDeposit(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userInfo.AdminId,
				AssetId:  params.WalletAssetId,
				ParentId: userInfo.ParentId,
				UserId:   userInfo.Id,
				SourceId: params.WalletAssetId,
				BillType: models.WalletBillTypeSystemAssetsDeposit,
				Money:    params.Money,
			}); err != nil {
				return err
			}

		//	资产系统扣款
		case models.WalletBillTypeSystemAssetsWithdraw:
			if err := users.UserAssetsSpending(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userInfo.AdminId,
				AssetId:  params.WalletAssetId,
				ParentId: userInfo.ParentId,
				UserId:   userInfo.Id,
				SourceId: params.WalletAssetId,
				BillType: models.WalletBillTypeSystemAssetsWithdraw,
				Money:    params.Money,
			}); err != nil {
				return err
			}
		}

		return nil
	})

	return "ok", err
}
