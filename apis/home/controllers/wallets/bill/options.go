package bill

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Options
// @Summary	用户账单类型
// @Tags Wallet
// @Param		type		body	int true "类型 1充值类型 11提现类型 21购买 51收益 61奖励"
// @Success		200			{object}	utils.ResponseJson{data=[]models.WalletBillTypeOptionsInfo}
// @Router		/api/v1/wallets/bill/index [post]
func Options(c *fiber.Ctx) error {
	params := &dtowallets.BillOptionsParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err := validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := wallets.BillOptions(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
