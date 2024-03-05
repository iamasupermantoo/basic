package views

import "basic/models"

type InputAttrsViews struct {
	Label    string      `json:"label"`    //	显示标题
	Field    string      `json:"field"`    //	字段名称
	Type     int         `json:"type"`     //	input类型
	Readonly bool        `json:"readonly"` //	只读
	Data     interface{} `json:"data"`     //	input配置
}

type InputViews struct {
	InputList []*InputAttrsViews     //	input组
	Params    map[string]interface{} //	input参数
}

// NewInputViews 创建Inputs
func NewInputViews() *InputViews {
	return &InputViews{
		Params:    map[string]interface{}{},
		InputList: make([]*InputAttrsViews, 0),
	}
}

// GetInputListInfo 获取参数信息
func (_InputViews *InputViews) GetInputListInfo() (interface{}, []*InputAttrsViews) {
	return _InputViews.Params, _InputViews.InputList
}

// Text 文本
func (_InputViews *InputViews) Text(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeText,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// TextArea 文本域
func (_InputViews *InputViews) TextArea(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeTextArea,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Editor 富文本
func (_InputViews *InputViews) Editor(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeEditor,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Number 数字
func (_InputViews *InputViews) Number(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeNumber,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Password 密码
func (_InputViews *InputViews) Password(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypePassword,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Select 选择框
func (_InputViews *InputViews) Select(label, field string, data []*InputOptions) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeSelect,
		Field: field,
		Label: label,
		Data:  data,
	})
	_InputViews.Params[field] = 0
	return _InputViews
}

// Radio 单选框
func (_InputViews *InputViews) Radio(label, field string, data []*InputOptions) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeRadio,
		Field: field,
		Label: label,
		Data:  data,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Checkbox 多选框
func (_InputViews *InputViews) Checkbox(label, field string, data []*InputOptions) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeCheckbox,
		Field: field,
		Label: label,
		Data:  data,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Toggle 开关
func (_InputViews *InputViews) Toggle(label, field string, data []*InputOptions) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeToggle,
		Field: field,
		Label: label,
		Data:  data,
	})
	_InputViews.Params[field] = false
	return _InputViews
}

// File 文件
func (_InputViews *InputViews) File(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeFile,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Image 图片
func (_InputViews *InputViews) Image(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeImage,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Images 图片组
func (_InputViews *InputViews) Images(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeImages,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// DatePicker 时间格式
func (_InputViews *InputViews) DatePicker(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeDatePicker,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// RangeDatePicker 时间范围
func (_InputViews *InputViews) RangeDatePicker(label, field string) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeRangeDatePicker,
		Field: field,
		Label: label,
	})
	_InputViews.Params[field] = nil
	return _InputViews
}

// Json 结构体
func (_InputViews *InputViews) Json(label, field string, value interface{}) *InputViews {

	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeJson,
		Field: field,
		Label: label,
		Data:  _InputViews.InputAttrsViewsToTwo(value.([]*InputAttrsViews)),
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// Children 子级
func (_InputViews *InputViews) Children(label, field string, value interface{}) *InputViews {
	_InputViews.InputList = append(_InputViews.InputList, &InputAttrsViews{
		Type:  models.AdminSettingTypeChildren,
		Field: field,
		Label: label,
		Data:  _InputViews.InputAttrsViewsToTwo(value.([]*InputAttrsViews)),
	})
	_InputViews.Params[field] = ""
	return _InputViews
}

// SetValue 设置值
func (_InputViews *InputViews) SetValue(field string, value interface{}) *InputViews {
	if _, ok := _InputViews.Params[field]; ok {
		_InputViews.Params[field] = value
	}
	return _InputViews
}

// SetReadonly 设置是否显示
func (_InputViews *InputViews) SetReadonly(field string, readonly bool) *InputViews {
	for _, v := range _InputViews.InputList {
		if v.Field == field {
			v.Readonly = readonly
		}
	}
	return _InputViews
}

// InputAttrsViewsToTwo 单微数组转二维
func (_InputViews *InputViews) InputAttrsViewsToTwo(inputList []*InputAttrsViews) [][]*InputAttrsViews {
	data := make([][]*InputAttrsViews, 0)

	for i := 0; i < len(inputList); i++ {
		data = append(data, []*InputAttrsViews{inputList[i]})
	}
	return data
}
