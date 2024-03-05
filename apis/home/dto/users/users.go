package dtousers

const (
	UpdatePassword         = 1
	UpdateSecurityPassword = 2
)

// UserInfo 用户信息
type UserInfo struct {
	Id          int    `json:"id"`                              //	用户ID
	ParentId    int    `json:"parentId"`                        //	上级ID
	AdminId     int    `json:"adminId"`                         //	管理ID
	UserName    string `json:"username" gorm:"column:username"` //	用户名
	NickName    string `json:"nickname" gorm:"column:nickname"` //	用户昵称
	Email       string `json:"email"`                           //	邮箱
	Telephone   string `json:"telephone"`                       //	手机号码
	Sex         int    `json:"sex"`                             //	性别
	Birthday    int    `json:"birthday"`                        //	生日
	Password    string `json:"-"`                               //	管理密码
	SecurityKey string `json:"-"`                               // 	管理密钥
	UpdatedAt   int    `json:"updatedAt"`                       //	更新时间
	CreatedAt   int    `json:"createdAt"`                       //	创建时间
}

// HomeUserInfo 前台用户信息
type HomeUserInfo struct {
	Id         int     `json:"id"`                                  //	用户ID
	UserName   string  `json:"username" gorm:"column:username"`     //  用户名
	NickName   string  `json:"nickname" gorm:"column:nickname"`     //	用户昵称
	Email      string  `json:"email"`                               //	邮箱
	Telephone  string  `json:"telephone"`                           //	手机号码
	Avatar     string  `json:"avatar"`                              //	用户头像
	Sex        int     `json:"sex"`                                 //	性别
	Money      float64 `json:"money"`                               //	余额
	Score      int     `json:"score"`                               //  信用分
	Level      int     `json:"level"`                               // 	等级
	AuthStatus int     `json:"authStatus" gorm:"column:authStatus"` //  实名状态
	Birthday   int     `json:"birthday"`                            //	生日
	Desc       string  `json:"desc"`                                //	个性签名
	CreatedAt  int     `json:"createdAt"`                           //注册时间
}

// UserRegisterAward 注册邀请奖励
type UserRegisterAward struct {
	Register float64 `json:"register"` //	注册奖励
	Share    float64 `json:"share"`    //	分享奖励
}

// UserLoginParams 用户登录参数
type UserLoginParams struct {
	AdminId    int    `json:"-"`                                         //	管理ID
	UserName   string `json:"username" validate:"required,min=5,max=20"` //	用户名
	Password   string `json:"password" validate:"required,min=5,max=20"` //	密码
	CaptchaId  string `json:"captchaId"`                                 //	验证码ID
	CaptchaVal string `json:"captchaVal"`                                //	验证码值
}

// UserRegisterParams 用户注册参数
type UserRegisterParams struct {
	UserLoginParams
	Email       string `json:"email"`       //	邮箱
	Telephone   string `json:"telephone"`   //	手机号码
	SecurityKey string `json:"securityKey"` //	安全密钥
	Code        string `json:"code"`        //	用户邀请码
}

// UserLoginData 登录数据
type UserLoginData struct {
	Token    string      `json:"token"`
	UserInfo interface{} `json:"userInfo"`
}

// UserUpdateParams 前台用户更新参数
type UserUpdateParams struct {
	Avatar      string `json:"avatar"`                             //	用户头像
	NickName    string `json:"nickname" gorm:"column:nickname"`    //	用户昵称
	Email       string `json:"email"`                              //	用户邮箱
	Telephone   string `json:"telephone"`                          //	用户手机号码
	Sex         int    `json:"sex" validate:"omitempty,oneof=1 2"` //	性别
	Birthday    int    `json:"-" `                                 //	生日
	Desc        string `json:"desc"`                               //	个性签名
	BirthdayStr string `json:"birthdayStr" gorm:"-"`               //	生日Str
}

// UserUpdatePasswordParams 前台用户更新密码||安全码参数
type UserUpdatePasswordParams struct {
	Type        int64  `json:"type" validate:"required,oneof=1 2"` //	更新类型 1更新登录密码 2更新安全密钥
	OldPassword string `json:"oldPassword"`                        //	旧密码
	NewPassword string `json:"newPassword" validate:"required"`    //	新密码
}
