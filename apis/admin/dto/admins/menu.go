package dtoadmins

import (
	"basic/utils"
)

// AdminMenuData 菜单数据
type AdminMenuData struct {
	Icon       string `json:"icon"`       //	图标
	ActiveIcon string `json:"activeIcon"` //	已激活图标
	VueTmp     string `json:"vueTmp"`     //	vue模版
	ViewsUrl   string `json:"viewsUrl"`   //	视图地址
	IsPc       bool   `json:"isPc"`       //	Pc是否显示
	IsMobile   bool   `json:"isMobile"`   //	手机是否显示
}

// MenuInfo 菜单信息
type MenuInfo struct {
	Id       int            `json:"id"`       //	ID
	Name     string         `json:"name"`     //	路由名称
	Route    string         `json:"route"`    //	路由
	Children []*MenuInfo    `json:"children"` //	子路由
	Data     *AdminMenuData `json:"data"`     //	数据
}

// AdminMenuIndexParams 菜单列表参数
type AdminMenuIndexParams struct {
	Name       string                 `json:"name"`       //菜单名称
	AdminName  string                 `json:"adminName"`  //管理名称
	ParentName string                 `json:"parentName"` //上级菜单名称
	Route      string                 `json:"route"`      //路由
	Status     int                    `json:"status"`     //状态
	Pagination *utils.Pagination      `json:"pagination"` //分页参数
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //过期时间
}

// AdminMenuIndexData 菜单列表数据
type AdminMenuIndexData struct {
	Id         int            `json:"id"`                                  //ID
	ParentId   int            `json:"parentId"`                            //上级ID
	AdminName  string         `json:"adminName" gorm:"column:adminName"`   //管理
	ParentName string         `json:"parentName" gorm:"column:parentName"` //父级菜单
	Name       string         `json:"name"`                                //菜单名称
	Route      string         `json:"route"`                               //路由
	Sort       int            `json:"sort"`                                //排序
	Data       string         `json:"data"`                                //数据
	DataJson   *AdminMenuData `json:"dataJson" gorm:"-"`                   //菜单详情
	Status     int            `json:"status"`                              //状态
	CreatedAt  int            `json:"createdAt"`                           //创建时间
}

// AdminMenuUpdateParams 更新菜单参数
type AdminMenuUpdateParams struct {
	Id       int    `json:"id" validate:"required" gorm:"-"`         //ID
	ParentId int    `json:"parentId"`                                //父级ID
	Name     string `json:"name"`                                    //名称
	Sort     int    `json:"sort"`                                    //排序
	Route    string `json:"route"`                                   //路由
	Data     string `json:"data"`                                    //数据
	Status   int    `json:"status" validate:"omitempty,oneof=-1 10"` //状态
}

// AdminMenuUpdateDescParams 更新菜单详情
type AdminMenuUpdateDescParams struct {
	Id       int            `json:"id" validate:"required" gorm:"-"` //ID
	DataJson *AdminMenuData `json:"dataJson"`
}

// AdminMenuCreateParams 新增菜单参数
type AdminMenuCreateParams struct {
	Name  string `json:"name"`  //菜单名称
	Route string `json:"route"` //路由
}
