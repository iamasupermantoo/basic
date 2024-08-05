package country

import (
	dtosystems "basic/apis/admin/dto/systems"
	"basic/apis/admin/service/systems"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Update 更新国家配置列表
func Update(c *fiber.Ctx) error {
	params := &dtosystems.CountryUpdateParams{}
	err := c.BodyParser(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	调用服务层方法
	adminId, userId := utils.GetContextClaims(c)
	data, err := systems.CountryUpdate(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
