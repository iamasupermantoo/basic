package dtousers

import "basic/utils"

// LevelOrderCreateParams 新增会员
type LevelOrderCreateParams struct {
	UserName string `json:"username" validate:"required" gorm:"-"` //	用户名称
	LevelId  int    `json:"levelId" validate:"required"`           //	等级Id
}

// LevelOrderIndexParams 会员参数
type LevelOrderIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理Id
	UserName   string                 `json:"username"`   //	用户Id
	LevelId    int                    `json:"levelId"`    //	等级Id
	Status     int                    `json:"status"`     //	状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	ExpiredAt  *utils.RangeDatePicker `json:"expiredAt"`  //	过期时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// LevelOrderIndexData 会员数据
type LevelOrderIndexData struct {
	Id        int    `json:"id"`                                //	Id
	AdminId   int    `json:"-"`                                 //	管理Id
	AdminName string `json:"adminName" gorm:"column:adminName"` //	用户名称或管理名称
	UserId    int    `json:"-"`                                 //	用户Id
	UserName  string `json:"username" gorm:"column:username"`   //	用户名称或管理名称
	LevelId   int    `json:"levelId"`                           //	等级Id
	Status    int    `json:"status"`                            //	状态
	Data      string `json:"data"`                              //	数据
	UpdatedAt int    `json:"updatedAt"`                         //	更新时间
	CreatedAt int    `json:"createdAt"`                         //	创造时间
	ExpiredAt int    `json:"expiredAt"`                         //	过期时间
}

// LevelOrderUpdateParams 更新会员
type LevelOrderUpdateParams struct {
	Id     int    `json:"id" validate:"required" gorm:"-"` //	Id
	Status int    `json:"status"`                          //	状态
	Data   string `json:"data"`                            //	数据
}
