package dtosystems

import "basic/utils"

// LangCreateParams 新增语言参数
type LangCreateParams struct {
	Id        int    `json:"-"`
	AdminId   int    `json:"-"`                         //	管理Id
	Name      string `json:"name" validate:"required"`  //	名称
	Alias     string `json:"alias" validate:"required"` //	别名
	Icon      string `json:"icon" validate:"required"`  //	图标
	Sort      int    `json:"-" gorm:"default:99"`       //	排序
	UpdatedAt int    `json:"-"`                         //	更新时间
	CreatedAt int    `json:"-"`                         //	创造时间
}

// LangIndexParams 语言列表参数
type LangIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理Id
	Name       string                 `json:"name"`       //	名称
	Alias      string                 `json:"alias"`      //	别名
	Status     int                    `json:"status"`     //	状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// LangUpdateParams 语言更新参数
type LangUpdateParams struct {
	Id     int    `json:"id" validate:"required" gorm:"-"` //	Id
	Name   string `json:"name"`                            //	名称
	Alias  string `json:"alias"`                           //	别名
	Icon   string `json:"icon"`                            //	图标
	Sort   int    `json:"sort"`                            //	排序
	Status int    `json:"status"`                          //	状态
	Data   string `json:"data"`                            //	数据
}

// LangIndexData 语言列表数据
type LangIndexData struct {
	Id        int    `json:"id"`                                //	Id
	AdminName string `json:"adminName" gorm:"column:adminName"` //	管理员名
	AdminId   int    `json:"adminId"`                           //	管理Id
	Name      string `json:"name"`                              //	名称
	Alias     string `json:"alias"`                             //	别名
	Icon      string `json:"icon"`                              //	图标
	Sort      int    `json:"sort"`                              //	排序
	Status    int    `json:"status"`                            //	状态
	Data      string `json:"data"`                              //	数据
	CreatedAt int    `json:"createdAt"`                         //	创造时间
}
