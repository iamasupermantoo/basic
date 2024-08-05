package userassets

import (
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Info
// @Summary	用户资产列表
// @Tags Wallet
// @Param		id	body		int	true	"请求数据"
// @Success		200			{object}	utils.ResponseJson{data=[]dtowallets.HomeUserAssetsInfo}
// @Router		/api/v1/wallets/user/assets/index [post]
func Info(c *fiber.Ctx) error {
	params := &dtowallets.UserAssetsInfoParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	// 从请求头获取发送请求使用的语言。和管理id
	adminId, userId := utils.GetContextClaims(c)

	data, err := wallets.UserAssetsInfo(adminId, userId, c.Get("Accept-Language"), params)
	if err != nil {

		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
