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

// BillIndex 订单列表
func BillIndex(adminId, userId int, lang string, params *dtowallets.BillIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: []*dtowallets.HomeBillIndexData{}}

	// 如果类型等于 -1 或者 -2 那么是全部
	if len(params.Types) == 1 && (params.Types[0] == models.WalletBillAccountDefault || params.Types[0] == models.WalletBillAccountAssets) {
		accountList := models.GetBillAccountTypeList(params.Types[0])
		params.Types = make([]int, 0)
		for _, v := range accountList {
			params.Types = append(params.Types, v.Value)
		}
	}

	filterParams := utils.NewFilterParams().
		InInt("type IN ?", params.Types).
		EqInt("assets_id = ?", params.AssetsId).
		BetweenDate("created_at BETWEEN ? AND ?", params.DateTime)

	if result := database.Db.Table("wallet_bill").
		Select("id", "assets_id", "name", "type", "money", "balance", "created_at").
		Where("user_id = ?", userId).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	账单名称翻译
	for _, v := range data.Items.([]*dtowallets.HomeBillIndexData) {
		v.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, v.Name)
		v.Money = models.GetBillMoney(v.Type, v.Money)
	}

	return data, nil
}

// BillOptions 钱包账单Options
func BillOptions(adminId, userId int, lang string, params *dtowallets.BillOptionsParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	optionsList := models.GetBillAccountTypeList(params.Type)
	for _, v := range optionsList {
		v.Label = rds.RedisFindTranslateField(rdsConn, adminId, lang, v.Label)
	}

	return optionsList, nil
}
