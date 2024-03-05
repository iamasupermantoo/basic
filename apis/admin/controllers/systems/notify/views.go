package notify

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/system/notify/index", "/system/notify/update")

	//	查询个别通知
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名称", "username").
		Text("通知标题", "name").
		Select("通知状态", "status", []*views.InputOptions{{Label: "未读消息", Value: models.SystemNotifyStatusUnread}, {Label: "已读消息", Value: models.SystemNotifyStatusRead}}).
		SetValue("type", 0).
		RangeDatePicker("发送时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Text("用户名称", "username").
		Text("通知标题", "name").
		Editor("通知内容", "content").
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/system/notify/delete", "批量删除通知").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增通知", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/system/notify/create", "新增通知").SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有通知字段
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名称", "username", false).
		EditText("通知标题", "name", true).
		Select("通知状态", "status", true, []*views.InputOptions{{Label: "已读消息", Value: models.SystemNotifyStatusRead}, {Label: "未读消息", Value: models.SystemNotifyStatusUnread}}).
		EditTextArea("通知内容", "content", true).
		EditTextArea("数据", "data", true).
		DatePicker("发送时间", "createdAt", true).
		DatePicker("读取时间", "updatedAt", true).
		GetColumnsListInfo()

	//	更新通知
	updateParams, updateInputList := views.NewInputViews().
		Text("通知标题", "name").
		Select("通知状态", "status", []*views.InputOptions{{Label: "未读消息", Value: models.SystemNotifyStatusUnread}, {Label: "已读消息", Value: models.SystemNotifyStatusRead}}).
		Editor("通知内容", "content").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/system/notify/update", "更新消息").
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
