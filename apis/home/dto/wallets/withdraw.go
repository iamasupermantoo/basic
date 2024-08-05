package dtowallets

// WithdrawCreateParams 用户提现新增
type WithdrawCreateParams struct {
	AccountId   int     `json:"accountId"  validate:"required,gt=0"` // 支付ID
	Money       float64 `json:"money" validate:"required,gt=0"`      // 金额
	SecurityKey string  `json:"securityKey"`                         // 安全密钥
}
