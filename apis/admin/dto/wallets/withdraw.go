package dtowallets

import "basic/utils"

// WithdrawInfo 提现详细信息
type WithdrawInfo struct {
	Id        int     `json:"id"`        // 主键
	AdminId   int     `json:"adminId"`   // 管理ID
	UserId    int     `json:"userId"`    // 用户ID
	SourceId  int     `json:"sourceId"`  // 来源ID
	Type      int     `json:"type"`      // 类型
	OrderSn   string  `json:"orderSn"`   // 订单编号
	Money     float64 `json:"money"`     // 金额
	Fee       float64 `json:"fee"`       // 手续费
	Status    int     `json:"status"`    // 状态
	Data      string  `json:"data"`      // 数据
	UpdatedAt string  `json:"updatedAt"` // 更新时间
	CreatedAt string  `json:"createdAt"` // 更新时间
}

// WithdrawCreateParams 用户提现新增
type WithdrawCreateParams struct {
	PaymentId int     `json:"paymentId" validate:"required"`
	UserName  string  `json:"username" validate:"required"`   // 用户名
	Money     float64 `json:"money" validate:"required,gt=0"` // 金额
}

// WithdrawUpdateParams 用户提现更新
type WithdrawUpdateParams struct {
	Id   int    `json:"id" validate:"required" gorm:"-"` // 钱包订单ID
	Data string `json:"data"`                            // 更新数据
}

// WithdrawIndexParams 用户提现列表查询数据
type WithdrawIndexParams struct {
	AdminName  string                 `json:"adminName"`  // 管理员
	UserName   string                 `json:"username"`   // 用户名
	RealName   string                 `json:"realName"`   // 真实姓名
	Number     string                 `json:"number"`     // 卡号
	Code       string                 `json:"code"`       // 银行代码
	OrderSn    string                 `json:"orderSn"`    // 订单编号
	Status     int                    `json:"status"`     // 状态
	Type       int                    `json:"type"`       // 类型
	Data       string                 `json:"data"`       // 数据
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  // 申请时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页
}

// WithdrawIndexData 用户提现列表返回数据
type WithdrawIndexData struct {
	WithdrawInfo
	AdminName string `json:"adminName" gorm:"column:adminName"` // 管理员
	UserName  string `json:"username" gorm:"column:username"`   // 用户
	PaymentId int    `json:"paymentId" gorm:"column:paymentId"` // 钱包ID
	ReaName   string `json:"reaName" gorm:"column:realName"`    // 真实姓名
	Number    string `json:"number" gorm:"column:number"`       // 卡号
	Code      string `json:"code" gorm:"column:code"`           // 银行代码
}

// WithdrawStatusParams 用户提现审核
type WithdrawStatusParams struct {
	Id     int    `json:"id" validate:"required"` // 主键
	Status int    `json:"status" validate:"required,oneof=-1 20"`
	Data   string `json:"data"` // 更新数据
}
