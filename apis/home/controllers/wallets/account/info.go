package account

import (
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Info
// @Summary	卡片详情
// @Tags Wallet
// @Param		id			body		int			true	"主键"
// @Success		200			{object}	utils.ResponseJson{data=[]dtowallets.HomeAccountInfo}
// @Router		/api/v1/wallets/account/info [post]
func Info(c *fiber.Ctx) error {
	params := &dtowallets.AccountInfoParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.AccountInfo(adminId, userId, c.Get("Accept-Language"), params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
