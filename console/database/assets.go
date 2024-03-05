package database

import "basic/models"

// WalletAssetsInit 钱包资产初始化
func WalletAssetsInit() []*models.WalletAssets {
	return []*models.WalletAssets{
		{Id: 1, AdminId: 1, Name: "USDT", Icon: "/assets/currency/usdt.png", Type: models.WalletAssetsTypeDigitalCurrency, Rate: 1},
		{Id: 2, AdminId: 1, Name: "USDC", Icon: "/assets/currency/usdc.png", Type: models.WalletAssetsTypeDigitalCurrency, Rate: 1},
		{Id: 3, AdminId: 1, Name: "BTC", Icon: "/assets/currency/btc.png", Type: models.WalletAssetsTypeDigitalCurrency, Rate: 38546},
		{Id: 4, AdminId: 1, Name: "ETH", Icon: "/assets/currency/eth.png", Type: models.WalletAssetsTypeDigitalCurrency, Rate: 2090},
		{Id: 5, AdminId: 1, Name: "TRX", Icon: "/assets/currency/trx.png", Type: models.WalletAssetsTypeDigitalCurrency, Rate: 0.1},
	}
}
