package menu

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/settings"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	params := &dtoadmins.AdminSettingMenuUpdateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}
	adminId, userId := utils.GetContextClaims(c)

	//调用服务层逻辑
	data, _ := settings.AdminSettingMenuUpdate(adminId, userId, params)
	return c.JSON(utils.SuccessJson(data))
}
