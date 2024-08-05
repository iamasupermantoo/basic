package setting

import (
	dtosettings "basic/apis/admin/dto/settings"
	"basic/apis/admin/service/settings"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Index 设置列表
func Index(c *fiber.Ctx) error {
	params := &dtosettings.SettingIndexParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := settings.SettingIndex(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
