package level

import (
	"basic/apis/common/service/users"
	"basic/models"
	"basic/utils"
	"basic/utils/views"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)
	config := views.NewTableViews("/user/level/index", "/user/level/update")

	//	查询管理会员设置列表
	levelList := users.GetLevelOptions(adminId)
	if len(levelList) == 0 {
		return c.JSON(utils.ErrorJson(errors.New("没有会员设置")))
	}

	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名称", "username").
		Select("等级", "levelId", levelList).
		Select("状态", "status", []*views.InputOptions{{Label: "启用", Value: models.UserLevelOrderStatusActive}, {Label: "禁用", Value: models.UserLevelOrderStatusDisable}}).
		RangeDatePicker("购买时间", "createdAt").
		RangeDatePicker("到期时间", "expiredAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Text("用户名称", "username").
		Select("会员等级", "levelId", levelList).
		SetValue("levelId", levelList[0].Value).
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/user/level/delete", "批量删除会员").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "username", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增会员", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/user/level/create", "新增会员").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有会员字段
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名称", "username", false).
		Select("等级", "levelId", true, levelList).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "开启", Value: models.UserLevelOrderStatusActive}, {Label: "禁用", Value: models.UserLevelOrderStatusDisable}}).
		Text("数据", "data", true).
		DatePicker("购买时间", "createdAt", true).
		DatePicker("到期时间", "expiredAt", true).
		GetColumnsListInfo()

	//	更新会员
	updateParams, updateInputList := views.NewInputViews().
		Select("状态", "status", []*views.InputOptions{{Label: "开启", Value: models.UserLevelOrderStatusActive}, {Label: "禁用", Value: models.UserLevelOrderStatusDisable}}).
		TextArea("数据", "data").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/user/level/update", "更新会员").
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
