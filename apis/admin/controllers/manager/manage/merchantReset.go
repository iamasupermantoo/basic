package manage

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// MerchantReset 商户重置
func MerchantReset(c *fiber.Ctx) error {
	params := &dtoadmins.MerchantParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	data, err := admins.MerchantReset(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))

}
