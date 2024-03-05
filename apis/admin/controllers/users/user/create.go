package user

import (
	deousers "basic/apis/admin/dto/users"
	"basic/apis/admin/service/users"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create 新增用户
func Create(c *fiber.Ctx) error {
	params := new(deousers.UserCreateParams)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	调用服务层注册流程
	adminId, userId := utils.GetContextClaims(c)
	data, err := users.UserCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
