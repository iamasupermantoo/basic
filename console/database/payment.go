package database

import (
	"basic/models"
	"basic/utils"
)

// WalletPaymentInit 初始化钱包支付类型
func WalletPaymentInit() []*models.WalletPayment {
	return []*models.WalletPayment{
		{AdminId: 0, Name: "bankFundsWithdraw", Icon: "/assets/icon/bank.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeWithdraw},
		{AdminId: 0, Name: "bankAssetsWithdraw", Icon: "/assets/icon/bank.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeAssetsWithdraw},
		{AdminId: 0, Name: "digitalFundsWithdraw", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeWithdraw},
		{AdminId: 0, Name: "digitalAssetsWithdraw", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsWithdraw},

		{AdminId: 0, Name: "bankFundsDeposit", Icon: "/assets/icon/bank.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeDeposit},
		{AdminId: 0, Name: "bankAssetsDeposit", Icon: "/assets/icon/bank.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeAssetsDeposit},
		{AdminId: 0, Name: "digitalFundsDeposit", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeDeposit},
		{AdminId: 0, Name: "digitalAssetsDeposit", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsDeposit},

		{Id: 100, AdminId: 1, Name: "DebitCard", Icon: "/assets/icon/bank.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "建设银行", RealName: "隔壁老王", Number: "888866665555", Code: "xxx666"})},
		{Id: 101, AdminId: 1, Name: "BankOfAmerica", Icon: "/assets/icon/bank-of-america-logo.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "建设银行", RealName: "隔壁老王", Number: "888866665555", Code: "xxx666"})},
		{Id: 102, AdminId: 1, Name: "ERC20-USDT", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "ETH", RealName: "USDT", Number: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", Code: "xxx666"})},
		{Id: 103, AdminId: 1, Name: "ERC20-USDC", Icon: "/assets/currency/usdc.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "ETH", RealName: "USDC", Number: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", Code: "xxx666"})},

		{Id: 104, AssetsId: 1, AdminId: 1, Name: "ERC20-USDT", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "ETH", RealName: "USDT", Number: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", Code: "xxx666"})},
		{Id: 105, AssetsId: 2, AdminId: 1, Name: "ERC20-USDC", Icon: "/assets/currency/usdc.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "ETH", RealName: "USDC", Number: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", Code: "xxx666"})},
		{Id: 106, AssetsId: 3, AdminId: 1, Name: "ERC20-BTC", Icon: "/assets/currency/btc.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "ETH", RealName: "BTC", Number: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", Code: "xxx666"})},
		{Id: 107, AssetsId: 4, AdminId: 1, Name: "ERC20-ETH", Icon: "/assets/currency/eth.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "ETH", RealName: "ETH", Number: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", Code: "xxx666"})},
		{Id: 108, AssetsId: 5, AdminId: 1, Name: "ERC20-TRX", Icon: "/assets/currency/trx.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsDeposit, Data: utils.JsonMarshal(&models.WalletPaymentData{Name: "ETH", RealName: "TRX", Number: "0x1Bdd1742C5dEd48bAa7F5D71dba59628D3A3130c", Code: "xxx666"})},

		{AdminId: 1, Name: "DebitCard", Icon: "/assets/icon/bank.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeWithdraw},
		{AdminId: 1, Name: "BankOfAmerica", Icon: "/assets/icon/bank-of-america-logo.png", Type: models.WalletPaymentTypeBank, Mode: models.WalletPaymentModeWithdraw},
		{AdminId: 1, Name: "ERC20-USDT", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeWithdraw},
		{AdminId: 1, Name: "ERC20-USDC", Icon: "/assets/currency/usdc.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeWithdraw},

		{AssetsId: 1, AdminId: 1, Name: "ERC20-USDT", Icon: "/assets/currency/usdt.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsWithdraw},
		{AssetsId: 2, AdminId: 1, Name: "ERC20-USDC", Icon: "/assets/currency/usdc.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsWithdraw},
		{AssetsId: 3, AdminId: 1, Name: "ERC20-BTC", Icon: "/assets/currency/btc.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsWithdraw},
		{AssetsId: 4, AdminId: 1, Name: "ERC20-ETH", Icon: "/assets/currency/eth.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsWithdraw},
		{AssetsId: 5, AdminId: 1, Name: "ERC20-TRX", Icon: "/assets/currency/trx.png", Type: models.WalletPaymentTypeDigital, Mode: models.WalletPaymentModeAssetsWithdraw},
	}
}
