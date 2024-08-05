package dtoadmins

import "basic/utils"

// RoleInfo 角色信息
type RoleInfo struct {
	Name      string `json:"name"`      //名称
	Type      int    `json:"type"`      //类型
	Desc      string `json:"desc"`      //详情
	Rule      string `json:"rule"`      //规则
	Data      string `json:"data"`      //数据
	UpdatedAt int    `json:"updatedAt"` //更新时间
	CreatedAt int    `json:"createdAt"` //创建时间
}

// RoleIndexParams 角色列表参数
type RoleIndexParams struct {
	Name       string                 ` json:"name"`      //角色名称
	Desc       string                 ` json:"desc"`      //角色详情
	CreatedAt  *utils.RangeDatePicker ` json:"createdAt"` //操作时间
	Pagination *utils.Pagination      `json:"pagination"` //分页信息
}

// RoleIndexData 角色列表数据
type RoleIndexData struct {
	RoleInfo
	CurrentyAuth string          `json:"currentAuth"` //	当前权限
	Authority    map[string]bool `json:"authority"`   //角色权限选中对列表
}

// RoleDeleteParams 角色删除参数
type RoleDeleteParams struct {
	NameList []string `json:"nameList" validate:"required"` //删除名称
}

// RoleCreateParams 角色新增参数
type RoleCreateParams struct {
	Name      string          `json:"name" validate:"required"` //名称
	Authority map[string]bool `json:"authority"`                //选中权限列表
	Desc      string          `json:"desc"`                     //描述
}

// RoleUpdateParams 更新角色参数
type RoleUpdateParams struct {
	Name      string          `json:"name" validate:"required"` //名称
	Desc      string          `json:"desc"`                     //描述
	Authority map[string]bool `json:"authority"`                //选中权限列表
}
