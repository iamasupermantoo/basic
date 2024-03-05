package user

import (
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Info
// @Summary	获取当前用户信息
// @Tags		User
// @Success	200	{object}	dtousers.HomeUserInfo
// @Router		/api/v1/user/info [post]
func Info(c *fiber.Ctx) error {
	//调用服务层
	_, userId := utils.GetContextClaims(c)

	data, err := users.UserInfo(userId)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	返回数据
	return c.JSON(utils.SuccessJson(data))
}
