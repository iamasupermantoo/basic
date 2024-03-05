package dtosettings

import "basic/utils"

// ArticleInfo 文章信息
type ArticleInfo struct {
	Id        int    `json:"id"`        // 主键
	AdminId   int    `json:"adminId"`   // 管理ID
	Type      int    `json:"type"`      // 类型 1帮助中心类型 2底部类型
	Image     string `json:"image"`     // 图片
	Name      string `json:"name"`      // 标题
	Content   string `json:"content"`   // 内容
	Status    int    `json:"status"`    // 状态 -1禁用 10开启
	Data      string `json:"data"`      // 数据
	UpdatedAt int    `json:"updatedAt"` // 更新时间
	CreatedAt int    `json:"createdAt"` // 创建时间
}

// ArticleCreateParams 文章创建
type ArticleCreateParams struct {
	Type    int    `json:"type" validate:"oneof=1 2 10"` // 类型 1帮助中心类型 2底部类型
	Name    string `json:"name"  validate:"required"`    // 标题
	Content string `json:"content"  validate:"required"` // 内容
}

// ArticleUpdateParams 文章修改
type ArticleUpdateParams struct {
	Id      int    `json:"id" gorm:"-" validate:"required"` // 主键
	Type    int    `json:"type"`                            // 类型 1帮助中心类型 2底部类型
	Image   string `json:"image"`                           // 图片
	Name    string `json:"name"`                            // 标题
	Content string `json:"content"`                         // 内容
	Status  int    `json:"status"`                          // 状态  -1禁用 10开启
	Data    string `json:"data"`                            // 数据
}

// ArticleIndexParams 文章列表
type ArticleIndexParams struct {
	AdminId    int                    `json:"adminId"`    //	管理Id
	AdminName  string                 `json:"adminName"`  // 管理名称
	Type       int                    `json:"type"`       // 类型 1帮助中心类型 2底部类型
	Name       string                 `json:"name"`       // 标题
	Content    string                 `json:"content"`    // 内容
	Status     int                    `json:"status"`     // 状态  -1禁用 10开启
	Data       string                 `json:"data"`       // 数据
	UpdatedAt  *utils.RangeDatePicker `json:"updatedAt"`  // 更新时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页
}

// ArticleIndexData 文章查询返回数据
type ArticleIndexData struct {
	ArticleInfo        // 文章详情
	AdminName   string `json:"adminName" gorm:"column:adminName"` // 管理名称
}
