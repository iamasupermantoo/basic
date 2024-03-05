package dtoindex

// FooterData 底部数据
type FooterData struct {
	SocialInfo *FooterInfo   `json:"socialInfo"` //	社区列表
	Items      []*FooterInfo `json:"items"`      //	列表数据
}

// FooterInfo 底部菜单列表
type FooterInfo struct {
	Id       int           `json:"id"`                //	Id
	Icon     string        `json:"icon"`              //	图标
	Name     string        `json:"name"`              //	名称
	Link     string        `json:"link"`              //	链接
	Children []*FooterInfo `json:"children" gorm:"-"` //	子数据
}
