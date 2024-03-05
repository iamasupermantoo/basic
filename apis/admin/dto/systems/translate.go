package dtosystems

import "basic/utils"

// TranslateCreateParams 新增字典参数
type TranslateCreateParams struct {
	Id      int    `json:"-"`                         //	主键ID
	AdminId int    `json:"-"`                         //	管理Id
	Name    string `json:"name" validate:"required"`  //	名称
	Lang    string `json:"-" gorm:"default:'zh-CN'"`  //	语言标识
	Type    int    `json:"type" validate:"required"`  //	类型
	Field   string `json:"field" validate:"required"` //	建铭
	Value   string `json:"value" validate:"required"` //	键值
}

// TranslateUpdateParams 更新字典参数
type TranslateUpdateParams struct {
	Id    int    `json:"id" validate:"required" gorm:"-"` //	Id
	Name  string `json:"name"`                            //	名称
	Type  int    `json:"type"`                            //	类型
	Field string `json:"field"`                           //	建铭
	Value string `json:"value"`                           //	键值
}

// TranslateIndexParams 字典列表参数
type TranslateIndexParams struct {
	AdminName  string                 `json:"adminName"`  //	管理Id
	Name       string                 `json:"name"`       //	名称
	Type       int                    `json:"type"`       //	类型
	CreatedAt  *utils.RangeDatePicker `json:"createdAt"`  //	创建时间范围
	Pagination *utils.Pagination      `json:"pagination"` //	分页
}

// TranslateIndexData 字典列表数据
type TranslateIndexData struct {
	Id        int    `json:"id"`                                //	Id
	AdminName string `json:"adminName" gorm:"column:adminName"` //	管理员名
	Lang      string `json:"lang"`                              //	语言标识
	AdminId   int    `json:"adminId"`                           //	管理Id
	Name      string `json:"name"`                              //	名称
	Type      int    `json:"type"`                              //	类型
	Field     string `json:"field"`                             //	键名
	Value     string `json:"value"`                             //	键值
	CreatedAt int    `json:"createdAt"`                         //	创造时间
}

// TranslateFiedParams 翻译语言字段参数
type TranslateFiedParams struct {
	AdminId int    `json:"adminId"` //	管理ID
	Field   string `json:"field"`   //	字段名称
}

// TranslateGoogleParams google 翻译参数
type TranslateGoogleParams struct {
	AdminId int    `json:"adminId"` //	管理ID
	Field   string `json:"field"`   //	字段名称
	Alias   string `json:"alias"`   //	翻译语言
}

// TranslateUpdateField 翻译更新字段
type TranslateUpdateField struct {
	AdminId int                   `json:"adminId"` //	管理ID
	Field   string                `json:"field"`   //	字段名称
	Data    []*TranslateFieldData `json:"data"`    //	数据
}

// TranslateFieldData 翻译语言字段数据
type TranslateFieldData struct {
	Id          int    `json:"id"`                                    //	语言ID
	Name        string `json:"name"`                                  //	语言名称
	Alias       string `json:"alias"`                                 //	语言别名
	Icon        string `json:"icon"`                                  //	语言图标
	TranslateId int    `json:"translateId" gorm:"column:translateId"` //	翻译ID
	Value       string `json:"value"`                                 //	翻译值
}
