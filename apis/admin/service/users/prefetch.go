package users

import (
	dtousers "basic/apis/admin/dto/users"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/utils"
)

// UserInit 用户初始化数据
func UserInit(domain, lang string) *dtousers.UserInitData {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	获取域名绑定的管理ID, 如果没有绑定那么使用超级管理ID
	adminId := rds.RedisFindDomainToAdminId(rdsConn, domain)
	if adminId == 0 {
		adminId = models.SuperAdminId
	}

	//	获取语言列表, 如果没有传递设置语言, 那么使用第一个语言
	langList := rds.RedisFindLangList(rdsConn, adminId)
	if lang == "" && len(langList) > 0 {
		lang = langList[0].Alias
	}

	adminDataInfo := rds.RedisFindAdminDataInfo(rdsConn, adminId)
	homeMenuList := rds.RedisFindHomeMenuList(rdsConn, adminId)
	return &dtousers.UserInitData{
		Config: &dtousers.UserInitConfig{
			Name:         rds.RedisFindAdminSettingField(rdsConn, adminId, "siteName"),
			Logo:         rds.RedisFindAdminSettingField(rdsConn, adminId, "siteLogo"),
			DefaultLang:  lang,
			Template:     adminDataInfo.Template,
			OnlineIcon:   rds.RedisFindAdminSettingField(rdsConn, adminId, "onlineImage"),
			OnlineLink:   rds.RedisFindAdminSettingField(rdsConn, adminId, "onlineLink"),
			DepositTips:  rds.RedisFindTranslateField(rdsConn, adminId, lang, rds.RedisFindAdminSettingField(rdsConn, adminId, "walletDepositDesc")),
			WithdrawTips: rds.RedisFindTranslateField(rdsConn, adminId, lang, rds.RedisFindAdminSettingField(rdsConn, adminId, "walletWithdrawDesc")),
			Settings: &dtousers.UserInitSettings{
				Register: utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateRegister")),
				Login:    utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateLogin")),
				Lang:     utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateLang")),
				Online:   utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateOnline")),
				Wallets:  utils.StringToBoolMaps(rds.RedisFindAdminSettingField(rdsConn, adminId, "templateWallet")),
			},
		},
		TabBars:      homeMenuList.TabbarList,
		QuickMenu:    homeMenuList.QuickList,
		UserMenu:     homeMenuList.MenuList,
		HomeMenu:     homeMenuList.HomeList,
		CountryList:  rds.RedisFindCountryList(rdsConn, adminId),
		LanguageList: rds.RedisFindLangList(rdsConn, adminId),
		Translate:    rds.RedisFindHomeTranslateLang(rdsConn, adminId, lang),
	}
}
