package bill

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/admin/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index  用户账单列表
func Index(c *fiber.Ctx) error {
	params := &dtowallets.BillIndexParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	// 获取路由管理ID
	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.BillIndex(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	return c.JSON(utils.NewResponseJson().Success(data))
}
