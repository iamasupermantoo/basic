package dtowallets

// DepositCreateParams 用户充值新增
type DepositCreateParams struct {
	PaymentId int     `json:"paymentId" validate:"required"`
	Money     float64 `json:"money" validate:"required,gt=0"` // 金额
	Voucher   string  `json:"voucher" validate:"required"`    // 支付凭证
}
