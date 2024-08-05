package menu

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/settings"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// UpdateDesc 更新菜单详情
func UpdateDesc(c *fiber.Ctx) error {
	params := &dtoadmins.AdminSettingMenuUpdateDescParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := settings.AdminSettingMenuUpdateDesc(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))

}
