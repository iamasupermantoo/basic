package dtowallets

import "basic/utils"

// DepositInfo 充值详细信息
type DepositInfo struct {
	Id        int     `json:"id"`        // 主键
	AdminId   int     `json:"adminId"`   // 管理ID
	UserId    int     `json:"userId"`    // 用户ID
	SourceId  int     `json:"sourceId"`  // 来源ID
	Type      int     `json:"type"`      // 类型
	OrderSn   string  `json:"orderSn"`   // 订单编号
	Money     float64 `json:"money"`     // 金额
	Voucher   string  `json:"voucher"`   // 支付凭证
	Fee       float64 `json:"fee"`       // 手续费
	Status    int     `json:"status"`    // 状态
	Data      string  `json:"data"`      // 数据
	UpdatedAt string  `json:"updatedAt"` // 更新时间
	CreatedAt string  `json:"createdAt"` // 更新时间
}

// DepositCreateParams 用户充值新增
type DepositCreateParams struct {
	PaymentId int     `json:"paymentId"  validate:"required,gt=0"` // 支付ID
	UserName  string  `json:"username" validate:"required"`        // 用户名
	Money     float64 `json:"money" validate:"required,gt=0"`      // 金额
}

// DepositUpdateParams 用户充值更新
type DepositUpdateParams struct {
	Id   int    `json:"id" validate:"required" gorm:"-"` // 钱包订单ID
	Data string `json:"data"`                            // 更新数据
}

// DepositIndexParams 用户充值列表
type DepositIndexParams struct {
	AdminName  string                 `json:"adminName"`  // 管理员
	UserName   string                 `json:"username"`   // 用户名
	SourceId   int                    `json:"sourceId"`   // 支付ID
	Status     int                    `json:"status"`     // 状态
	Type       int                    `json:"type"`       // 类型
	Data       string                 `json:"data"`       // 数据
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  // 申请时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页
}

// DepositIndexData 用户充值列表返回数据
type DepositIndexData struct {
	DepositInfo
	AdminName string `json:"adminName" gorm:"column:adminName"` // 管理员
	UserName  string `json:"username" gorm:"column:username"`   // 用户
}

// DepositStatusParams 用户充值更新
type DepositStatusParams struct {
	Id        int    `json:"id" validate:"required" gorm:"-"`  // 钱包订单ID
	Status    int    `json:"status" validate:"required=-1 20"` // 更新状态
	Data      string `json:"data"`                             // 更新数据
	UpdatedAt int    `json:"-"`                                //	更新时间
}
