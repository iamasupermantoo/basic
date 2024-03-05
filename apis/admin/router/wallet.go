package router

import (
	"basic/apis/admin/controllers/wallets/account"
	"basic/apis/admin/controllers/wallets/assets"
	"basic/apis/admin/controllers/wallets/bill"
	"basic/apis/admin/controllers/wallets/deposit"
	"basic/apis/admin/controllers/wallets/payment"
	userassets "basic/apis/admin/controllers/wallets/userAssets"
	"basic/apis/admin/controllers/wallets/withdraw"

	"github.com/gofiber/fiber/v2"
)

func WalletRouter(router fiber.Router) {
	router.Route("/wallets", func(router fiber.Router) {
		// 用户充值管理
		router.Route("/deposit", func(router fiber.Router) {
			router.Post("/status", deposit.Status).Name("用户充值审核")
			router.Post("/create", deposit.Create).Name("用户充值新增")
			router.Post("/delete", deposit.Delete).Name("用户充值删除")
			router.Post("/update", deposit.Update).Name("用户充值更新")
			router.Post("/index", deposit.Index).Name("用户充值列表")
			router.Post("/views", deposit.Views).Name("用户充值Json模版")
		})

		// 用户账单
		router.Route("/bill", func(router fiber.Router) {
			router.Post("/index", bill.Index).Name("用户账单列表")
			router.Post("/views", bill.Views).Name("用户账单Json模版")
		})

		// 钱包账户
		router.Route("/account", func(router fiber.Router) {
			router.Post("/create", account.Create).Name("用户提现账户新增")
			router.Post("/delete", account.Delete).Name("用户提现账户删除")
			router.Post("/update", account.Update).Name("用户提现账户更新")
			router.Post("/index", account.Index).Name("用户提现账户列表")
			router.Post("/views", account.Views).Name("用户提现账户Json模版")
		})

		//钱包支付方式
		router.Route("/payment", func(router fiber.Router) {
			router.Post("/create", payment.Create).Name("支付创建")
			router.Post("/delete", payment.Delete).Name("支付删除")
			router.Post("/update", payment.Update).Name("支付更新")
			router.Post("/desc", payment.UpdateDesc).Name("支付更新数据")
			router.Post("/index", payment.Index).Name("支付列表")
			router.Post("/views", payment.Views).Name("支付Json模版")
		})

		// 用户提现管理
		router.Route("/withdraw", func(router fiber.Router) {
			router.Post("/create", withdraw.Create).Name("用户提现新增")
			router.Post("/delete", withdraw.Delete).Name("用户提现删除")
			router.Post("/update", withdraw.Update).Name("用户提现更新")
			router.Post("/index", withdraw.Index).Name("用户提现列表")
			router.Post("/status", withdraw.Status).Name("用户提现审核")
			router.Post("/views", withdraw.Views).Name("用户提现Json模版")
		})

		// 资产管理
		router.Route("/assets", func(router fiber.Router) {
			router.Post("/create", assets.Create).Name("资产创建")
			router.Post("/delete", assets.Delete).Name("资产删除")
			router.Post("/update", assets.Update).Name("资产更新")
			router.Post("/index", assets.Index).Name("资产列表")
			router.Post("/views", assets.Views).Name("资产Json模版")
		})
	})

	// 用户资产管理
	router.Route("/user", func(router fiber.Router) {
		router.Route("/assets", func(router fiber.Router) {
			router.Post("/balance", userassets.Balance).Name("用户资产修改")
			router.Post("/delete", userassets.Delete).Name("用户资产删除")
			router.Post("/index", userassets.Index).Name("用户资产列表")
			router.Post("/update", userassets.Update).Name("用户资产更新")
			router.Post("/views", userassets.Views).Name("用户资产json模版")
		})
	})

}
