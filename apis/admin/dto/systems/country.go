package dtosystems

import "basic/utils"

// CountryInfo 国家信息
type CountryInfo struct {
	AdminId   int    `json:"adminId"`   //	管理Id
	Id        int    `json:"id"`        //	主键
	Name      string `json:"name"`      //	名称
	Alias     string `json:"alias"`     //	别名
	Iso1      string `json:"iso1"`      //	二位代码
	Icon      string `json:"icon"`      //	图标
	Code      string `json:"code"`      //	区号
	Sort      int    `json:"sort"`      //	排序
	Status    int    `json:"status"`    //	状态
	Data      string `json:"data"`      //	数据
	CreatedAt int    `json:"createdAt"` //	创造时间
}

// CountryIndexData 国家列表数据
type CountryIndexData struct {
	CountryInfo
	AdminName string ` json:"adminName" gorm:"column:adminName"` //	管理员名
}

// CountryCreateParams 新增国家参数
type CountryCreateParams struct {
	Id        int    `json:"-"`                         //	新增ID
	AdminId   int    `json:"-"`                         //	管理Id
	Name      string `json:"name" validate:"required"`  //	名称
	Alias     string `json:"alias" validate:"required"` //	别名
	Iso1      string `json:"iso1" validate:"required"`  //	二位代码
	Icon      string `json:"icon" validate:"required"`  //	图标
	Code      string `json:"code" validate:"required"`  //	区号
	Sort      int    `gorm:"default:99"`
	Data      string //	数据
	UpdatedAt int    //	更新时间
	CreatedAt int    //	创建时间
}

// CountryIndexParams 国家列表参数
type CountryIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理Id
	Name       string                 `json:"name"`       //	名称
	Alias      string                 `json:"alias"`      //	别名
	Icon       string                 `json:"icon"`       //	图标
	Iso1       string                 `json:"iso1"`       //	二位代码
	Sort       int                    `json:"sort"`       //	排序
	Code       string                 `json:"code"`       //	区号
	Status     int                    `json:"status"`     //	状态
	Data       string                 `json:"data"`       //	数据
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// CountryUpdateParams 国家更新参数
type CountryUpdateParams struct {
	Id     int    `json:"id" validate:"required" gorm:"-"` //	Id
	Name   string `json:"name"`                            //	名称
	Alias  string `json:"alias"`                           //	别名
	Code   string `json:"code"`                            //	区号
	Icon   string `json:"icon"`                            //	图标
	Iso1   string `json:"iso1"`                            //	二位代码
	Sort   int    `json:"sort"`                            //	排序
	Status int    `json:"status"`                          //	状态
	Data   string `json:"data"`                            //	数据
}
