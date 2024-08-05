package manage

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create 创建管理
func Create(c *fiber.Ctx) error {
	params := &dtoadmins.ManageCreateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	调用服务层流程
	adminId, userId := utils.GetContextClaims(c)
	data, err := admins.ManageCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
