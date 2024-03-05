package dtoindex

// HomeHelpersInfo 帮助中心信息
type HomeHelpersInfo struct {
	ArticleList []*HomeArticleInfo `json:"articleList"` //	文章列表
	SocialList  []*HomeSocialInfo  `json:"socialList"`  //	社区列表
}

// HomeSocialInfo 社区信息
type HomeSocialInfo struct {
	Name   string `json:"name"`   //	名称
	Icon   string `json:"icon"`   //	图标
	Link   string `json:"link"`   //	链接
	Status int    `json:"status"` //	状态
}
