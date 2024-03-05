package invite

import (
	"basic/apis/common/service/users"
	"basic/apis/common/validator"
	dtousers "basic/apis/home/dto/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

func Create(c *fiber.Ctx) error {
	params := new(dtousers.InviteCreateParams)
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
	data, err := users.InviteCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
