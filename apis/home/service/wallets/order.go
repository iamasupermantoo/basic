package wallets

import (
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
)

// OrderIndex 钱包订单列表
func OrderIndex(adminId, userId int, lang string, params *dtowallets.OrderIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: []*dtowallets.HomeWalletOrderIndexData{}}
	filterParams := utils.NewFilterParams().
		EqInt("assets_id = ?", params.AssetsId).
		InInt("type in ?", params.Types)
	if result := database.Db.Table("wallet_order as wo").
		Select("wo.id", "wo.type", "wa.icon", "wo.money", "wo.status", "wo.data", "wo.updated_at").
		Joins("left join wallet_assets as wa on wa.id = wo.assets_id").
		Where("wo.user_id = ?", userId).
		Where("wo.status > ?", models.WalletOrderStatusDelete).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	billTypeList := models.GetBillTypeList()
	for _, v := range data.Items.([]*dtowallets.HomeWalletOrderIndexData) {
		v.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, billTypeList[v.Type])
	}

	return data, nil
}
