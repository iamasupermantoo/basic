package database

import "basic/models"

// WalletsUserAccountInit 用户提现账户
func WalletsUserAccountInit() []*models.WalletUserAccount {
	return []*models.WalletUserAccount{
		{AdminId: models.AdminMerchat, UserId: 1, PaymentId: 36, Name: "建设银行", RealName: "隔壁张三", Number: "88888888", Code: "建设银行分行"},
		{AdminId: models.AdminMerchat, UserId: 1, PaymentId: 40, Name: "", RealName: "", Number: "0x888skklxlxlckjlssdfsdfdsfd", Code: ""},
	}
}
