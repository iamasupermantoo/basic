package manage

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 管理列表
func Index(c *fiber.Ctx) error {
	params := &dtoadmins.ManageIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	//调用服务层流程
	adminId, userId := utils.GetContextClaims(c)
	data, _ := admins.ManageIndex(adminId, userId, params)
	return c.JSON(utils.SuccessJson(data))
}
