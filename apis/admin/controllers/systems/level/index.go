package level

import (
	dtosystems "basic/apis/admin/dto/systems"
	"basic/apis/admin/service/systems"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 查询等级配置列表
func Index(c *fiber.Ctx) error {
	params := &dtosystems.LevelIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	调用服务层方法
	adminId, userId := utils.GetContextClaims(c)
	data, err := systems.LevelIndex(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
