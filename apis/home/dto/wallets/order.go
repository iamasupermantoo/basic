package dtowallets

import "basic/utils"

// OrderIndexParams  钱包订单列表参数
type OrderIndexParams struct {
	AssetsId   int               `json:"assetsId"` //	资产ID
	Types      []int             `json:"types"`
	Pagination *utils.Pagination `json:"pagination"` // 分页数据
}

// HomeWalletOrderIndexData  钱包订单列表返回数据
type HomeWalletOrderIndexData struct {
	Id        int     `json:"id"`        // 主键
	Icon      string  `json:"icon"`      //	图标
	Name      string  `json:"name"`      //	名称
	Type      int     `json:"type"`      //	类型 1充值类型 2资产充值类型 11提现类型 12资产提现类型
	Money     float64 `json:"money"`     // 金额
	Fee       float64 `json:"fee"`       // 手续费
	Status    int     `json:"status"`    // 状态
	Data      string  `json:"data"`      //	拒绝理由
	UpdatedAt string  `json:"updatedAt"` // 更新时间
}
