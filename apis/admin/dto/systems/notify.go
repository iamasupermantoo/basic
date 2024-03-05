package dtosystems

import "basic/utils"

// NotifyCreateParams 新增通知参数
type NotifyCreateParams struct {
	UserId    int    `json:"-"`                        //	用户Id
	AdminId   int    `json:"-"`                        //	管理Id
	UserName  string `json:"username" gorm:"-"`        //	用户名称
	Name      string `json:"name" validate:"required"` //	通知标题
	Content   string `json:"content"`                  //	内容
	UpdatedAt int    `json:"-"`                        //	更新时间
	CreatedAt int    `json:"-"`                        //	创造时间
}

// NotifyIndexParams 通知列表参数
type NotifyIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理员名
	UserName   string                 `json:"username"`   //	用户名
	Name       string                 `json:"name"`       //	通知标题
	Status     int                    `json:"status"`     //	状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// NotifyIndexData 通知列表数据
type NotifyIndexData struct {
	Id        int    `json:"id"`                                //	Id
	AdminName string `json:"adminName" gorm:"column:adminName"` //	管理员名
	UserName  string `json:"username" gorm:"column:username"`   //	管理员名
	UserId    int    `json:"userId"`                            //	用户Id
	AdminId   int    `json:"adminId"`                           //	管理Id
	Name      string `json:"name"`                              //	通知标题
	Type      int    `json:"type"`                              //	类型
	Content   string `json:"content"`                           //	内容
	Status    int    `json:"status"`                            //	状态
	Data      string `json:"data"`                              //	数据
	CreatedAt int    `json:"createdAt"`                         //	创造时间
	UpdatedAt int    `json:"updatedAt"`                         //	读取时间
}

// NotifyUpdateParams 语言更新参数
type NotifyUpdateParams struct {
	Id        int    `json:"id" validate:"required" gorm:"-"` //	Id
	Name      string `json:"name"`                            //	通知标题
	Content   string `json:"content"`                         //	内容
	Status    int    `json:"status"`                          //	状态
	Data      string `json:"data"`                            //	数据
	UpdatedAt int    `json:"-"`                               //	更新时间
}
