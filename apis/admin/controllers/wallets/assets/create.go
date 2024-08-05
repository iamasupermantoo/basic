package assets

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/admin/service/wallets"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Create 资产创建
func Create(c *fiber.Ctx) error {
	params := &dtowallets.AssetsCreateParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	if err := validator.NewValidator(c).Validate(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.AssetsCreate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
