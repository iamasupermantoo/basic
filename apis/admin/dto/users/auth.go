package dtousers

import "basic/utils"

// AuthCreateParams 新增认证
type AuthCreateParams struct {
	AdminId   int    `json:"-"`
	UserId    int    `json:"-"`
	UserName  string `json:"username" gorm:"-"`                 //	用户名称
	RealName  string `json:"realName" validate:"required"`      //	真实姓名
	Number    string `json:"number" validate:"required"`        //	卡号
	Photo1    string `json:"photo1" validate:"required"`        //	证件照1
	Photo2    string `json:"photo2" validate:"required"`        //	证件照2
	Photo3    string `json:"photo3"`                            //	证件照3
	Telephone string `json:"telephone"`                         //	手机号码
	Address   string `json:"address"`                           //	详细地址
	Type      int    `json:"type"  validate:"required,oneof=1"` //	类型
	Status    int    `json:"-"`                                 //	状态
	UpdatedAt int    `json:"-"`                                 //	更新时间
	CreatedAt int    `json:"-"`                                 //	创造时间
}

// AuthIndexParams 认证参数
type AuthIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理Id
	UserName   string                 `json:"username"`   //	用户Id
	RealName   string                 `json:"realName"`   //	真实姓名
	Address    string                 `json:"address"`    //	详细地址
	Type       int                    `json:"type"`       //	类型
	Status     int                    `json:"status"`     //	状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// AuthIndexData 认证数据
type AuthIndexData struct {
	Id        int    `json:"id"`                                //	Id
	AdminId   int    `json:"-"`                                 //	管理Id
	AdminName string `json:"adminName" gorm:"column:adminName"` //	管理名称
	UserId    int    `json:"-"`                                 //	用户Id
	UserName  string `json:"username" gorm:"column:username"`   //	用户名称
	RealName  string `json:"realName"`                          //	真实姓名
	Number    string `json:"number"`                            //	卡号
	Photo1    string `json:"photo1"`                            //	证件照1
	Photo2    string `json:"photo2"`                            //	证件照2
	Photo3    string `json:"photo3"`                            //	证件照3
	Telephone string `json:"telephone"`                         //	手机号码
	Address   string `json:"address"`                           //	详细地址
	Type      int    `json:"type"`                              //	类型
	Status    int    `json:"status"`                            //	状态
	Data      string `json:"data"`                              //	数据
	UpdatedAt int    `json:"updatedAt"`                         //	更新时间
	CreatedAt int    `json:"createdAt"`                         //	创造时间
}

// AuthUpdateParams 更新认证
type AuthUpdateParams struct {
	Id        int    `json:"id" validate:"required" gorm:"-"` //	Id
	RealName  string `json:"realName"`                        //	真实姓名
	Number    string `json:"number"`                          //	卡号
	Photo1    string `json:"photo1"`                          //	证件照1
	Photo2    string `json:"photo2"`                          //	证件照2
	Photo3    string `json:"photo3"`                          //	证件照3
	Telephone string `json:"telephone"`                       //	手机号码
	Address   string `json:"address"`                         //	详细地址
	Type      int    `json:"type"`                            //	类型
	Status    int    `json:"status"`                          //	状态
	Data      string `json:"data"`                            //	数据
}
