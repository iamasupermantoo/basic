package index

import (
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	dtoindex "basic/apis/home/dto/index"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"encoding/json"
)

// HelpersInfo 获取帮助列表
func HelpersInfo(host, lang string) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	域名获取管理ID
	adminId := rds.RedisFindDomainToAdminId(rdsConn, host)
	socialListStr := rds.RedisFindAdminSettingField(rdsConn, adminId, "socialList")

	data := &dtoindex.HomeHelpersInfo{
		ArticleList: make([]*dtoindex.HomeArticleInfo, 0),
		SocialList:  make([]*dtoindex.HomeSocialInfo, 0),
	}

	//	获取社区缓存设置
	socialList := make([]*dtoindex.HomeSocialInfo, 0)
	if err := json.Unmarshal([]byte(socialListStr), &socialList); err != nil {
		return nil, err
	}
	for _, social := range socialList {
		if social.Status == dto.ActivateStatus {
			social.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, social.Name)
			data.SocialList = append(data.SocialList, social)
		}
	}

	//	查询文章，帮助中心
	if result := database.Db.Model(&models.Article{}).
		Where("admin_id = ?", adminId).
		Where("type = ?", models.ArticleTypeHelpers).
		Find(&data.ArticleList); result.Error != nil {
		return nil, result.Error
	}
	for _, article := range data.ArticleList {
		article.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, article.Name)
		article.Content = rds.RedisFindTranslateField(rdsConn, adminId, lang, article.Content)
	}

	return data, nil
}
