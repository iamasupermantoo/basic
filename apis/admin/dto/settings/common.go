package dtosettings

import (
	"basic/models"
	"basic/utils/views"
)

// AdminSettingGroupOptions 管理设置分组Options
var AdminSettingGroupOptions = []*views.InputOptions{
	{Label: "基础分组", Value: models.AdminSettingGroupBasic},
	{Label: "钱包分组", Value: models.AdminSettingGroupWallet},
	{Label: "收益分组", Value: models.AdminSettingGroupEarnings},
	{Label: "模版分组", Value: models.AdminSettingGroupTemplate},
}

// AdminSettingTypeOptions 管理配置类型Options
var AdminSettingTypeOptions = []*views.InputOptions{
	{Label: "文本类型", Value: models.AdminSettingTypeText},
	{Label: "文本域类型", Value: models.AdminSettingTypeTextArea},
	{Label: "富文本类型", Value: models.AdminSettingTypeEditor},
	{Label: "数字类型", Value: models.AdminSettingTypeNumber},
	{Label: "密码类型", Value: models.AdminSettingTypePassword},
	{Label: "选择类型", Value: models.AdminSettingTypeSelect},
	{Label: "单选类型", Value: models.AdminSettingTypeRadio},
	{Label: "多选类型", Value: models.AdminSettingTypeCheckbox},
	{Label: "开关类型", Value: models.AdminSettingTypeToggle},
	{Label: "文件类型", Value: models.AdminSettingTypeFile},
	{Label: "图片类型", Value: models.AdminSettingTypeImage},
	{Label: "图片组类型", Value: models.AdminSettingTypeImages},
	{Label: "时间类型", Value: models.AdminSettingTypeDatePicker},
	{Label: "时间范围类型", Value: models.AdminSettingTypeRangeDatePicker},
	{Label: "Json类型", Value: models.AdminSettingTypeJson},
	{Label: "混合类型", Value: models.AdminSettingTypeChildren},
	{Label: "翻译类型", Value: models.AdminSettingTypeTranslate},
}
