package payment

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/admin/service/wallets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// UpdateDesc 更新支付数据
func UpdateDesc(c *fiber.Ctx) error {
	params := &dtowallets.PaymentDescParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.PaymentUpdateDesc(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
