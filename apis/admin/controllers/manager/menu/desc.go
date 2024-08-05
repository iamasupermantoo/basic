package menu

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Desc 更新菜单详情
func UpdateDesc(c *fiber.Ctx) error {
	params := &dtoadmins.AdminMenuUpdateDescParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := admins.AdminMenuUpdateDesc(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
