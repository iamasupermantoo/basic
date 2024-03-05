package account

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Update
// @Summary	修改卡片
// @Tags Wallet
// @Param		id			body		int			true	"主键"
// @Param		paymentId	body		int			false	"主键"
// @Param		code		body		string		false	"银行代码｜token"
// @Param		name		body		string		false	"银行名称"
// @Param		realName	body		string		false	"真实姓名"
// @Param		number		body		string		false	"卡号"
// @Success		200			{object}	utils.ResponseJson
// @Router		/api/v1/wallets/account/update [post]
func Update(c *fiber.Ctx) error {
	params := &dtowallets.AccountUpdateParams{}

	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	if err := validator.NewValidator(c).Validate(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.AccountUpdate(adminId, userId, c.Get("Accept-Language"), params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
