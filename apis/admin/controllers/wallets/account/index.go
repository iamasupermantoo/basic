package account

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/admin/service/wallets"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 用户提现账户列表
func Index(c *fiber.Ctx) error {
	params := &dtowallets.AccountIndexParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	// 获取路由管理权限
	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.AccountIndex(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
