package index

import (
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 首页信息
func Index(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)

	//	调用服务层获取管理信息
	data, _ := admins.Dashboard(adminId)
	return c.JSON(utils.NewResponseJson().Success(data))
}
