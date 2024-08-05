package logs

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 日志列表
func Index(c *fiber.Ctx) error {
	params := &dtoadmins.AdminLogsIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)

	//调用服务层逻辑
	data, _ := admins.AdminLogsIndex(adminId, userId, params)
	return c.JSON(utils.SuccessJson(data))
}
