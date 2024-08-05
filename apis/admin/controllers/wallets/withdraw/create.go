package withdraw

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/admin/service/wallets"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create 用户提现创建
func Create(c *fiber.Ctx) error {
	params := &dtowallets.WithdrawCreateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.WithdrawCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
