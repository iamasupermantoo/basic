package logs

import (
	"basic/utils"
	"basic/utils/views"

	"github.com/gofiber/fiber/v2"
)

// Views 模版JSON信息
func Views(c *fiber.Ctx) error {
	config := views.NewTableViews("/manage/logs/index", "")

	//	查询参数
	config.Search.Params, config.Search.InputList = views.NewInputViews().
		Text("管理名称", "adminName").
		Text("操作名称", "name").
		Text("Ip4地址", "ip").
		Text("路由", "route").
		RangeDatePicker("操作时间", "createdAt").
		GetInputListInfo()

	//	数据表格字段名称
	config.Table.Columns = views.NewColumnsViews().
		Text("ID", "id", true).
		Text("管理名称", "adminName", false).
		Text("操作名称", "name", true).
		Text("ip4地址", "ip", true).
		Text("路由", "route", true).
		Text("头信息", "headers", true).
		Text("内容", "body", true).
		Text("数据", "data", true).
		DatePicker("操作时间", "createdAt", true).
		GetColumnsListInfo()

	return c.JSON(utils.SuccessJson(config))
}
