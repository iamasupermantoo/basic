package dtosettings

import "basic/utils"

// SettingInfo 设置信息
type SettingInfo struct {
	Id        int    `json:"id"`        // 主键
	AdminId   int    `json:"adminId"`   // 管理ID
	GroupId   int    `json:"groupId"`   // 分组ID
	Name      string `json:"name"`      // 设置名称
	Type      int    `json:"type"`      // input类型 1文本类型 2文本域类型 3富文本类型 4数字类型 5密码类型 10选择类型 11单选类型 12多选类型 13开关类型 21文件类型 22图片类型 23图片组类型 31时间类型 32时间范围类型 41对象类型 42子集对象类型 52编辑文本类型 53编辑富文本类型 54编辑开关类型 61翻译类型
	Field     string `json:"field"`     // 键名
	Value     string `json:"value"`     // 键值
	Data      string `json:"data"`      // 配置
	UpdatedAt int    `json:"updatedAt"` // 更新时间
	CreatedAt int    `json:"createdAt"` // 创建时间
}

// SettingCreateParams 设置创建
type SettingCreateParams struct {
	GroupId int    `json:"groupId" validate:"required"` // 分组ID
	Type    int    `json:"type" validate:"required"`    // 类型
	Name    string `json:"name" validate:"required"`    // 设置名称
	Field   string `json:"field" validate:"required"`   // 键名
	Value   string `json:"value"`                       // 键值
	Data    string `json:"data"`                        // 配置
}

// SettingUpdateParams 设置修改
type SettingUpdateParams struct {
	Id        int    `json:"id" validate:"required" gorm:"-"`                            // 主键
	GroupId   int    `json:"groupId"`                                                    // 分组ID
	Name      string `json:"name"`                                                       // 设置名称
	Type      int    `json:"type"`                                                       // input类
	Field     string `json:"field"`                                                      // 键名
	Value     string `json:"value"`                                                      // 键值
	Data      string `json:"data"`                                                       // 配置
	UpdatedAt int    `json:"-" gorm:"column:updated_at;type:datetime(0);autoUpdateTime"` // 更新时间
}

// SettingUpdateDescParams 管理设置更新详情
type SettingUpdateDescParams struct {
	Id    int         `json:"id" validate:"required" gorm:"-"` // 主键
	Type  int         `json:"type" validate:"required"`        //	类型
	Value interface{} `json:"value"`                           // 值
}

// SettingIndexParams 设置列表
type SettingIndexParams struct {
	AdminName  string                 `json:"adminName"`  // 管理名称
	Name       string                 `json:"name"`       // 设置名称
	GroupId    int                    `json:"groupId"`    // 分组
	Type       int                    `json:"type"`       // 类型
	Field      string                 `json:"field"`      // 键名
	Value      string                 `json:"value"`      // 键值
	Data       string                 `json:"data"`       // 配置
	UpdatedAt  *utils.RangeDatePicker `json:"updatedAt"`  // 更新时间
	Pagination *utils.Pagination      `json:"pagination"` // 分页
}

// SettingIndexData 设置查询返回数据
type SettingIndexData struct {
	SettingInfo
	AdminName string `json:"adminName" gorm:"column:adminName"` // 管理名称
}
