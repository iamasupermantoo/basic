package database

import "basic/models"

// WalletsUserAssetsInit 初始化用户资产
func WalletsUserAssetsInit() []*models.WalletUserAssets {
	return []*models.WalletUserAssets{
		{AdminId: models.AdminMerchat, UserId: 1, WalletAssetsId: 6, Money: 100},
		{AdminId: models.AdminMerchat, UserId: 1, WalletAssetsId: 7, Money: 200},
	}
}
