package dtowallets

import "basic/utils"

// AccountInfo 帐户信息
type AccountInfo struct {
	Id        int    `json:"id"`        // 主键
	AdminId   int    `json:"adminId"`   // 管理ID
	UserId    int    `json:"userId"`    // 用户ID
	PaymentId int    `json:"paymentId"` // 支付ID
	Name      string `json:"name"`      // 名称
	RealName  string `json:"realName"`  // 真实姓名
	Number    string `json:"number"`    // 卡号
	Code      string `json:"code"`      // 识别码
	Status    int    `json:"status"`    // 状态
	Data      string `json:"data"`      // 申请理由
	UpdatedAt int    `json:"updatedAt"` // 更新时间
	CreatedAt int    `json:"createdAt"` // 创建时间
}

// AccountCreateParams 用户户新增参数
type AccountCreateParams struct {
	UserName  string `json:"userName" validate:"required"`  // 用户名
	PaymentId int    `json:"paymentId" validate:"required"` // 支付ID
	RealName  string `json:"realName" validate:"required"`  // 真实姓名
	Number    string `json:"number" validate:"required"`    // 卡号
	Code      string `json:"code"`                          // 银行代码
}

// AccountUpdateParams 用户账户修改参数
type AccountUpdateParams struct {
	Id       int    `json:"id" validate:"required" gorm:"-"` // 主键
	Name     string `json:"name"`                            // 名称
	RealName string `json:"realName"`                        // 真实姓名
	Number   string `json:"number"`                          // 卡号
	Code     string `json:"code"`                            // 识别码
	Data     string `json:"data"`                            // 数据
	Status   int    `json:"status"`                          // 状态
}

// AccountIndexParams 用户账户列表参数
type AccountIndexParams struct {
	AdminName  string                 `json:"adminName"`  // 管理员名
	UserName   string                 `json:"username"`   // 用户名
	PaymentId  int                    `json:"paymentId"`  // 支付ID
	Name       string                 `json:"name"`       // 备注
	RealName   string                 `json:"realName"`   // 姓名
	Number     string                 `json:"number"`     // 卡号
	Code       string                 `json:"code"`       // 银行代码
	Status     int                    `json:"status"`     // 状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  // 更新时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页数据
}

// AccountIndexData 用户账户管理列表数据
type AccountIndexData struct {
	AccountInfo
	AdminName string `json:"adminName" gorm:"column:adminName"` // 管理员名
	UserName  string `json:"username" gorm:"column:username"`   // 用户名
}
