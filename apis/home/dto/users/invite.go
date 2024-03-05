package dtousers

// InviteInfo 邀请信息
type InviteInfo struct {
	Code string `json:"code"` //	邀请码
}

// InviteCreateParams 新增邀请
type InviteCreateParams struct {
	UserName string `json:"username" gorm:"-"`                  // 用户名称｜管理名称
	Type     int    `json:"type" validate:"required,oneof=1 2"` // 管理1 用户2
}
