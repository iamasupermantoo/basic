package dtowallets

import "basic/utils"

// BillIndexParams 用户账单列表查询数据
type BillIndexParams struct {
	AssetsId   int                    `json:"assetsId"`                       // 资产ID
	Types      []int                  `json:"types" validate:"required"`      // 类型
	DateTime   *utils.RangeDatePicker `json:"createdAt"`                      // 创建时间
	Pagination *utils.Pagination      `json:"pagination" validate:"required"` // 分页数据
}

// BillOptionsParams 账单Options参数
type BillOptionsParams struct {
	Type int `json:"type" validate:"required"`
}

// HomeBillIndexData 用户账单列表返回数据
type HomeBillIndexData struct {
	Id        int     `json:"id"`
	Type      int     `json:"type"`      // 类型 1充值类型 11提现类型 21购买 51收益 61奖励
	Name      string  `json:"name"`      // 名称
	Money     float64 `json:"money"`     // 金额
	Balance   float64 `json:"balance"`   // 余额
	CreatedAt int     `json:"createdAt"` // 创建时间
}
