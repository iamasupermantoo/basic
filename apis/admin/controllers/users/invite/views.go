package invite

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/user/invite/index", "/user/invite/update")

	//	查询个别邀请
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名称", "username").
		Text("邀请码", "code").
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.UserInviteStatusActive}, {Label: "禁用", Value: models.UserInviteStatusDisable}}).
		RangeDatePicker("创建时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Text("用户名称｜管理名称", "username").
		Select("类型", "type", []*views.InputOptions{{Label: "管理邀请码", Value: models.UserInviteTypeAdmin}, {Label: "用户邀请码", Value: models.UserInviteTypeUser}}).
		SetValue("type", models.UserInviteTypeUser).
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/user/invite/delete", "批量删除邀请").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "code", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增邀请码", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/user/invite/create", "新增邀请").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有邀请字段
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名称", "username", false).
		EditText("邀请码", "code", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "开启", Value: models.UserInviteStatusActive}, {Label: "禁用", Value: models.UserInviteStatusDisable}}).
		EditText("数据", "data", true).
		DatePicker("创建时间", "createdAt", true).
		DatePicker("更新时间", "updatedAt", true).
		GetColumnsListInfo()

	//	更新邀请
	updateParams, updateInputList := views.NewInputViews().
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.UserInviteStatusActive}, {Label: "禁用", Value: models.UserInviteStatusDisable}}).
		Text("邀请码", "code").
		TextArea("数据", "data").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/user/invite/update", "更新消息").SetSizingSmall().
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
