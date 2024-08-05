package setting

import (
	dtosettings "basic/apis/admin/dto/settings"
	"basic/apis/admin/service/settings"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create 设置创建
func Create(c *fiber.Ctx) error {
	params := &dtosettings.SettingCreateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := settings.SettingCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
