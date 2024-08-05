package user

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

// Views 模版JSON信息
func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/user/index", "/user/update")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").Text("上级用户", "parentName").Text("用户名", "username").
		Text("昵称", "nickname").Text("邮箱", "email").Text("号码", "telephone").
		Select("性别", "sex", []*views.InputOptions{{Label: "男", Value: models.UserSexMale}, {Label: "女", Value: models.UserSexWoman}}).
		RangeDatePicker("生日", "birthday").
		Select("类型", "type", []*views.InputOptions{{Label: "虚拟用户", Value: models.UserTypeVirtual}, {Label: "普通用户", Value: models.UserTypeDefault}, {Label: "会员用户", Value: models.UserTypeLevel}}).
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.UserStatusDisable}, {Label: "激活", Value: models.UserStatusActive}}).
		RangeDatePicker("注册时间", "createdAt").GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Text("用户名", "username").Password("密码", "password").GetInputListInfo()

	balanceParams, balanceInputList := views.NewInputViews().
		Text("用户名", "username").
		Select("类型", "type", []*views.InputOptions{{Label: "加款", Value: models.WalletBillTypeSystemDeposit}, {Label: "减款", Value: models.WalletBillTypeSystemWithdraw}}).
		Number("金额", "money").
		SetValue("type", models.WalletBillTypeSystemDeposit).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/user/delete", "批量删除用户组").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "username", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增用户", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/user/create", "新增平台用户").SetSizingSmall().SetParams(createParams).SetInputList(createInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "资金加减款", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/user/balance", "用户资金加减款").SetSizingSmall().SetParams(balanceParams).SetInputList(balanceInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).Text("管理名称", "adminName", false).Image("头像", "avatar", false).Text("上级用户", "parentName", false).
		Text("用户名", "username", true).EditNumber("信用分", "score", true).EditText("昵称", "nickname", true).EditText("邮箱", "email", true).EditText("手机号码", "telephone", true).
		Text("金额", "money", true).
		Select("类型", "type", true, []*views.InputOptions{{Label: "虚拟用户", Value: models.UserTypeVirtual}, {Label: "普通用户", Value: models.UserTypeDefault}, {Label: "会员用户", Value: models.UserTypeLevel}}).
		Select("状态", "status", true, []*views.InputOptions{{Label: "激活", Value: models.UserStatusActive}, {Label: "禁用", Value: models.UserStatusDisable}}).
		DatePicker("最近时间", "updatedAt", true).DatePicker("注册时间", "createdAt", true).GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Image("头像", "avatar").
		Text("昵称", "nickname").
		Text("邮箱", "email").
		Text("电话号码", "telephone").
		Select("性别", "sex", []*views.InputOptions{{Label: "男", Value: models.UserSexMale}, {Label: "女", Value: models.UserSexWoman}}).
		DatePicker("生日", "birthday").
		Password("密码", "password").
		Password("安全密码", "securityKey").
		Select("类型", "type", []*views.InputOptions{{Label: "虚拟用户", Value: models.UserTypeVirtual}, {Label: "普通用户", Value: models.UserTypeDefault}}).
		Select("状态", "status", []*views.InputOptions{{Label: "激活", Value: models.UserStatusActive}, {Label: "禁用", Value: models.UserStatusDisable}}).
		Text("数据", "data").TextArea("个性签名", "desc").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config:      views.NewDialogViews("update", "/user/update", "更新用户信息").SetParams(updateParams).SetInputList(updateInputList).SetFullHeight(),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
