package menu

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/settings"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create 新增菜单
func Create(c *fiber.Ctx) error {
	params := &dtoadmins.AdminSettingMenuCreateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}
	adminId, userId := utils.GetContextClaims(c)

	//调用服务层逻辑
	data, err := settings.AdminSettingMenuCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
