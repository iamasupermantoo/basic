package dtousers

import (
	"basic/utils"
)

// UserIndexParams 用户列表参数
type UserIndexParams struct {
	AdminName  string                 `json:"adminName"`                               //	管理名称
	ParentName string                 `json:"parentName"`                              //	父级用户名
	UserName   string                 `json:"username"`                                //	用户名
	NickName   string                 `json:"nickname"`                                //	用户昵称
	Email      string                 `json:"email"`                                   //	用户邮箱
	Telephone  string                 `json:"telephone"`                               //	用户手机号码
	Sex        int                    `json:"sex" validate:"omitempty,oneof=1 2"`      //	男女
	Birthday   *utils.RangeDatePicker `json:"birthday"`                                //	用户生日时间类型
	Type       int                    `json:"type" validate:"omitempty,oneof=1 11 21"` //	类型
	Status     int                    `json:"status" validate:"omitempty,oneof=10 -1"` //	状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`                               //	注册时间范围
	Pagination *utils.Pagination      `json:"pagination"`                              //	分页信息
}

//// UserLoginParams 用户登录参数
//type UserLoginParams struct {
//	AdminId    int    `json:"-"`                                         //	管理ID
//	UserName   string `json:"username" validate:"required,min=5,max=20"` //	用户名
//	Password   string `json:"password" validate:"required,min=5,max=20"` //	密码
//	CaptchaId  string `json:"captchaId"`                                 //	验证码ID
//	CaptchaVal string `json:"captchaVal"`                                //	验证码值
//}

// UserCreateParams 用户创建参数
type UserCreateParams struct {
	//UserLoginParams
	Email       string `json:"email"`       //	邮箱
	Telephone   string `json:"telephone"`   //	手机号码
	SecurityKey string `json:"securityKey"` //	安全密钥
	Code        string `json:"code"`        //	用户邀请码

	AdminId    int    `json:"-"`                                         //	管理ID
	UserName   string `json:"username" validate:"required,min=5,max=20"` //	用户名
	Password   string `json:"password" validate:"required,min=5,max=20"` //	密码
	CaptchaId  string `json:"captchaId"`                                 //	验证码ID
	CaptchaVal string `json:"captchaVal"`                                //	验证码值
}

// UserUpdateParams 用户更新数据
type UserUpdateParams struct {
	Id          int    `json:"id" validate:"required" gorm:"-"`            //	用户ID
	Avatar      string `json:"avatar"`                                     //	用户头像
	NickName    string `json:"nickname" gorm:"column:nickname"`            //	用户昵称
	Email       string `json:"email"`                                      //	用户邮箱
	Telephone   string `json:"telephone"`                                  //	用户手机号码
	Sex         int    `json:"sex" validate:"omitempty,oneof=1 2"`         //	性别
	Birthday    string `json:"birthday" gorm:"-"`                          //	用户生日时间类型
	Type        int    `json:"type" validate:"omitempty,oneof=1 11 21"`    //	类型
	Score       int    `json:"score"`                                      //	信用分
	Status      int    `json:"status" validate:"omitempty,oneof=10 -1 -2"` //	状态
	Data        string `json:"data"`                                       //	数据
	Desc        string `json:"desc"`                                       //	个性签名
	BirthdayInt int    `gorm:"column:birthday"`                            //	生日
}

// UserBalanceParam 用户余额变动
type UserBalanceParam struct {
	UserName string  `json:"username" validate:"required"`        //	用户名
	Type     int     `json:"type" validate:"required,oneof=3 13"` //	类型
	Money    float64 `json:"money" validate:"required,gt=0"`      //	金额
}

// UserIndexData 用户列表
type UserIndexData struct {
	UserInfo
	AdminName  string  `json:"adminName" gorm:"column:adminName"`    //	管理名称
	ParentName string  `json:"parentName" gorm:"column:parentName;"` //	上级用户名
	Avatar     string  `json:"avatar"`                               //	用户头像
	Money      float64 `json:"money"`                                //	金额
	Score      int     `json:"score"`                                //	信用分
	Type       int     `json:"type"`                                 //	类型
	Status     int     `json:"status"`                               //	状态
	Data       string  `json:"data"`                                 //	数据
	Desc       string  `json:"desc"`                                 //	个性签名
}

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
	SecurityKey string `json:"-"`                               // 管理密钥
	UpdatedAt   int    `json:"updatedAt"`                       //	更新时间
	CreatedAt   int    `json:"createdAt"`                       //	创建时间
}

// UserRelationTree 用户关系树
type UserRelationTree struct {
	Id          int                 `json:"id"`                              //	用户ID
	Header      string              `json:"header"`                          //	层级
	Avatar      string              `json:"avatar"`                          //	头像
	UserName    string              `json:"username" gorm:"column:username"` //	用户名
	NickName    string              `json:"nickname" gorm:"column:nickname"` //	用户昵称
	Email       string              `json:"email"`                           //	邮箱
	Telephone   string              `json:"telephone"`                       //	手机号码
	SumPeople   int                 `json:"sumPeople"`                       //	总人数
	SumDeposit  float64             `json:"sumDeposit"`                      //	总充值
	SumWithdraw float64             `json:"sumWithdraw"`                     //	总提现
	SumEarnings float64             `json:"sumEarnings"`                     //	总收益
	UpdatedAt   int                 `json:"updatedAt"`                       //	更新时间
	CreatedAt   int                 `json:"createdAt"`                       //	创建时间
	Children    []*UserRelationTree `json:"children" gorm:"-"`               //	子级
}
