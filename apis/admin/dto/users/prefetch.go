package dtousers

import "basic/apis/common/dto"

// UserInitData 用户初始化数据
type UserInitData struct {
	Config       *UserInitConfig         `json:"config"`       //	初始化配置
	TabBars      []*HomeMenuInfo         `json:"tabBars"`      //	tabBar菜单
	QuickMenu    []*HomeMenuInfo         `json:"quickMenu"`    //	快捷菜单
	UserMenu     []*HomeMenuInfo         `json:"userMenu"`     //	我的菜单
	HomeMenu     []*HomeMenuInfo         `json:"homeMenu"`     // 前台菜单
	Translate    []*dto.TranslateOptions `json:"translate"`    //	翻译列表
	LanguageList []*UserInitLanguage     `json:"languageList"` //	语言列表
	CountryList  []*UserInitCountry      `json:"countryList"`  //	国家列表
}

// UserInitConfig 用户初始化配置
type UserInitConfig struct {
	Name         string            `json:"name"`         //	项目名称
	Logo         string            `json:"logo"`         //	项目Logo
	DefaultLang  string            `json:"defaultLang"`  //	默认语言
	Template     string            `json:"template"`     //	模版
	Settings     *UserInitSettings `json:"settings"`     //	配置
	DepositTips  string            `json:"depositTips"`  //	充值注意事项
	WithdrawTips string            `json:"withdrawTips"` //	提现注意事项
	OnlineIcon   string            `json:"onlineIcon"`   //	客服图标
	OnlineLink   string            `json:"onlineLink"`   //	客服链接
}

// UserInitLanguage 初始化语言
type UserInitLanguage struct {
	Id    int    `json:"id"`    //	语言ID
	Name  string `json:"name"`  //	语言名称
	Alias string `json:"alias"` //	别名
	Icon  string `json:"icon"`  //	语言图标
}

// UserInitCountry 初始化国家
type UserInitCountry struct {
	Id   int    `json:"id"`                       //	国家Id
	Name string `json:"name" gorm:"column:alias"` //	国家名称
	Icon string `json:"icon"`                     //	国家图标
	Code string `json:"code"`                     //	国家代码
}

// UserInitSettings 用户初始化配置
type UserInitSettings struct {
	Register map[string]bool `json:"register"` //	注册配置
	Login    map[string]bool `json:"login"`    //	登录配置
	Lang     map[string]bool `json:"lang"`     //	语言配置
	Online   map[string]bool `json:"online"`   //	在线客服配置
	Wallets  map[string]bool `json:"wallets"`  //	钱包配置
}
