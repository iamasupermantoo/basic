package dtowallets

import (
	"basic/utils"
)

// BillInfo 用户账单信息
type BillInfo struct {
	Id        int     `json:"id"`
	AdminId   int     `json:"adminId"`   // 管理ID
	UserId    int     `json:"userId"`    // 用户ID
	AssetsId  int     `json:"assetsId"`  // 资产ID
	SourceId  int     `json:"sourceId"`  // 来源ID
	Type      int     `json:"type"`      // 类型
	Name      string  `json:"name"`      // 名称
	Money     float64 `json:"money"`     // 金额
	Balance   float64 `json:"balance"`   // 余额
	Data      string  `json:"data"`      // 数据
	CreatedAt int     `json:"createdAt"` // 创建时间
}

// BillIndexParams 用户账单列表查询数据
type BillIndexParams struct {
	AdminName  string                 `json:"adminName"`  // 管理员名
	UserName   string                 `json:"username"`   // 用户名
	AssetsId   int                    `json:"assetsId"`   // 资产ID
	Type       int                    `json:"type"`       // 类型
	DateTime   *utils.RangeDatePicker `json:"createdAt"`  // 创建时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页数据
}

// BillIndexData 用户账单列表返回数据
type BillIndexData struct {
	BillInfo
	AdminName string `json:"adminName" gorm:"column:adminName"` //管理员名称
	UserName  string `json:"username" gorm:"column:username"`   //用户名
}
