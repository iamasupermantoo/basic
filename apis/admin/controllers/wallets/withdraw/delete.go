package withdraw

import (
	"basic/apis/admin/service/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Delete 用户提现删除
func Delete(c *fiber.Ctx) error {
	params := &dto.DeleteParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := wallets.WithdrawDelete(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
