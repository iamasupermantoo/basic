package dtowallets

import "basic/utils"

// AssetsInfo 资产信息
type AssetsInfo struct {
	Id        int     `json:"id"`        // 主键
	AdminId   int     `json:"adminId"`   // 管理ID
	Name      string  `json:"name"`      // 名称
	Icon      string  `json:"icon"`      // 图标
	Type      int     `json:"type"`      // 类型 1法币资产 11 数字资产 21虚拟货币资产
	Rate      float64 `json:"rate"`      // 汇率
	Status    int     `json:"status"`    // 状态 -1禁用 10开启
	Data      string  `json:"data"`      // 数据
	UpdatedAt int     `json:"updatedAt"` // 更新时间
	CreatedAt int     `json:"createdAt"` // 创建时间
}

// AssetsCreateParams 资产创建
type AssetsCreateParams struct {
	AdminId   int     `json:"-"`                        // 名称
	Name      string  `json:"name" validate:"required"` // 名称
	Icon      string  `json:"icon" validate:"required"` // 图标
	Type      int     `json:"type" validate:"required"` // 类型 1法币资产 11 数字资产 21虚拟货币资产
	Rate      float64 `json:"rate"`                     // 汇率
	CreatedAt int     `json:"-"`                        // 创建时间
	UpdatedAt int     `json:"-"`                        // 更新时间
}

// AssetsUpdateParams 资产更新
type AssetsUpdateParams struct {
	Id     int     `json:"id" validate:"required"`                    // 主键
	Name   string  `json:"name"`                                      // 名称
	Icon   string  `json:"icon"`                                      // 图标
	Rate   float64 `json:"rate"`                                      // 汇率
	Status int     `json:"status"  validate:"omitempty,oneof= -1 10"` // 状态 -1禁用 10开启
	Data   string  `json:"data"`                                      // 数据
}

// AssetsIndexParams 资产列表参数
type AssetsIndexParams struct {
	AdminName  string                 `json:"adminName"`  // 管理员
	Name       string                 `json:"name"`       // 名称
	Type       int                    `json:"type"`       // 类型 1法币资产 11 数字资产 21虚拟货币资产
	Status     int                    `json:"status"`     // 状态 -1禁用 10开启
	Data       string                 `json:"data"`       // 数据
	UpdatedAt  *utils.RangeDatePicker `json:"updatedAt"`  // 更新时间
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  // 创建时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页
}

// AssetsIndexData 资产列表返回数据
type AssetsIndexData struct {
	AssetsInfo
	AdminName string `json:"adminName" gorm:"column:adminName"` // 管理员
}
