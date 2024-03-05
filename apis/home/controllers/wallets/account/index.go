package account

import (
	"basic/apis/common/validator"
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Index
// @Summary	卡片列表
// @Tags Wallet
// @Success		200			{object}	utils.ResponseJson{data=[]dtowallets.HomeAccountIndexData}
// @Router		/api/v1/wallets/account/index [post]
func Index(c *fiber.Ctx) error {
	params := &dtowallets.HomeAccountIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	lang := c.Get("Accept-Language")
	data, err := wallets.AccountIndex(adminId, userId, lang, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
