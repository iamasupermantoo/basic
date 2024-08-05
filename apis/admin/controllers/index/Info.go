package index

import (
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Info 管理信息
func Info(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)

	//	调用服务层获取管理信息
	data, err := admins.AdminInfo(adminId)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}
	return c.JSON(utils.NewResponseJson().Success(data))
}
