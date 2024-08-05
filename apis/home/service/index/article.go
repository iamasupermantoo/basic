package index

import (
	"basic/apis/common/rds"
	dtoindex "basic/apis/home/dto/index"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
)

// ArticleIndex  根据type获取文章标题和内容
func ArticleIndex(host, lang string, params *dtoindex.ArticleIndexParams) (interface{}, error) {
	//	过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt(" type IN ?", params.Types)

	//	通过缓存换取adminId
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	通过域名获取adminId
	adminId := rds.RedisFindDomainToAdminId(rdsConn, host)
	var data []*dtoindex.HomeArticleInfo

	//	type不为0则展示缓存adminId的type中的数据
	if result := database.Db.Table("article").
		Where("admin_id = ?", adminId).
		Scopes(filterParams.Scopes()).
		Find(&data); result.Error != nil {
		return nil, result.Error
	}
	for _, item := range data {
		item.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, item.Name)
		item.Content = rds.RedisFindTranslateField(rdsConn, adminId, lang, item.Content)
	}
	return data, nil
}

// ArticleInfo  根据id获取文章标题和内容
func ArticleInfo(lang string, params *dtoindex.ArticleInfoParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	通过id获取内容和adminId
	data := &dtoindex.HomeArticleInfo{}
	if result := database.Db.Model(models.Article{}).Where("id = ?", params.Id).Take(data); result.Error != nil {
		return nil, result.Error
	}

	//	翻译标题和内容
	data.Name = rds.RedisFindTranslateField(rdsConn, data.AdminId, lang, data.Name)
	data.Content = rds.RedisFindTranslateField(rdsConn, data.AdminId, lang, data.Content)
	return data, nil
}
