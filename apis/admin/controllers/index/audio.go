package index

import (
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Audio 音源
func Audio(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)

	//	调用服务层音源
	data, _ := admins.Audio(adminId)
	return c.JSON(utils.NewResponseJson().Success(data))
}
