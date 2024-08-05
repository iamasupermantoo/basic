package withdraw

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/admin/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 用户提现列表
func Index(c *fiber.Ctx) error {
	params := &dtowallets.WithdrawIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.WithdrawIndex(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
