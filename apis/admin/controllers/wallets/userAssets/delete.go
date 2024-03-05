package userassets

import (
	"basic/apis/admin/service/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Delete 用户资产删除
func Delete(c *fiber.Ctx) error {
	params := &dto.DeleteParams{}
	if err := c.BodyParser(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	if err := validator.NewValidator(c).Validate(params); err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.UserAssetsDelete(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
