package dtoindex

// HomeTranslateParams 翻译列表参数
type HomeTranslateParams struct {
	Lang string `json:"lang" validate:"required"` //	语言别名
}
