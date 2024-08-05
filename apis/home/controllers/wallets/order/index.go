package order

import (
	dtowallets "basic/apis/home/dto/wallets"
	"basic/apis/home/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index
// @Summary	钱包订单列表
// @Tags Wallet
// @Param		assetsId	body		int					false	"资产ID"
// @Param		types		body		[]int				false	"类型 1充值类型 2资产充值类型 11提现类型 12资产提现类型"
// @Param		pagination	body		utils.Pagination	false	"分页参数"
// @Success		200			{object}	utils.ResponseJson{data=[]dtowallets.HomeWalletOrderIndexData}
// @Router		/api/v1/wallets/order/index [post]
func Index(c *fiber.Ctx) error {
	params := &dtowallets.OrderIndexParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.OrderIndex(adminId, userId, c.Get("Accept-Language"), params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
