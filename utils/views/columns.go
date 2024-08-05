package views

import "basic/models"

// columnsAttrsViews 表格数据显示
type ColumnsAttrsViews struct {
	Label    string      `json:"label"`    //	显示标题
	Field    string      `json:"field"`    //	字段名称
	Type     int         `json:"type"`     //	input类型
	Name     string      `json:"name"`     //	字段名称
	Align    string      `json:"align"`    //	字符对齐
	Sortable bool        `json:"sortable"` //	是否排序
	Data     interface{} `json:"data"`     //	input配置
}

type ColumnsViews struct {
	ColumnsList []*ColumnsAttrsViews
}

func NewColumnsViews() *ColumnsViews {
	return &ColumnsViews{
		ColumnsList: make([]*ColumnsAttrsViews, 0),
	}
}

// GetColumnsListInfo 获取数据字段列表
func (_ColumnsViews *ColumnsViews) GetColumnsListInfo() []*ColumnsAttrsViews {
	return _ColumnsViews.ColumnsList
}

func (_ColumnsViews *ColumnsViews) Text(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeText,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) Translate(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeTranslate,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) Select(label, field string, sortable bool, data []*InputOptions) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeSelect,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
		Data:     data,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) Image(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeImage,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) Images(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeImages,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) DatePicker(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeDatePicker,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) EditText(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeEditText,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) EditNumber(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeEditNumber,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) EditTextArea(label, field string, sortable bool) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeEditTextArea,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
	})
	return _ColumnsViews
}

func (_ColumnsViews *ColumnsViews) EditToggle(label, field string, sortable bool, data []*InputOptions) *ColumnsViews {
	_ColumnsViews.ColumnsList = append(_ColumnsViews.ColumnsList, &ColumnsAttrsViews{
		Type:     models.AdminSettingTypeEditToggle,
		Label:    label,
		Field:    field,
		Name:     field,
		Align:    "left",
		Sortable: sortable,
		Data:     data,
	})
	return _ColumnsViews
}
