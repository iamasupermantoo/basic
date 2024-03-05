package userassets

import (
	"basic/apis/home/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index
// @Summary	用户资产列表
// @Tags Wallet
// @Success		200			{object}	utils.ResponseJson{data=[]dtowallets.HomeUserAssetsInfo}
// @Router		/api/v1/wallets/user/assets/index [post]
func Index(c *fiber.Ctx) error {
	adminId, userId := utils.GetContextClaims(c)

	data, err := wallets.UserAssetsIndex(adminId, userId, c.Get("Accept-Language"))
	if err != nil {

		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
