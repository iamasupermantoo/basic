package dtowallets

import dtoadmins "basic/apis/admin/dto/admins"

// HomeUserAssetsIndexData 用户资产查询返回数据。
type HomeUserAssetsIndexData struct {
	MoneySum       float64               `json:"moneySum"`                // 金额
	MoneyRateSum   float64               `json:"moneyRateSum"`            // 计算利率之后的金额总和
	UserAssetsList []*HomeUserAssetsInfo `json:"userAssetsList" gorm:"-"` // 各项资产详情
	Echarts        *dtoadmins.Echarts    `json:"echarts" gorm:"-"`        // 折线图数据
}

// HomeUserAssetsInfo 用户资产详细信息。
type HomeUserAssetsInfo struct {
	Id             int     `json:"id"`             // 主键
	WalletAssetsId int     `json:"walletAssetsId"` // 钱包资产ID
	Icon           string  `json:"icon"`           //	图标
	Name           string  `json:"name"`           // 资产名
	Money          float64 `json:"money"`          // 金额
	MoneyRate      float64 `json:"moneyRate"`      // 计算后金额
	UpdatedAt      int     `json:"updatedAt"`      // 更新时间
	CreatedAt      int     `json:"createdAt"`      // 创建时间
}

// UserAssetsInfoParams 用户资产信息查询参数。
type UserAssetsInfoParams struct {
	AssetsId int `json:"assetsId"` // 资产ID
}
