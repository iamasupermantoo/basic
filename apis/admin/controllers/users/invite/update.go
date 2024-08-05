package invite

import (
	dtousers "basic/apis/admin/dto/users"
	"basic/apis/admin/service/users"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

func Update(c *fiber.Ctx) error {
	params := new(dtousers.InviteUpdateParams)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := users.InviteUpdate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
