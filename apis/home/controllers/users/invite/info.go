package invite

import (
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Info
// @Tags User
// @Summary	查询邀请码
// @Success	200			{string}	Code "邀请码"
// @Router		/api/v1/user/invite/index [post]
func Info(c *fiber.Ctx) error {
	adminId, userId := utils.GetContextClaims(c)
	data, err := users.InviteInfo(adminId, userId)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
