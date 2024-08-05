package database

import "basic/models"

// SystemArticleInit 系统文章初始化
func SystemArticleInit() []*models.Article {
	return []*models.Article{
		{AdminId: 1, Type: models.ArticleTypeFooterAbout, Name: "aboutUsLabel", Content: "aboutUsDesc"},
		{AdminId: 1, Type: models.ArticleTypeFooterAbout, Name: "serviceAgreementLabel", Content: "serviceAgreementDesc"},
		{AdminId: 1, Type: models.ArticleTypeFooterAbout, Name: "privacyAgreementLabel", Content: "privacyAgreementDesc"},

		{AdminId: 1, Type: models.ArticleTypeFooterService, Name: "service1Label", Content: "service1Desc"},
		{AdminId: 1, Type: models.ArticleTypeFooterService, Name: "service2Label", Content: "service2Desc"},
		{AdminId: 1, Type: models.ArticleTypeFooterService, Name: "service3Label", Content: "service3Desc"},

		//	帮助中心
		{AdminId: 1, Type: models.ArticleTypeHelpers, Name: "helpers1Label", Content: "helpers1Desc"},
		{AdminId: 1, Type: models.ArticleTypeHelpers, Name: "helpers2Label", Content: "helpers2Desc"},
		{AdminId: 1, Type: models.ArticleTypeHelpers, Name: "helpers3Label", Content: "helpers3Desc"},
		{AdminId: 1, Type: models.ArticleTypeHelpers, Name: "helpers4Label", Content: "helpers4Desc"},
		{AdminId: 1, Type: models.ArticleTypeHelpers, Name: "helpers5Label", Content: "helpers5Desc"},
		{AdminId: 1, Type: models.ArticleTypeHelpers, Name: "helpers6Label", Content: "helpers6Desc"},
		{AdminId: 1, Type: models.ArticleTypeHelpers, Name: "helpers7Label", Content: "helpers7Desc"},
	}
}
