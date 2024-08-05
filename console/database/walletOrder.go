package database

import (
	"basic/models"
	"basic/utils"
)

// WalletsOrderInit 初始化钱包订单
func WalletsOrderInit() []*models.WalletOrder {
	return []*models.WalletOrder{
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 100, Type: models.WalletOrderTypeDeposit, Money: 100, Status: models.WalletOrderStatusActive, OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 100, Type: models.WalletOrderTypeDeposit, Money: 100, Status: models.WalletOrderStatusRefuse, Data: "拒绝提现, 您还未达到标准提现", OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 102, Type: models.WalletOrderTypeDeposit, Money: 100, Status: models.WalletOrderStatusComplete, OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 102, Type: models.WalletOrderTypeDeposit, Money: 100, OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 102, Type: models.WalletOrderTypeDeposit, Money: 100, OrderSn: utils.NewRandom().OrderSn()},

		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 104, AssetsId: 6, Type: models.WalletOrderTypeAssetsDeposit, Money: 200, Status: models.WalletOrderStatusActive, OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 104, AssetsId: 6, Type: models.WalletOrderTypeAssetsDeposit, Money: 200, Status: models.WalletOrderStatusRefuse, Data: "拒绝提现, 您还未达到标准提现", OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 106, AssetsId: 6, Type: models.WalletOrderTypeAssetsDeposit, Money: 200, Status: models.WalletOrderStatusComplete, OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 106, AssetsId: 7, Type: models.WalletOrderTypeAssetsDeposit, Money: 200, OrderSn: utils.NewRandom().OrderSn()},
		{AdminId: models.AdminMerchat, UserId: 1, SourceId: 106, AssetsId: 7, Type: models.WalletOrderTypeAssetsDeposit, Money: 200, OrderSn: utils.NewRandom().OrderSn()},
	}
}
