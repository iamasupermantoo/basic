package settings

import (
	dtosettings "basic/apis/admin/dto/settings"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
)

// ArticleIndex 文章列表
func ArticleIndex(adminId, userId int, params *dtosettings.ArticleIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtosettings.ArticleIndexData, 0)}
	// 过滤where 条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		EqInt("type = ?", params.Type).
		Like("name LIKE ?", params.Name).
		Like("content LIKE ?", params.Content).
		EqInt("status = ?", params.Status).
		Like("data LIKE ?", params.Data).
		BetweenDate("updated_at BETWEEN ? AND ?", params.UpdatedAt)

	result := database.Db.Table("article AS a").
		Select("a.id", "a.admin_id", "au.username AS adminName", "a.type", "a.image", "a.name", "a.content", "a.status", "a.data", "a.updated_at", "a.created_at").
		Joins("left join admin_user AS au on a.admin_id = au.id").
		Where("a.status > ?", models.WalletAssetsStatusDelete).
		Where("a.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)
	return data, result.Error
}

// ArticleCreate 文章创建
func ArticleCreate(adminId, userId int, params *dtosettings.ArticleCreateParams) (interface{}, error) {
	result := database.Db.Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Create(&models.Article{
		AdminId: adminId,
		Type:    params.Type,
		Name:    params.Name,
		Content: params.Content,
	})
	return "ok", result.Error
}

// ArticleDelete 文章删除
func ArticleDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	result := database.Db.Model(&models.Article{}).
		Where("admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id IN ?", params.Ids).
		Update("status", models.ArticleStatusDelete)
	return "ok", result.Error
}

// ArticleUpdate 文章修改
func ArticleUpdate(adminId, userId int, params *dtosettings.ArticleUpdateParams) (interface{}, error) {
	result := database.Db.Model(&models.Article{}).Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Where("id = ?", params.Id).Updates(params)
	return "ok", result.Error
}
