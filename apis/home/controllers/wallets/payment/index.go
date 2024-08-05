package payment

import (
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index
// @Summary	支付列表
// @Tags Wallet
// @Param		modes	body		[]int	true	"模式 1充值模式 2资产充值模式 11提现模式 12资产提现模式"
// @Success		200			{object}	utils.ResponseJson{data=[]dtowallets.HomePaymentIndexData}
// @Router		/api/v1/wallets/payment/index [post]
func Index(c *fiber.Ctx) error {
	params := &dtowallets.PaymentIndexParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, _ := utils.GetContextClaims(c)
	data, err := wallets.PaymentIndex(adminId, c.Get("Accept-Language"), params)

	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
