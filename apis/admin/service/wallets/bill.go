package wallets

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
)

// BillIndex 用户账单列表
func BillIndex(adminId, userId int, params *dtowallets.BillIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtowallets.BillIndexData, 0)}
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		InInt("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		EqInt("assets_id = ?", params.AssetsId).
		EqInt("type = ?", params.Type).
		BetweenDate("created_at BETWEEN ? AND ?", params.DateTime)

	if result := database.Db.Table("wallet_bill AS wb").
		Select("wb.id", "wb.admin_id", "wb.assets_id", "wb.user_id", "au.username as adminName", "u.username", "wb.type", "wb.money", "wb.balance", "wb.data", "wb.created_at").
		Joins("left JOIN admin_user AS au ON au.id = wb.admin_id").
		Joins("left JOIN user AS u ON u.id = wb.user_id").
		Joins("left JOIN wallet_assets AS wa ON wa.id = wb.assets_id").
		Where("wb.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
