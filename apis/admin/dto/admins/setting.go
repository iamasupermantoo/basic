package dtoadmins

import "basic/utils"

// AdminSettingTeamEarningSetting 团队收益配置
type AdminSettingTeamEarningSetting struct {
	Options map[string]bool `json:"options"` //	账单功能
	Rate    float64         `json:"rate"`    //	收益率
}

// AdminSettingWithdrawSetting 提现次数手续费配置
type AdminSettingWithdrawSetting struct {
	Days int     `json:"days"` //	天数
	Nums int     `json:"nums"` //	次数
	Fee  float64 `json:"fee"`  //	手续费
}

// AdminSettingMenuIndexData 菜单列表信息
type AdminSettingMenuIndexData struct {
	Id         int            `json:"id"`                                  //ID
	ParentId   int            `json:"parentId"`                            //上级ID
	AdminName  string         `json:"adminName" gorm:"column:adminName"`   //管理
	ParentName string         `json:"parentName" gorm:"column:parentName"` //父级菜单
	Name       string         `json:"name"`                                //菜单名称
	Route      string         `json:"route"`                               //路由
	Sort       int            `json:"sort"`                                //排序
	Type       int            `json:"type"`                                //菜单类型
	Data       string         `json:"data"`                                //数据
	DataJson   *AdminMenuData `json:"dataJson" gorm:"-"`                   //菜单详情
	Status     int            `json:"status"`                              //状态
	CreatedAt  int            `json:"createdAt"`                           //创建时间
}

// AdminSettingMenuIndexParams 菜单列表参数
type AdminSettingMenuIndexParams struct {
	Name       string                 `json:"name"`       //菜单名称
	AdminId    int                    `json:"adminId"`    //	管理ID
	AdminName  string                 `json:"adminName"`  //管理名称
	ParentName string                 `json:"parentName"` //上级菜单名称
	Route      string                 `json:"route"`      //路由
	Status     int                    `json:"status"`     //状态
	Type       int                    `json:"type"`       //类型
	Pagination *utils.Pagination      `json:"pagination"` //分页参数
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //过期时间
}

// AdminSettingMenuCreateParams 新增菜单参数
type AdminSettingMenuCreateParams struct {
	Id    int    `json:"="`
	Name  string `json:"name" validate:"required"` //菜单名称
	Route string `json:"route"`                    //路由
	Type  int    `json:"type"`                     //菜单类型
}

// AdminSettingMenuUpdateParams 更新菜单参数
type AdminSettingMenuUpdateParams struct {
	Id       int    `json:"id" validate:"required" gorm:"-"`         //ID
	ParentId int    `json:"parentId"`                                //父级ID
	Name     string `json:"name"`                                    //名称
	Sort     int    `json:"sort"`                                    //排序
	Route    string `json:"route"`                                   //路由
	Type     int    `json:"type"`                                    //类型
	Status   int    `json:"status" validate:"omitempty,oneof=-1 10"` //状态
}

// AdminSettingMenuUpdateDescParams 更新菜单详情
type AdminSettingMenuUpdateDescParams struct {
	Id       int            `json:"id" validate:"required" gorm:"-"` //ID
	DataJson *AdminMenuData `json:"dataJson"`
}
