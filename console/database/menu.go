package database

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/models"
	"basic/utils"
)

// AdminMenuInit 管理菜单初始化
func AdminMenuInit() []*models.AdminMenu {
	return []*models.AdminMenu{
		//	用户管理
		{Id: 100, Status: 10, ParentId: 0, Name: "用户管理", Route: "", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "manage_accounts"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 101, Status: 10, ParentId: 100, Name: "用户列表", Route: "/user/index", Sort: 0, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "people", VueTmp: "/views", ViewsUrl: "/user/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 102, Status: 10, ParentId: 100, Name: "用户资产", Route: "/user/assets/index", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "card_giftcard", VueTmp: "/views", ViewsUrl: "/user/assets/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 103, Status: 10, ParentId: 100, Name: "会员列表", Route: "/user/level/index", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "diamond", VueTmp: "/views", ViewsUrl: "/user/level/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 104, Status: 10, ParentId: 100, Name: "认证申请", Route: "/user/auth/index", Sort: 3, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "badge", VueTmp: "/views", ViewsUrl: "/user/auth/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 105, Status: 10, ParentId: 100, Name: "邀请管理", Route: "/user/invite/index", Sort: 4, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "share", VueTmp: "/views", ViewsUrl: "/user/invite/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 106, Status: 10, ParentId: 100, Name: "用户关系", Route: "/user/relation/index", Sort: 5, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "diversity_1", VueTmp: "/user/relation", ViewsUrl: "/user/relation/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},

		//	钱包管理
		{Id: 120, Status: 10, ParentId: 0, Name: "钱包管理", Route: "", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "account_balance_wallet"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 121, Status: 10, ParentId: 120, Name: "充值订单", Route: "/wallets/deposit/index", Sort: 0, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "format_textdirection_r_to_l", VueTmp: "/views", ViewsUrl: "/wallets/deposit/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 122, Status: 10, ParentId: 120, Name: "提现订单", Route: "/wallets/withdraw/index", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "format_textdirection_l_to_r", VueTmp: "/views", ViewsUrl: "/wallets/withdraw/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 123, Status: 10, ParentId: 120, Name: "账单明细", Route: "/wallets/bill/index", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "assignment", VueTmp: "/views", ViewsUrl: "/wallets/bill/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 124, Status: 10, ParentId: 120, Name: "资产管理", Route: "/wallets/assets/index", Sort: 3, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "card_giftcard", VueTmp: "/views", ViewsUrl: "/wallets/assets/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 125, Status: 10, ParentId: 120, Name: "支付管理", Route: "/wallets/payment/index", Sort: 4, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "currency_exchange", VueTmp: "/views", ViewsUrl: "/wallets/payment/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 126, Status: 10, ParentId: 120, Name: "卡片管理", Route: "/wallets/account/index", Sort: 5, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "credit_card", VueTmp: "/views", ViewsUrl: "/wallets/account/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},

		//	后台管理
		{Id: 1, Status: 10, ParentId: 0, Name: "后台管理", Route: "", Sort: 90, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "manage_accounts"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 2, Status: 10, ParentId: 1, Name: "管理列表", Route: "/manage/index", Sort: 0, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "manage_accounts", VueTmp: "/views", ViewsUrl: "/manage/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 3, Status: 10, ParentId: 1, Name: "角色配置", Route: "/manage/role/index", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "6_ft_apart", VueTmp: "/views", ViewsUrl: "/manage/role/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 4, Status: 10, ParentId: 1, Name: "菜单管理", Route: "/manage/menu/index", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "menu_open", VueTmp: "/views", ViewsUrl: "/manage/menu/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 5, Status: 10, ParentId: 1, Name: "操作日志", Route: "/manage/logs/index", Sort: 3, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "event_note", VueTmp: "/views", ViewsUrl: "/manage/logs/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},

		//	配置管理
		{Id: 20, Status: 10, ParentId: 0, Name: "配置管理", Route: "", Sort: 91, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "settings"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 21, Status: 10, ParentId: 20, Name: "参数配置", Route: "/setting/index", Sort: 0, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "settings_ethernet", VueTmp: "/views", ViewsUrl: "/setting/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 22, Status: 10, ParentId: 20, Name: "菜单配置", Route: "/setting/menu/index", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "menu_open", VueTmp: "/views", ViewsUrl: "/setting/menu/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 23, Status: 10, ParentId: 20, Name: "文章配置", Route: "/setting/article/index", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "article", VueTmp: "/views", ViewsUrl: "/setting/article/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},

		//	系统配置
		{Id: 40, Status: 10, ParentId: 0, Name: "系统配置", Route: "", Sort: 92, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "widgets"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 41, Status: 10, ParentId: 40, Name: "等级配置", Route: "/system/level", Sort: 0, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "diamond", VueTmp: "/views", ViewsUrl: "/system/level/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 42, Status: 10, ParentId: 40, Name: "国家配置", Route: "/system/country", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "flag", VueTmp: "/views", ViewsUrl: "/system/country/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 43, Status: 10, ParentId: 40, Name: "语言配置", Route: "/system/lang", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "language", VueTmp: "/views", ViewsUrl: "/system/lang/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 44, Status: 10, ParentId: 40, Name: "翻译配置", Route: "/system/translate", Sort: 3, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "translate", VueTmp: "/views", ViewsUrl: "/system/translate/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},
		{Id: 45, Status: 10, ParentId: 40, Name: "通知配置", Route: "/system/notify", Sort: 4, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "notifications", VueTmp: "/views", ViewsUrl: "/system/notify/views"}), AdminId: 1, Type: models.AdminMenuTypeAdminMenu},

		//	用户我的菜单
		{AdminId: models.SuperAdminId, Sort: 20, Id: 200, Name: "walletManage", Type: models.AdminMenuTypeUserMenu, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/wallets-manage.png"})},
		{AdminId: models.SuperAdminId, Sort: 1, ParentId: 200, Name: "accountManage", Route: "/wallets/account/index", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/account.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 2, ParentId: 200, Name: "myWallet", Route: "/wallets/index", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/wallets.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 3, ParentId: 200, Name: "myAssets", Route: "/wallets/assets/index", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/assets.png"}), Type: models.AdminMenuTypeUserMenu},

		{AdminId: models.SuperAdminId, Sort: 21, Id: 220, Name: "teamManage", Type: models.AdminMenuTypeUserMenu, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/team-manage.png"})},
		{AdminId: models.SuperAdminId, Sort: 1, ParentId: 220, Name: "inviteFriends", Route: "/team/share", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/share.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 2, ParentId: 220, Name: "myTeam", Route: "/team/index", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/team.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 3, ParentId: 220, Name: "teamEarnings", Route: "/team/earnings", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/earnings.png"}), Type: models.AdminMenuTypeUserMenu},

		{AdminId: models.SuperAdminId, Sort: 22, Id: 230, Name: "serviceManage", Type: models.AdminMenuTypeUserMenu, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/service-manage.png"})},
		{AdminId: models.SuperAdminId, Sort: 1, ParentId: 230, Name: "memberVip", Route: "/user/level", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/level.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 2, ParentId: 230, Name: "realAuth", Route: "/user/auth", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/realauth.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 3, ParentId: 230, Name: "helpers", Route: "/helpers", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/helpers.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 4, ParentId: 230, Name: "settings", Route: "/user/settings", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/settings.png"}), Type: models.AdminMenuTypeUserMenu},
		{AdminId: models.SuperAdminId, Sort: 5, ParentId: 230, Name: "download", Route: "/download", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/download.png"}), Type: models.AdminMenuTypeUserMenu},

		//	首页菜单
		{AdminId: models.SuperAdminId, Id: 260, Name: "myWallet", Route: "/wallets/index", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/wallets.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeHomeMenu},
		{AdminId: models.SuperAdminId, Id: 261, Name: "myAssets", Route: "/wallets/assets/index", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/assets.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeHomeMenu},
		{AdminId: models.SuperAdminId, Id: 262, Name: "myTeam", Route: "/team/index", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/team.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeHomeMenu},
		{AdminId: models.SuperAdminId, Id: 263, Name: "inviteFriends", Route: "/team/share", Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/share.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeHomeMenu},

		//	用户端 H5 TabBar菜单
		{Id: 300, AdminId: models.SuperAdminId, Name: "home", Route: "/", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/home.png", ActiveIcon: "/assets/icon/tabbar/home_active.png", IsPc: false, IsMobile: true}), Type: models.AdminMenuTypeTabularMenu},
		{Id: 301, AdminId: models.SuperAdminId, Name: "markets", Route: "/markets", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/markets.png", ActiveIcon: "/assets/icon/tabbar/markets_active.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeTabularMenu},
		{Id: 302, AdminId: models.SuperAdminId, Name: "contract", Route: "/contract", Sort: 3, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/contract.png", ActiveIcon: "/assets/icon/tabbar/contract_active.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeTabularMenu},
		{Id: 303, AdminId: models.SuperAdminId, Name: "futures", Route: "/futures", Sort: 4, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/futures.png", ActiveIcon: "/assets/icon/tabbar/futures_active.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeTabularMenu},
		{Id: 304, AdminId: models.SuperAdminId, Name: "user", Route: "/user", Sort: 5, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/user.png", ActiveIcon: "/assets/icon/tabbar/user_active.png", IsPc: false, IsMobile: true}), Type: models.AdminMenuTypeTabularMenu},

		//	子级菜单
		{ParentId: 301, AdminId: models.SuperAdminId, Name: "digitalMarkets", Route: "/markets/digital", Sort: 0, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/markets_active.png", IsPc: true, IsMobile: false}), Type: models.AdminMenuTypeTabularMenu},
		{ParentId: 301, AdminId: models.SuperAdminId, Name: "forexMarkets", Route: "/markets/forex", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/contract_active.png", IsPc: true, IsMobile: false}), Type: models.AdminMenuTypeTabularMenu},
		{ParentId: 301, AdminId: models.SuperAdminId, Name: "futuresMarkets", Route: "/markets/futures", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/tabbar/futures_active.png", IsPc: true, IsMobile: false}), Type: models.AdminMenuTypeTabularMenu},

		//	用户快捷菜单
		{Id: 400, AdminId: models.SuperAdminId, Name: "deposit", Route: "/wallets/deposit?mode=1", Sort: 1, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/deposit.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeQuickMenu},
		{Id: 401, AdminId: models.SuperAdminId, Name: "withdraw", Route: "/wallets/withdraw?mode=11", Sort: 2, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/withdraw.png", IsPc: true, IsMobile: true}), Type: models.AdminMenuTypeQuickMenu},
		{Id: 402, AdminId: models.SuperAdminId, Name: "helpers", Route: "/helpers", Sort: 3, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/chatbot.png", IsPc: true, IsMobile: false}), Type: models.AdminMenuTypeQuickMenu},
		{Id: 403, AdminId: models.SuperAdminId, Name: "download", Route: "/download", Sort: 4, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/download.png", IsPc: true, IsMobile: false}), Type: models.AdminMenuTypeQuickMenu},
		{Id: 404, AdminId: models.SuperAdminId, Name: "memberVip", Route: "/user/level", Sort: 4, Data: utils.JsonMarshal(&dtoadmins.AdminMenuData{Icon: "/assets/icon/menu/level.png", IsPc: true, IsMobile: false}), Type: models.AdminMenuTypeQuickMenu},
	}
}
