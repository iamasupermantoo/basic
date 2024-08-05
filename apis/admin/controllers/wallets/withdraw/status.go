package withdraw

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/admin/service/wallets"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Status 用户提现审核
func Status(c *fiber.Ctx) error {
	params := new(dtowallets.WithdrawStatusParams)
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminInd, userId := utils.GetContextClaims(c)
	data, err := wallets.WithdrawStatus(adminInd, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
