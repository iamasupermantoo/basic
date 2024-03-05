package auth

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/user/auth/index", "/user/auth/update")

	//	查询个别认证
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("用户名称", "username").
		Text("真实姓名", "realName").
		Text("详细地址", "address").
		Select("类型", "type", []*views.InputOptions{{Label: "身份证", Value: models.UserRealNameTypeIdCard}}).
		SetValue("type", models.UserRealNameTypeIdCard).
		Select("状态", "status", []*views.InputOptions{{Label: "拒绝", Value: models.UserRealNameStatusRefuse}, {Label: "完成", Value: models.UserRealNameStatusComplete}, {Label: "审核", Value: models.UserRealNameStatusActive}}).
		RangeDatePicker("申请时间", "createdAt").
		GetInputListInfo()

	//	Table 工具栏目，用于新增和删除
	createParams, createInputList := views.NewInputViews().
		Select("类型", "type", []*views.InputOptions{{Label: "身份证", Value: models.UserRealNameTypeIdCard}}).
		Text("用户名称", "username").
		Text("真实姓名", "realName").
		Text("卡号", "number").
		Image("证件正面", "photo1").
		Image("证件反面", "photo2").
		SetValue("type", models.UserRealNameTypeIdCard).
		GetInputListInfo()
	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/user/auth/delete", "批量删除认证").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "realName", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增认证", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/user/auth/create", "新增认证").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	查询所有认证字段
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("用户名称", "username", false).
		EditText("真实姓名", "realName", true).
		EditNumber("卡号｜地址", "number", true).
		Image("证件正面", "photo1", false).
		Image("证件反面", "photo2", false).
		Image("手持照片", "photo3", false).
		EditText("手机号码", "telephone", true).
		EditText("详细地址", "address", true).
		Select("类型", "type", true, []*views.InputOptions{{Label: "身份证", Value: models.UserRealNameTypeIdCard}}).
		Select("状态", "status", true, []*views.InputOptions{{Label: "拒绝", Value: models.UserRealNameStatusRefuse}, {Label: "完成", Value: models.UserRealNameStatusComplete}, {Label: "审核", Value: models.UserRealNameStatusActive}}).
		EditText("数据", "data", true).
		DatePicker("创建时间", "createdAt", true).
		DatePicker("更新时间", "updatedAt", true).
		GetColumnsListInfo()

	//	更新认证
	updateParams, updateInputList := views.NewInputViews().
		Image("证件正面", "photo1").
		Image("证件反面", "photo2").
		Image("手持照片", "photo3").
		Text("真实姓名", "realName").
		Text("卡号｜地址", "number").
		Text("手机号码", "telephone").
		Text("详情地址", "address").
		Select("类型", "type", []*views.InputOptions{{Label: "身份证", Value: models.UserRealNameTypeIdCard}}).
		Select("状态", "status", []*views.InputOptions{{Label: "拒绝", Value: models.UserRealNameStatusRefuse}, {Label: "审核", Value: models.UserRealNameStatusActive}, {Label: "完成", Value: models.UserRealNameStatusComplete}}).
		TextArea("数据", "data").
		GetInputListInfo()

	pendingParams, pendingInputList := views.NewInputViews().
		Select("状态", "status", []*views.InputOptions{{Label: "拒绝", Value: models.UserRealNameStatusRefuse}, {Label: "完成", Value: models.UserRealNameStatusComplete}}).
		TextArea("拒绝理由", "data").
		SetValue("status", models.UserRealNameStatusComplete).
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/user/auth/update", "更新认证").
				SetParams(updateParams).SetInputList(updateInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "审核", Color: views.ColorAccent, Size: "xs", Eval: "scope.row.status == 10"},
			Config: views.NewDialogViews("update", "/user/auth/update", "审核用户认证状态").
				SetParams(pendingParams).SetInputList(pendingInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
