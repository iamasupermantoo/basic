package views

const (
	SizingSmall    = "small"     //	dialog 小型宽度
	SizingMedium   = "medium"    //	dialog	中型宽度 默认值
	ButtonSizeXS   = "xs"        //	小按钮
	ButtonSizeSM   = "sm"        //	一般
	ButtonSizeMD   = "md"        //	正常
	ColorPrimary   = "primary"   //	主颜色
	ColorSecondary = "secondary" //	辅色
	ColorPositive  = "positive"  //	绿色
	ColorNegative  = "negative"  //	红色
	ColorAccent    = "accent"    //	紫色
	ColorWarning   = "warning"   //	警告色

	OperateCheckbox = "checkbox" //	选择操作
	OperateSetting  = "setting"  //	配置操作
)

// InputOptions select|checkbox|toggle|radio
type InputOptions struct {
	Label string      `json:"label"` //	名称
	Value interface{} `json:"value"` //	值
}

// InputJsonOptions Json格式
type InputJsonOptions struct {
	Items [][]*InputAttrsViews `json:"items"` //	input组
}

// InputChildrenOptions 子级
type InputChildrenOptions struct {
	Items    [][]*InputAttrsViews `json:"items"`    //	input组
	IsCreate bool                 `json:"isCreate"` //	是否新增
	IsDelete bool                 `json:"isDelete"` //	是否删除
}

// CheckboxListOptions Table多选配置参数
type CheckboxListOptions struct {
	Operate string `json:"operate"` //	操作名称
	Name    string `json:"name"`    //	名称
	Field   string `json:"field"`   //	提交的字段名称
	Scan    string `json:"scan"`    //	提取的字段名称
}

// SettingOptions
type SettingOptions struct {
	Operate string `json:"operate"` //	操作名称
	Name    string `json:"name"`    //	名称
	Params  string `json:"params"`  //	参数
	Type    string `json:"type"`    //  类型字段
	Value   string `json:"value"`   // 	值字段
	Input   string `json:"input"`   //	输入框
}
