package database

import (
	"basic/models"
	"time"
)

// WalletsBillInit 初始化账单
func WalletsBillInit() []*models.WalletBill {
	return []*models.WalletBill{
		{AdminId: models.AdminMerchat, UserId: 1, AssetsId: 0, SourceId: 0, Type: models.WalletBillTypeTeamEarnings, Name: models.WalletBillDepositTypeList[models.WalletBillTypeTeamEarnings], Money: 100, Balance: 0, CreatedAt: int(time.Now().Unix())},
		{AdminId: models.AdminMerchat, UserId: 1, AssetsId: 0, SourceId: 0, Type: models.WalletBillTypeWithdraw, Name: models.WalletBillSpendTypeList[models.WalletBillTypeWithdraw], Money: 100, Balance: 0, CreatedAt: int(time.Now().Unix())},
		{AdminId: models.AdminMerchat, UserId: 1, AssetsId: 0, SourceId: 0, Type: models.WalletBillTypeTeamEarnings, Name: models.WalletBillDepositTypeList[models.WalletBillTypeTeamEarnings], Money: 100, Balance: 0, CreatedAt: int(time.Now().Unix())},

		{AdminId: models.AdminMerchat, UserId: 1, AssetsId: 6, SourceId: 0, Type: models.WalletBillTypeTeamAssetsEarnings, Name: models.WalletBillDepositTypeList[models.WalletBillTypeTeamAssetsEarnings], Money: 200, Balance: 0, CreatedAt: int(time.Now().Unix())},
		{AdminId: models.AdminMerchat, UserId: 1, AssetsId: 7, SourceId: 0, Type: models.WalletBillTypeTeamAssetsEarnings, Name: models.WalletBillDepositTypeList[models.WalletBillTypeTeamAssetsEarnings], Money: 200, Balance: 0, CreatedAt: int(time.Now().Unix())},

		// 下级团队收益
		{AdminId: models.AdminMerchat, UserId: 2, AssetsId: 0, SourceId: 0, Type: models.WalletBillTypeTeamEarnings, Name: models.WalletBillDepositTypeList[models.WalletBillTypeTeamEarnings], Money: 100, Balance: 0, CreatedAt: int(time.Now().Unix())},
		{AdminId: models.AdminMerchat, UserId: 3, AssetsId: 0, SourceId: 0, Type: models.WalletBillTypeTeamEarnings, Name: models.WalletBillDepositTypeList[models.WalletBillTypeTeamEarnings], Money: 100, Balance: 0, CreatedAt: int(time.Now().Unix())},
		{AdminId: models.AdminMerchat, UserId: 4, AssetsId: 0, SourceId: 0, Type: models.WalletBillTypeTeamEarnings, Name: models.WalletBillDepositTypeList[models.WalletBillTypeTeamEarnings], Money: 100, Balance: 0, CreatedAt: int(time.Now().Unix())},
	}
}
