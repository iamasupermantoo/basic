package wallets

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
)

// AssetsIndex 资产列表
func AssetsIndex(adminId, userId int, params *dtowallets.AssetsIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtowallets.AssetsIndexData, 0)}
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		EqStr("name LIKE ?", params.Name+"%").
		EqInt("type = ?", params.Type).
		EqInt("status = ?", params.Status).
		Like("data LIKE ?", params.Data).
		BetweenDate("updated_at BETWEEN ? AND  ?", params.UpdatedAt).
		BetweenDate("created_at BETWEEN ? AND  ?", params.CreatedAt)

	if result := database.Db.Table("wallet_assets AS wa").
		Select("wa.id", "wa.admin_id", "au.username AS adminName", "wa.name", "wa.icon", "wa.type", "wa.rate", "wa.status", "wa.data", "wa.updated_at", "wa.created_at").
		Joins("LEFT JOIN 	admin_user AS au ON au.id = wa.admin_id").
		Where("wa.status > ?", models.WalletAssetsStatusDelete).
		Where("wa.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

// AssetsDelete 资产更新
func AssetsDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletAssets{}).
		Where("admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id IN ?", params.Ids).
		Update("status", models.WalletAssetsStatusDelete); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// AssetsUpdate 资产更新
func AssetsUpdate(adminId, userId int, params *dtowallets.AssetsUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletAssets{}).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id = ?", params.Id).
		Updates(params); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// AssetsCreate 资产更新
func AssetsCreate(adminId, userId int, params *dtowallets.AssetsCreateParams) (interface{}, error) {
	if result := database.Db.Create(&models.WalletAssets{
		AdminId: adminId,
		Name:    params.Name,
		Icon:    params.Icon,
		Type:    params.Type,
		Rate:    params.Rate,
	}); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}
