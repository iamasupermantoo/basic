package account

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Delete
// @Summary	删除卡片
// @Tags Wallet
// @Param		id			body		int			true	"主键"
// @Param		securityKey	body		int			false	"安全码"
// @Success		200			{object}	utils.ResponseJson
// @Router		/api/v1/wallets/account/delete [post]
func Delete(c *fiber.Ctx) error {
	params := &dtowallets.HomeAccountDeleteParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)

	data, err := wallets.AccountDelete(adminId, userId, c.Get("Accept-Language"), params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
