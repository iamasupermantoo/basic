package index

import (
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	dtoindex "basic/apis/home/dto/index"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"github.com/goccy/go-json"
	"strconv"
)

// FooterIndex 获取底部列表
func FooterIndex(host string, lang string) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	域名获取管理ID
	adminId := rds.RedisFindDomainToAdminId(rdsConn, host)

	data := &dtoindex.FooterData{
		Items: make([]*dtoindex.FooterInfo, 0),
	}

	//	关于我们文章
	aboutList := make([]*dtoindex.FooterInfo, 0)
	aboutArticleList := make([]*models.Article, 0)
	if result := database.Db.Model(&models.Article{}).
		Where("admin_id = ?", adminId).
		Where("type = ?", models.ArticleTypeFooterAbout).
		Where("status = ?", models.ArticleStatusActive).
		Limit(4).
		Find(&aboutArticleList); result.Error != nil {
		return result.Error, nil
	}
	for _, article := range aboutArticleList {
		aboutList = append(aboutList, &dtoindex.FooterInfo{
			Name:     rds.RedisFindTranslateField(rdsConn, adminId, lang, article.Name),
			Link:     "/article/details?id=" + strconv.Itoa(article.Id),
			Children: make([]*dtoindex.FooterInfo, 0),
		})
	}

	//	平台服务
	siteServiceList := make([]*dtoindex.FooterInfo, 0)
	siteServiceArticleList := make([]*models.Article, 0)
	if result := database.Db.Model(&models.Article{}).
		Where("admin_id = ?", adminId).
		Where("type = ?", models.ArticleTypeFooterService).
		Where("status = ?", models.ArticleStatusActive).
		Limit(4).
		Find(&siteServiceArticleList); result.Error != nil {
		return result.Error, nil
	}
	for _, article := range siteServiceArticleList {
		siteServiceList = append(siteServiceList, &dtoindex.FooterInfo{
			Name:     rds.RedisFindTranslateField(rdsConn, adminId, lang, article.Name),
			Link:     "/article/details?id=" + strconv.Itoa(article.Id),
			Children: make([]*dtoindex.FooterInfo, 0),
		})
	}

	//	帮助中心
	helpersList := make([]*dtoindex.FooterInfo, 0)
	helpersArticleList := make([]*models.Article, 0)
	if result := database.Db.Model(&models.Article{}).
		Where("admin_id = ?", adminId).
		Where("type = ?", models.ArticleTypeHelpers).
		Where("status = ?", models.ArticleStatusActive).
		Limit(4).
		Find(&helpersArticleList); result.Error != nil {
		return result.Error, nil
	}
	for _, article := range helpersArticleList {
		helpersList = append(helpersList, &dtoindex.FooterInfo{
			Name:     rds.RedisFindTranslateField(rdsConn, adminId, lang, article.Name),
			Link:     "/article/details?id=" + strconv.Itoa(article.Id),
			Children: make([]*dtoindex.FooterInfo, 0),
		})
	}

	//	联系我们
	contactListStr := rds.RedisFindAdminSettingField(rdsConn, adminId, "contactList")
	contactList := make([]*dtoindex.HomeSocialInfo, 0)
	contactFooterList := make([]*dtoindex.FooterInfo, 0)
	if err := json.Unmarshal([]byte(contactListStr), &contactList); err != nil {
		return nil, err
	}
	for _, contact := range contactList {
		if contact.Status == dto.ActivateStatus {
			contactFooterList = append(contactFooterList, &dtoindex.FooterInfo{
				Icon:     contact.Icon,
				Name:     rds.RedisFindTranslateField(rdsConn, adminId, lang, contact.Name),
				Link:     contact.Link,
				Children: make([]*dtoindex.FooterInfo, 0),
			})
		}
	}

	//	社区列表
	socialListStr := rds.RedisFindAdminSettingField(rdsConn, adminId, "socialList")
	socialList := make([]*dtoindex.HomeSocialInfo, 0)
	socialListFooter := make([]*dtoindex.FooterInfo, 0)
	if err := json.Unmarshal([]byte(socialListStr), &socialList); err != nil {
		return nil, err
	}
	for _, social := range socialList {
		if social.Status == dto.ActivateStatus {
			socialListFooter = append(socialListFooter, &dtoindex.FooterInfo{
				Name:     rds.RedisFindTranslateField(rdsConn, adminId, lang, social.Name),
				Icon:     social.Icon,
				Link:     social.Link,
				Children: make([]*dtoindex.FooterInfo, 0),
			})
		}
	}
	data.SocialInfo = &dtoindex.FooterInfo{
		Name:     rds.RedisFindTranslateField(rdsConn, adminId, lang, "social"),
		Children: socialListFooter,
	}

	//	页脚数据
	data.Items = append(data.Items, []*dtoindex.FooterInfo{
		{Name: rds.RedisFindTranslateField(rdsConn, adminId, lang, "aboutUs"), Children: aboutList},
		{Name: rds.RedisFindTranslateField(rdsConn, adminId, lang, "siteService"), Children: siteServiceList},
		{Name: rds.RedisFindTranslateField(rdsConn, adminId, lang, "helpers"), Children: helpersList},
		{Name: rds.RedisFindTranslateField(rdsConn, adminId, lang, "contactUs"), Children: contactFooterList},
	}...)
	return data, nil
}
