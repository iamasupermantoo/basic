package user

import (
	deousers "basic/apis/admin/dto/users"
	"basic/apis/admin/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 用户列表
func Index(c *fiber.Ctx) error {
	params := new(deousers.UserIndexParams)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := users.UserIndex(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
