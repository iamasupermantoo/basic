package wallets

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/rds"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"time"
)

// UserAssetsIndex 用户资产列表
func UserAssetsIndex(adminId, userId int, lang string) (interface{}, error) {
	// 查询用户资产总金额
	data := &dtowallets.HomeUserAssetsIndexData{UserAssetsList: []*dtowallets.HomeUserAssetsInfo{}}
	if result := database.Db.Table("wallet_user_assets as wua").
		Select("sum(wua.money) as MoneySum", "sum(wua.money*wa.rate) as MoneyRateSum").
		Joins("left join wallet_assets as wa on wua.wallet_assets_id=wa.id").
		Where("wua.user_id = ?", userId).
		Where("wua.status = ?", models.WalletUserAccountStatusActive).
		Scan(data); result.Error != nil {
		return nil, result.Error
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 查询用户资产各项金额
	if result := database.Db.Table("wallet_user_assets as wua").
		Select("wua.id", "wa.id AS WalletAssetsId", "wua.money", "wa.name", " wua.money*wa.rate as MoneyRate", "wa.icon", " wua.updated_at", " wua.created_at").
		Joins("left join wallet_assets as wa on wua.wallet_assets_id=wa.id").
		Where("wua.user_id = ?", userId).
		Where("wua.wallet_assets_id <> 0").
		Where("wua.status = ?", models.WalletUserAccountStatusActive).
		Find(&data.UserAssetsList); result.Error != nil {
		return nil, result.Error
	}
	if data.UserAssetsList != nil {
		for _, v := range data.UserAssetsList {
			v.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, v.Name)
		}
	}

	// 折线图数据
	data.Echarts = &dtoadmins.Echarts{
		Category:   make([]string, 0),
		SeriesList: make([]*dtoadmins.Series, 0),
	}

	for assetsIndex, assets := range data.UserAssetsList {
		series := &dtoadmins.Series{
			Name:   assets.Name,
			Type:   "line",
			Smooth: true,
			Data:   make([]any, 0),
		}

		n := -14
		for i := n; i <= 0; i++ {
			nowTimeTmp := time.Now().AddDate(0, 0, i)
			sourceTime := time.Date(nowTimeTmp.Year(), nowTimeTmp.Month(), nowTimeTmp.Day(), 0, 0, 0, 0, time.Local)
			staTime := sourceTime.Unix()

			// 名称
			if assetsIndex == 0 {
				data.Echarts.Category = append(data.Echarts.Category, sourceTime.Format("01/02"))
			}

			endWalletBill := &models.WalletBill{}
			model := database.Db.Model(endWalletBill).Where("assets_id=?", assets.WalletAssetsId).Where("user_id=?", userId)
			if i == n {
				model.Where("created_at <= ?", staTime+86399)
			} else {
				model.Where("created_at BETWEEN ? AND ?", staTime, staTime+86399)
			}
			result := model.Order("id desc").Find(&endWalletBill)
			var endBalance float64 = 0
			if result.RowsAffected == 0 {
				if i > n {
					endBalance = series.Data[len(series.Data)-1].(float64)
				}
			} else {
				endBalance = endWalletBill.Balance + models.GetBillMoney(endWalletBill.Type, endWalletBill.Money)
			}
			series.Data = append(series.Data, endBalance)
		}

		data.Echarts.SeriesList = append(data.Echarts.SeriesList, series)
	}

	return data, nil
}

// UserAssetsInfo 用户资产信息
func UserAssetsInfo(adminId, userId int, lang string, params *dtowallets.UserAssetsInfoParams) (interface{}, error) {
	walletUserAssetsInfo := &dtowallets.HomeUserAssetsInfo{}
	if result := database.Db.Table("wallet_user_assets as wua").
		Select("wua.id", "wua.wallet_assets_id", "wua.money", "wa.rate*wua.money as MoneyRate", "wa.icon", "wa.name", "wua.updated_at", " wua.created_at").
		Joins("left join wallet_assets as wa on wua.wallet_assets_id=wa.id").
		Where("wua.status = ?", models.WalletUserAssetsStatusActive).
		Where("wua.user_id = ?", userId).
		Where("wua.wallet_assets_id = ?", params.AssetsId).
		Take(walletUserAssetsInfo); result.Error != nil {
		return nil, result.Error
	}
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	if walletUserAssetsInfo != nil {
		walletUserAssetsInfo.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, walletUserAssetsInfo.Name)
	}

	return walletUserAssetsInfo, nil
}
