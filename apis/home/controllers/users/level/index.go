package level

import (
	"basic/apis/home/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index
// @Summary	等级信息列表
// @Tags		User
// @Success	200	{object}	dtousers.LevelIndexData
// @Router		/api/v1/user/level/index [post]
func Index(c *fiber.Ctx) error {
	adminId, userId := utils.GetContextClaims(c)

	lang := c.Get("Accept-Language")
	data, err := users.LevelIndex(adminId, userId, lang)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
