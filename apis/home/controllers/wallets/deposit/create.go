package deposit

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create
// @Summary	用户充值新增
// @Tags Wallet
// @Param		paymentId	body		int			true	"支付ID,支付模式：1充值模式，2资产充值模式"
// @Param		money		body		number		true	"金额"
// @Param		voucher		body		string 		true 	"支付凭证"
// @Success		200			{object}	utils.ResponseJson
// @Router		/api/v1/wallets/deposit/create [post]
func Create(c *fiber.Ctx) error {
	params := &dtowallets.DepositCreateParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	if err := validator.NewValidator(c).Validate(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := wallets.DepositCreate(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
