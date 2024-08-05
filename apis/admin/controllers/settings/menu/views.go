package menu

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

// Views 菜单Json模板
func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/setting/menu/index", "/setting/menu/update")

	//	查询参数设置
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("父级菜单", "parentName").
		Text("菜单名称", "name").
		Text("菜单路由", "route").
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.AdminMenuStatusDisable}, {Label: "开启", Value: models.AdminMenuStatusActive}}).
		Select("类型", "type", []*views.InputOptions{{Label: "前台菜单", Value: models.AdminMenuTypeHomeMenu}, {Label: "用户菜单", Value: models.AdminMenuTypeUserMenu}, {Label: "快捷菜单", Value: models.AdminMenuTypeQuickMenu}, {Label: "H5 Tabbar菜单", Value: models.AdminMenuTypeTabularMenu}}).
		RangeDatePicker("创建时间", "CreatedAt").GetInputListInfo()

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("父级ID", "parentId", true).
		Text("父级菜单", "parentName", false).
		Translate("菜单名称", "name", true).
		EditText("菜单路由", "route", false).
		EditNumber("排序", "sort", true).
		Select("类型", "type", true, []*views.InputOptions{{Label: "前台菜单", Value: models.AdminMenuTypeHomeMenu}, {Label: "用户菜单", Value: models.AdminMenuTypeUserMenu}, {Label: "快捷菜单", Value: models.AdminMenuTypeQuickMenu}, {Label: "H5 Tabbar菜单", Value: models.AdminMenuTypeTabularMenu}}).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "开启", Value: models.AdminMenuStatusActive}, {Label: "禁用", Value: models.AdminMenuStatusDisable}}).
		EditTextArea("数据", "data", true).
		DatePicker("创建时间", "createdAt", true).
		GetColumnsListInfo()

	//	新增角色参数设置
	createParams, createInputList := views.NewInputViews().
		Select("菜单类型", "type", []*views.InputOptions{{Label: "前台菜单", Value: models.AdminMenuTypeHomeMenu}, {Label: "用户菜单", Value: models.AdminMenuTypeUserMenu}, {Label: "快捷菜单", Value: models.AdminMenuTypeQuickMenu}, {Label: "H5 Tabbar菜单", Value: models.AdminMenuTypeTabularMenu}}).
		Text("名称", "name").
		Text("路由", "route").
		SetValue("type", models.AdminMenuTypeHomeMenu).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "新增菜单", Color: views.ColorPrimary, Size: "md"},
			Config: views.NewDialogViews("create", "/setting/menu/create", "新增菜单").
				SetParams(createParams).SetSizingSmall().
				SetInputList(createInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "删除菜单", Color: views.ColorNegative, Size: "md"},
			Config:      views.NewDialogViews("delete", "/setting/menu/delete", "批量删除菜单").SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
	}

	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	更新菜单
	updateParams, updateInputList := views.NewInputViews().
		Number("父级ID", "parentId").
		Text("菜单名称", "name").
		Text("菜单路由", "route").
		Number("排序", "sort").
		Select("菜单类型", "type", []*views.InputOptions{{Label: "前台菜单", Value: models.AdminMenuTypeHomeMenu}, {Label: "用户菜单", Value: models.AdminMenuTypeUserMenu}, {Label: "快捷菜单", Value: models.AdminMenuTypeQuickMenu}, {Label: "H5 Tabbar菜单", Value: models.AdminMenuTypeTabularMenu}}).
		TextArea("数据", "data").
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.AdminMenuStatusDisable}, {Label: "开启", Value: models.AdminMenuStatusActive}}).
		GetInputListInfo()

	//	更新详情
	dataJsonParam, dataJsonInputList := views.NewInputViews().
		Image("默认图标", "icon").Image("激活图标", "activeIcon").
		Text("vue模版文件", "vueTmp").Text("视图请求路由", "viewsUrl").
		Toggle("是否显示Pc端", "isPc", []*views.InputOptions{{Label: "显示", Value: true}, {Label: "隐藏", Value: false}}).Toggle("是否显示手机端", "isMobile", []*views.InputOptions{{Label: "显示", Value: true}, {Label: "隐藏", Value: false}}).
		GetInputListInfo()

	optionsParam, optionsInputList := views.NewInputViews().
		Json("", "dataJson", dataJsonInputList).SetValue("dataJson", dataJsonParam).
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/setting/menu/update", "更新菜单信息").
				SetParams(updateParams).SetSizingSmall().
				SetInputList(updateInputList),
		},
		{
			ButtonViews: views.ButtonViews{Label: "详情", Color: views.ColorSecondary, Size: "xs"},
			Config: views.NewDialogViews("update", "/setting/menu/update/desc", "更新菜单详情").
				SetParams(optionsParam).SetSizingSmall().
				SetInputList(optionsInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
