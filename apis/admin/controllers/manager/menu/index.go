package menu

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 菜单列表
func Index(c *fiber.Ctx) error {
	params := &dtoadmins.AdminMenuIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}
	adminId, userId := utils.GetContextClaims(c)

	//调用服务层逻辑
	data, _ := admins.AdminMenuIndex(adminId, userId, params)
	return c.JSON(utils.SuccessJson(data))
}
