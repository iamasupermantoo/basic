package article

import (
	"basic/models"
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

// Views 模版JSON信息
func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/setting/article/index", "/setting/article/update")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("标题", "name").
		Text("内容", "content").
		Text("数据", "data").
		Select("类型", "type", []*views.InputOptions{{Label: "基础文章", Value: models.ArticleTypeBasic}, {Label: "帮助中心", Value: models.ArticleTypeHelpers}, {Label: "底部类型", Value: models.ArticleTypeFooterAbout}, {Label: "平台服务", Value: models.ArticleTypeFooterService}}).
		Select("状态", "status", []*views.InputOptions{{Label: "禁用", Value: models.ArticleStatusDisable}, {Label: "开启", Value: models.ArticleStatusActive}}).
		RangeDatePicker("更新时间", "updatedAt").
		GetInputListInfo()

	//	Table 工具栏目
	createParams, createInputList := views.NewInputViews().
		Select("类型", "type", []*views.InputOptions{{Label: "基础文章", Value: models.ArticleTypeBasic}, {Label: "帮助中心", Value: models.ArticleTypeHelpers}, {Label: "底部类型", Value: models.ArticleTypeFooterAbout}, {Label: "平台服务", Value: models.ArticleTypeFooterService}}).
		Text("标题", "name").
		Editor("内容", "content").
		SetValue("type", models.ArticleTypeBasic).
		GetInputListInfo()

	toolsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "批量删除", Color: views.ColorNegative, Size: "md"},
			Config: views.NewDialogViews("delete", "/setting/article/delete", "文章批量删除").
				SetSizingSmall().SetParams(&views.CheckboxListOptions{Operate: views.OperateCheckbox, Name: "name", Field: "ids", Scan: "id"}),
		},
		{
			ButtonViews: views.ButtonViews{Label: "新增文章", Color: views.ColorPrimary, Size: "md"},
			Config:      views.NewDialogViews("create", "/setting/article/create", "新增文章").SetParams(createParams).SetInputList(createInputList),
		},
	}
	config.Table.Tools = append(config.Table.Tools, toolsButtonDialogList...)

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Select("类型", "type", true, []*views.InputOptions{{Label: "基础文章", Value: models.ArticleTypeBasic}, {Label: "帮助中心", Value: models.ArticleTypeHelpers}, {Label: "关于我们", Value: models.ArticleTypeFooterAbout}, {Label: "平台服务", Value: models.ArticleTypeFooterService}}).
		Translate("标题", "name", true).
		Translate("内容", "content", true).
		EditTextArea("数据", "data", true).
		EditToggle("状态", "status", true, []*views.InputOptions{{Label: "开启", Value: models.ArticleStatusActive}, {Label: "禁用", Value: models.ArticleStatusDisable}}).
		DatePicker("更新时间", "updatedAt", true).
		GetColumnsListInfo()

	//	数据表格参数
	updateParams, updateInputList := views.NewInputViews().
		Image("广告图", "image").
		Select("类型", "type", []*views.InputOptions{{Label: "基础文章", Value: models.ArticleTypeBasic}, {Label: "帮助中心", Value: models.ArticleTypeHelpers}, {Label: "底部类型", Value: models.ArticleTypeFooterAbout}, {Label: "平台服务", Value: models.ArticleTypeFooterService}}).
		Text("标题", "name").
		Editor("内容", "content").
		GetInputListInfo()

	optionsButtonDialogList := []*views.DialogButtonViews{
		{
			ButtonViews: views.ButtonViews{Label: "更新", Color: views.ColorPrimary, Size: "xs"},
			Config: views.NewDialogViews("update", "/setting/article/update", "更新文章").
				SetParams(updateParams).SetInputList(updateInputList),
		},
	}
	config.Table.Options = append(config.Table.Options, optionsButtonDialogList...)

	return c.JSON(utils.SuccessJson(config))
}
