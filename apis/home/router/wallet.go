package router

import (
	"basic/apis/home/controllers/wallets/account"
	"basic/apis/home/controllers/wallets/bill"
	"basic/apis/home/controllers/wallets/deposit"
	"basic/apis/home/controllers/wallets/order"
	"basic/apis/home/controllers/wallets/payment"
	userassets "basic/apis/home/controllers/wallets/userAssets"
	"basic/apis/home/controllers/wallets/withdraw"

	"github.com/gofiber/fiber/v2"
)

// WalletRouter 钱包路由
func WalletRouter(router fiber.Router) {
	router.Route("/wallets", func(router fiber.Router) {
		router.Post("/account/index", account.Index).Name("卡片列表")
		router.Post("/account/info", account.Info).Name("卡片详情")
		router.Post("/account/update", account.Update).Name("修改卡片")
		router.Post("/account/create", account.Create).Name("新增卡片")
		router.Post("/account/delete", account.Delete).Name("删除卡片")
		router.Post("/bill/index", bill.Index).Name("账单列表")
		router.Post("/bill/options", bill.Options).Name("账单类型Options")
		router.Post("/order/index", order.Index).Name("钱包订单")
		router.Post("/payment/index", payment.Index).Name("支付列表")
		router.Post("/withdraw/create", withdraw.Create).Name("提现创建")
		router.Post("/deposit/create", deposit.Create).Name("充值创建")
		router.Post("/user/assets/index", userassets.Index).Name("用户资产列表")
		router.Post("/user/assets/info", userassets.Info).Name("用户资产详细信息")
	})

}
