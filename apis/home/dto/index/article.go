package dtoindex

// ArticleIndexParams 获取文章列表参数
type ArticleIndexParams struct {
	Types []int `json:"types"` //	文章类型
}

// ArticleInfoParams 获取文章信息参数
type ArticleInfoParams struct {
	Id int `json:"id" validate:"required"` //	文章Id
}

// HomeArticleInfo 文章信息
type HomeArticleInfo struct {
	Id        int    `json:"id"`        //	文章Id
	AdminId   int    `json:"adminId"`   //	管理员Id
	Name      string `json:"name"`      //	标题
	Content   string `json:"content"`   //	内容
	CreatedAt int    `json:"createdAt"` //	文章发布时间
}
