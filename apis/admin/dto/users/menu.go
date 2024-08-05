package dtousers

// HomeMenuList 前台菜单组合
type HomeMenuList struct {
	TabbarList []*HomeMenuInfo `json:"tabbarList"` //	tabbar列表			手机端Tabbar 和 电脑端 navs 菜单
	QuickList  []*HomeMenuInfo `json:"quickList"`  //	快捷菜单列表	 	手机端我的页面 快捷按钮 和 navs 更多菜单
	MenuList   []*HomeMenuInfo `json:"menuList"`   //	用户菜单列表	 	手机端我的菜单 电脑端 设置菜单
	HomeList   []*HomeMenuInfo `json:"homeList"`   //	首页快捷菜单列表		手机端首页菜单 电脑端头像子菜单
}

// HomeMenuDataInfo 前台菜单数据
type HomeMenuDataInfo struct {
	IsMobile  bool `json:"isMobile"`  //	是否显示手机端
	IsDesktop bool `json:"isDesktop"` //	是否显示桌面端
}

// HomeMenuInfo
type HomeMenuInfo struct {
	Id         int               `json:"id"`         //	ID
	Name       string            `json:"name"`       //	名称
	Route      string            `json:"route"`      //	路由
	Icon       string            `json:"icon"`       //	图标
	ActiveIcon string            `json:"activeIcon"` //	激活图标
	Data       *HomeMenuDataInfo `json:"data"`       //	菜单数据
	Children   []*HomeMenuInfo   `json:"children"`   //	子路由
}
