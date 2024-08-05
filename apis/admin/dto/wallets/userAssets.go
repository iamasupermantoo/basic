package dtowallets

import "basic/utils"

// UserAssetsInfo 用户资产详细信息
type UserAssetsInfo struct {
	Id             int     `json:"id"`             // 主键
	AdminUserId    int     `json:"adminUserId"`    // 管理ID
	UserId         int     `json:"userId"`         // 用户ID
	WalletAssetsId int     `json:"walletAssetsId"` // 钱包资产ID
	Status         int     `json:"status"`         // 状态
	Money          float64 `json:"money"`          // 金额
	Data           string  `json:"data"`           // 数据
	UpdatedAt      int     `json:"updatedAt"`      // 更新时间
	CreatedAt      int     `json:"createdAt"`      // 创建时间
}

// UserAssetsUpdateParams 用户资产更新
type UserAssetsUpdateParams struct {
	Id     int `json:"id" validate:"required"`                  // 资产ID
	Status int `json:"status" validate:"omitempty,oneof=-1 10"` // 状态
}

// UserAssetsIndexParams 用户资产查询
type UserAssetsIndexParams struct {
	AdminName      string                 `json:"adminName"`      // 管理员名
	UserName       string                 `json:"UserName"`       // 用户名
	WalletAssetsId int                    `json:"walletAssetsId"` // 资产类型
	Status         int                    `json:"status"`         // 状态
	CreatedAt      *utils.RangeDatePicker `json:"createdAt"`      // 激活时间
	Pagination     *utils.Pagination      `json:"pagination"`     // 分页
}

// UserAssetsIndexData 用户资产查询返回数据
type UserAssetsIndexData struct {
	UserAssetsInfo          // 资产详细信息
	AdminName        string `json:"adminName" gorm:"column:adminName"`               // 管理员名
	UserName         string `json:"username" gorm:"column:userName"`                 // 用户名
	WalletAssetsName string `json:"walletAssetsName" gorm:"column:walletAssetsName"` // 钱包资产名
}

// UserAssetsBalanceParams 用户资产余额变动
type UserAssetsBalanceParams struct {
	UserName      string  `json:"username"`                               //	用户名
	Type          int     `json:"type" validate:"required,oneof=103 113"` //	类型
	WalletAssetId int     `json:"walletAssetsId" validate:"gte=0"`        //	钱包资产ID
	Money         float64 `json:"money" validate:"required,gt=0"`         //	金额
}
