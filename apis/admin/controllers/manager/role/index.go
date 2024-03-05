package role

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 角色列表
func Index(c *fiber.Ctx) error {
	params := &dtoadmins.RoleIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	//调用服务层流程
	data, _ := admins.RoleIndex(params)

	return c.JSON(utils.SuccessJson(data))
}
