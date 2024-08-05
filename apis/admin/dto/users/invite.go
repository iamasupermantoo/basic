package dtousers

import "basic/utils"

// InviteCreateParams 新增邀请
type InviteCreateParams struct {
	UserName string `json:"username" gorm:"-"`                  // 用户名称｜管理名称
	Type     int    `json:"type" validate:"required,oneof=1 2"` // 管理1 用户2
}

// InviteIndexParams 邀请参数
type InviteIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理Id
	UserName   string                 `json:"username"`   //	用户Id
	Code       string                 `json:"code"`       //	邀请码
	Status     int                    `json:"status"`     //	状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// InviteIndexData 邀请数据
type InviteIndexData struct {
	Id        int    `json:"id"`                                //	Id
	AdminId   int    `json:"-"`                                 //	管理Id
	AdminName string `json:"adminName" gorm:"column:adminName"` //	用户名称或管理名称
	UserId    int    `json:"-"`                                 //	用户Id
	UserName  string `json:"username" gorm:"column:username"`   //	用户名称或管理名称
	Code      string `json:"code"`                              //	验证码
	Status    int    `json:"status"`                            //	状态
	Data      string `json:"data"`                              //	数据
	UpdatedAt int    `json:"updatedAt"`                         //	更新时间
	CreatedAt int    `json:"createdAt"`                         //	创造时间
}

// InviteUpdateParams 更新邀请
type InviteUpdateParams struct {
	Id        int    `json:"id" validate:"required" gorm:"-"` //	Id
	Status    int    `json:"status"`                          //	状态
	Code      string `json:"code"`                            //	邀请码
	Data      string `json:"data"`                            //	数据
	UpdatedAt int    `json:"-"`                               //	更新时间
}
