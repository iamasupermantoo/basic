package menu

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/settings"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 菜单列表
func Index(c *fiber.Ctx) error {
	params := &dtoadmins.AdminSettingMenuIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}
	adminId, userId := utils.GetContextClaims(c)

	//调用服务层逻辑
	data, _ := settings.AdminSettingMenuIndex(adminId, userId, params)
	return c.JSON(utils.SuccessJson(data))
}
