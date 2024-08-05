package dtowallets

import (
	"basic/models"
	"basic/utils"
)

// PaymentInfo 钱包支付信息
type PaymentInfo struct {
	Id        int    `json:"id"`        // 主键
	AdminId   int    `json:"adminId"`   // 管理ID
	AssetsId  int    `json:"assetsId"`  // 钱包资产ID
	Name      string `json:"name"`      // 名称
	Icon      string `json:"icon"`      // 图标
	Type      int    `json:"type"`      // 类型 1银行卡类型 11数字货币类型 21第三方支付
	Mode      int    `json:"mode"`      // 模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式
	Level     int    `json:"level"`     // 等级
	Status    int    `json:"status"`    // 状态 -2删除 -1禁用 10开启
	Data      string `json:"data"`      // 数据
	Desc      string `json:"desc"`      // 详情
	UpdatedAt int    `json:"updatedAt"` //	更新时间
	CreatedAt int    `json:"createdAt"` //	更新时间
}

// PaymentCreateParams  钱包支付创建
type PaymentCreateParams struct {
	AdminId  int    `json:"-"`
	Icon     string `json:"icon" validate:"required"`  // 图标
	Name     string `json:"name" validate:"required"`  // 名称
	Mode     int    `json:"mode" validate:"required"`  // 模式
	Type     int    `json:"type" validate:"required"`  // 类型
	AssetsId int    `json:"assetsId" validate:"gt=-1"` // 钱包资产ID
}

// PaymentIndexParams  钱包支付查询
type PaymentIndexParams struct {
	AdminName  string                 `json:"adminName"`  // 管理员
	Name       string                 `json:"name"`       // 名称
	Data       string                 `json:"data"`       // 数据
	Desc       string                 `json:"desc"`       // 详情
	AssetsId   int                    `json:"assetsId"`   // 资产ID
	Status     int                    `json:"status"`     // 状态 -1禁用 10开启
	Mode       int                    `json:"mode"`       // 模式
	Type       int                    `json:"type"`       // 类型
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  // 创建时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页数据
}

// PaymentIndexData  钱包支付查询返回数据
type PaymentIndexData struct {
	PaymentInfo
	DataJson  *models.WalletPaymentData `json:"dataJson" gorm:"-"`                 // 信息数据
	AdminName string                    `json:"adminName" gorm:"column:adminName"` // 管理员
}

// PaymentUpdateParams 钱包支付修改
type PaymentUpdateParams struct {
	Id     int    `json:"id" validate:"required" gorm:"-"`          // 主键
	Name   string `json:"name"`                                     // 名称
	Status int    `json:"status" validate:"omitempty,oneof= -1 10"` // 状态
	Icon   string `json:"icon"`                                     // 图标
	Data   string `json:"data"`                                     // 数据
	Desc   string `json:"desc"`                                     // 详情
}

// PaymentDescParams 数据参数
type PaymentDescParams struct {
	Id       int                       `json:"id" validate:"required"` // 主键
	DataJson *models.WalletPaymentData `json:"dataJson"`               // 数据
}
