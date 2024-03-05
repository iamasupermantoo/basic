package account

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Create
// @Summary	新增卡片
// @Tags Wallet
// @Param		paymentId	body		integer			true	"支付Id"
// @Param		realName	body		string 			true 	"真实姓名"
// @Param		number 		body 		string 			true 	"卡号｜token"
// @Param		code 		body 		string 			false 	"银行代码｜地址"
// @Success		200			{object}	utils.ResponseJson
// @Router		/api/v1/wallets/account/create [post]
func Create(c *fiber.Ctx) error {
	params := &dtowallets.AccountCreateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)

	data, err := wallets.AccountCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
