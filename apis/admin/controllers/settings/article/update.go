package article

import (
	dtosettings "basic/apis/admin/dto/settings"
	"basic/apis/admin/service/settings"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Update 文章更新
func Update(c *fiber.Ctx) error {
	params := &dtosettings.ArticleUpdateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := settings.ArticleUpdate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
