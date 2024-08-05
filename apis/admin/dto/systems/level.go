package dtosystems

import "basic/utils"

// LevelCreateParams 新增等级参数
type LevelCreateParams struct {
	AdminId   int     `json:"-"`                         //	管理Id
	Level     int     `json:"level"`                     //	等级标识
	Name      string  `json:"name" validate:"required"`  //	名称
	Icon      string  `json:"icon" validate:"required"`  //	图标
	Money     float64 `json:"money" validate:"required"` //	金额
	Days      int     `json:"days" validate:"required"`  //	天数
	Desc      string  `json:"desc"`                      //	详情
	UpdatedAt int     //	更新时间
	CreatedAt int     //	创建时间
}

// LevelIndexParams 等级列表数据
type LevelIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理Id
	Name       string                 `json:"name"`       //	名称
	Status     int                    `json:"status"`     //	状态
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// LevelIndexData 等级数据
type LevelIndexData struct {
	Id        int     `json:"id"`                                //	Id
	AdminId   int     `json:"adminId"`                           //	管理Id
	AdminName string  `json:"adminName" gorm:"column:adminName"` //	管理员名称
	Level     int     `json:"level"`                             //	等级标识
	Name      string  `json:"name"`                              //	名称
	Icon      string  `json:"icon"`                              //	图标
	Money     float64 `json:"money"`                             //	金额
	Days      int     `json:"days"`                              //	天数
	Status    int     `json:"status"`                            //	状态
	Data      string  `json:"data"`                              //	数据
	Desc      string  `json:"desc"`                              //	详情
	CreatedAt int     `json:"createdAt"`                         //	创造时间
}

// LevelUpdateParams 等级更新参数
type LevelUpdateParams struct {
	Id     int    `json:"id" validate:"required" gorm:"-"` //	主键
	Level  int    `json:"level"`                           //	等级标识
	Name   string `json:"name"`                            //	名称
	Icon   string `json:"icon"`                            //	图标
	Days   int    `json:"days"`                            //	天数
	Status int    `json:"status"`                          //	状态
	Data   string `json:"data"`                            //	数据
	Desc   string `json:"desc"`                            //	详情
}
