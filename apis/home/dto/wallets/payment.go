package dtowallets

import "basic/models"

// HomePaymentIndexData 支付信息列表
type HomePaymentIndexData struct {
	PaymentInfo
	Items []*HomePaymentIndexData `json:"items"` // 子项
}

// PaymentInfo 支付信息
type PaymentInfo struct {
	Id       int                       `json:"id"`                // 主键
	AdminId  int                       `json:"-"`                 //	管理ID
	Name     string                    `json:"name"`              // 名称
	Icon     string                    `json:"icon"`              // 图标
	Type     int                       `json:"type"`              // 类型 1银行卡类型 11数字货币类型 21第三方支付
	Mode     int                       `json:"mode"`              // 模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式
	Data     string                    `json:"data"`              //	数据字符串
	DataJson *models.WalletPaymentData `json:"dataJson" gorm:"-"` //	数据
}

// PaymentIndexParams 支付信息列表参数
type PaymentIndexParams struct {
	Modes []int `json:"modes"` // 模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式
}
