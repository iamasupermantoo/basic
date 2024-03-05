package dtoadmins

import (
	"basic/utils"
)

const (
	UpdatePassword         = 1
	UpdateSecurityPassword = 2
)

// AdminLoginParams 管理登录参数
type AdminLoginParams struct {
	UserName   string `json:"username" validate:"required,min=5,max=28"` //	用户名
	Password   string `json:"password" validate:"required,min=5,max=28"` //	密码
	CaptchaId  string `json:"captchaId" validate:"required"`             //	验证码ID
	CaptchaVal string `json:"captchaVal" validate:"required,len=4"`      //	验证码值
}

// AdminLoginData 管理登录信息
type AdminLoginData struct {
	Token string `json:"token"` //	Token信息
}

// AdminInitData 管理初始化数据
type AdminInitData struct {
	Config     *AdminInitConfig `json:"config"`     //	配置
	RouterList []string         `json:"routerList"` //	管理菜单
	MenuList   []*MenuInfo      `json:"menuList"`   //	管理菜单
	AdminInfo  *AdminInfo       `json:"adminInfo"`  //	管理信息
}

// AdminInitConfig 管理信息配置
type AdminInitConfig struct {
	Name string `json:"name"` //	项目名称
	Logo string `json:"logo"` //	项目Logo
}

// AdminInfo 管理信息
type AdminInfo struct {
	Id          int     `json:"id"`                              //	管理ID
	ParentId    int     `json:"parentId"`                        //	上级ID
	UserName    string  `gorm:"column:username" json:"username"` //	管理账号
	NickName    string  `gorm:"column:nickname" json:"nickname"` //	管理昵称
	Email       string  `json:"email"`                           //	管理邮箱
	Avatar      string  `json:"avatar"`                          //	管理头像
	Money       float64 `json:"money"`                           //	管理金额
	Password    string  `json:"-"`                               //	管理密码
	SecurityKey string  `json:"-"`                               // 管理密钥
	Status      int     `json:"status"`                          //状态
	Domains     string  `json:"domains"`                         //域名
	ExpiredAt   int     `json:"expiredAt"`                       //	过期时间
	UpdatedAt   int     `json:"updatedAt"`                       //	更新时间
	CreatedAt   int     `json:"createdAt"`                       //	创建时间
}

// AdminInfoData 管理信息数据
type AdminInfoData struct {
	Template string `json:"template"` //	管理员模版
	Nums     int    `json:"nums"`     //	代理数量
}

// AdminUpdateParams 更新当前管理员参数
type AdminUpdateParams struct {
	Id       int    `json:"id" gorm:"-"`                     //	管理ID
	Email    string `json:"email"`                           // 管理邮箱
	NickName string `json:"nickname" gorm:"column:nickname"` // 管理昵称
	Avatar   string `json:"avatar"`                          // 管理头像
}

// AdminUpdateDesc 更新商户配置
type AdminUpdateDesc struct {
	Id       int            `json:"id"`
	DataJson *AdminInfoData `json:"dataJson"` //	管理配置
}

// AdminUpdatePasswordParams 更新当前管理员密码参数
type AdminUpdatePasswordParams struct {
	Id          int    `gorm:"-"`                                            //	管理ID
	Type        int64  `json:"type" validate:"required,oneof=1 2"`           //	更新类型 1更新登录密码 2更新安全密钥
	OldPassword string `json:"oldPassword" validate:"required,min=5,max=28"` //	旧密码
	NewPassword string `json:"newPassword" validate:"required,min=5,max=28"` //	新密码
	CmfPassword string `json:"cmfPassword" validate:"required,min=5,max=28"` //	确认密码
}

type Statistics struct {
	Name      string `json:"name"`      //	名称
	Icon      string `json:"icon"`      //	图标
	Color     string `json:"color"`     //	背景颜色
	Today     any    `json:"today"`     //	今日
	Yesterday any    `json:"yesterday"` //	昨日
	Total     any    `json:"total"`     //	总数
}

