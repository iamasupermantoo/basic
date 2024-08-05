package auth

import (
	"basic/apis/home/service/users"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Info
// @Tags User
// @Summary	查询认证
// @Success	200			{object}	dtousers.HomeAuthIndexData
// @Router		/api/v1/user/auth/info [post]
func Info(c *fiber.Ctx) error {
	//	调用服务层注册流程
	adminId, userId := utils.GetContextClaims(c)
	data, err := users.AuthInfo(adminId, userId)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
