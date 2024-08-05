package withdraw

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create
// @Summary	用户提现新增
// @Tags Wallet
// @Param		accountId		body		int			true	"卡片ID"
// @Param		money			body		number		true	"金额"
// @Param		securityKey		body		string		false	"安全密钥"
// @Success		200			{object}	utils.ResponseJson
// @Router		/api/v1/wallets/withdraw/create [post]
func Create(c *fiber.Ctx) error {
	params := &dtowallets.WithdrawCreateParams{}

	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	if err := validator.NewValidator(c).Validate(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := wallets.WithdrawCreate(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