type Series struct {
	Name   string `json:"name"`   //	名称
	Type   string `json:"type"`   //	线类型
	Smooth bool   `json:"smooth"` // 	平滑
	Data   []any  `json:"data"`   //	数据
}

type Echarts struct {
	Category   []string  `json:"category"` //	日期
	SeriesList []*Series `json:"series"`   //	数据
}

// IndexPageData 首页信息数据
type IndexPageData struct {
	Items       [][]*Statistics `json:"items"`   //	图标信息
	EchartsList *Echarts        `json:"echarts"` //	图标数据
}

// AdminLogsIndexParams 管理日志列表参数
type AdminLogsIndexParams struct {
	AdminName  string                 `json:"adminName"`  //管理名称
	Name       string                 `json:"name"`       //操作名称
	Route      string                 `json:"route"`      //路由
	Ip         string                 `json:"ip"`         //IP地址
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //操作时间
	Pagination *utils.Pagination      `json:"pagination"` //分页设置
}

// AdminLogsIndexData 管理日志列表信息
type AdminLogsIndexData struct {
	Id        int    `json:"id"`                                //ID
	AdminName string `json:"adminName" gorm:"column:adminName"` //管理ID
	Name      string `json:"name"`                              //操作名称
	Ip        string `json:"ip"        gorm:"column:ip"`        //IP地址
	Headers   string `json:"headers"`                           //头信息
	Route     string `json:"route"`                             //路由
	Body      string `json:"body"`                              //内容
	Data      string `json:"data"`                              //数据
	CreatedAt int    `json:"createdAt"`                         //操作时间
}

// ManageIndexParams 管理列表参数
type ManageIndexParams struct {
	ParentName string                 `json:"parentName" ` //上级名称
	Username   string                 `json:"username"`    //管理名称
	Nickname   string                 `json:"nickname"`    //昵称
	Email      string                 `json:"email"`       //邮箱
	Domains    string                 `json:"domains"`     //域名
	Role       string                 `json:"role"`        //角色
	Status     int                    `json:"status"`      //状态
	Pagination *utils.Pagination      `json:"pagination"`  //分页设置
	ExpiredAt  *utils.RangeDatePicker `json:"expiredAt"`   //过期时间
}

// AdminIndexData 管理列表信息
type AdminIndexData struct {
	ParentName string          `json:"parentName" gorm:"column:parentName"` //父级名称
	RolesStr   string          `json:"rolesStr" gorm:"-"`                   //当前角色列表
	Roles      map[string]bool `json:"roles" gorm:"-"`                      //选中角色对列表
	Data       string          `json:"data"`                                // 数据
	DataJson   *AdminInfoData  `json:"dataJson" gorm:"-"`                   //管理详情
	AdminInfo
}

// ManageCreateParams 新增管理参数
type ManageCreateParams struct {
	Roles     map[string]bool `json:"roles" validate:"required"`                 //选中角色对列表
	UserName  string          `json:"username" validate:"required"`              //管理名称
	Domains   string          `json:"domains"`                                   //域名
	Password  string          `json:"password" validate:"required,min=5,max=28"` //密码
	ExpiredAt string          `json:"expiredAt"`                                 //过期时间
}

// ManageUpdateParams 管理更新参数
type ManageUpdateParams struct {
	Id          int             `json:"id" validate:"required" gorm:"-"`
	Roles       map[string]bool `json:"roles" gorm:"-"`                          //选中角色列表
	Domains     string          `json:"domains"`                                 //域名
	Email       string          `json:"email"`                                   //邮箱
	Nickname    string          `json:"nickname"`                                //昵称
	Avatar      string          `json:"avatar"`                                  //头像
	Password    string          `json:"password"`                                //密码
	SecurityKey string          `json:"securityKey"`                             //安全码
	Status      int             `json:"status" validate:"omitempty,oneof=-1 10"` //状态
	ExpiredAt   int             `json:"-"`                                       //过期时间
	ExpiredTime string          `json:"expiredTime" gorm:"-"`                    //过期时间
}

// MerchantParams 商户管理参数
type MerchantParams struct {
	Id int `json:"id" validate:"gt=1"` // 管理员ID
}
