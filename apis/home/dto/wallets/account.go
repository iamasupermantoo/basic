package dtowallets

// HomeAccountInfo 卡片信息
type HomeAccountInfo struct {
	Id          int    `json:"id"`          // 	主键
	Name        string `json:"name"`        //	类型名称
	PaymentId   int    `json:"paymentId"`   //	支付ID
	PaymentType int    `json:"paymentType"` //	支付类型
	PaymentName string `json:"paymentName"` //	支付名称
	Icon        string `json:"icon"`        // 	图标
	RealName    string `json:"realName"`    // 	真实姓名
	Number      string `json:"number"`      // 	卡号
	Code        string `json:"code"`        // 	银行代码
	Status      int    `json:"status"`      // 	状态
	CreatedAt   int    `json:"createdAt"`   //	创建时间
}

// HomeAccountIndexParams 提现账户
type HomeAccountIndexParams struct {
	Modes []int `json:"modes"` // 模式 1资金模式 2资产模式
}

// HomeAccountIndexData 卡片管理列表返回数据
type HomeAccountIndexData struct {
	AccountList []*HomeAccountInfo `json:"accountList"` // 卡片信息数组
}

// AccountUpdateParams 编辑卡片
type AccountUpdateParams struct {
	Id          int    `json:"id" validate:"required"` // 主键
	RealName    string `json:"realName"`               // 真实姓名
	Number      string `json:"number"`                 // 卡号
	Code        string `json:"code"`                   // 银行代码｜token
	SecurityKey string `json:"securityKey"`            // 安全密钥
}

// AccountCreateParams 新增卡片
type AccountCreateParams struct {
	PaymentId int    `json:"paymentId" validate:"required,gt=0"` // 支付ID
	RealName  string `json:"realName"`                           // 真实姓名
	Number    string `json:"number" validate:"required"`         // 卡号｜token
	Code      string `json:"code"`                               // 银行代码｜token
}

// HomeAccountDeleteParams 删除卡片
type HomeAccountDeleteParams struct {
	Id          int    `json:"id" validate:"required"` // 主键
	SecurityKey string `json:"securityKey"`            // 安全码
}

// AccountInfoParams 卡片详情
type AccountInfoParams struct {
	Id int `json:"id" validate:"required"` // 主键
}
