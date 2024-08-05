package country

import (
	"basic/apis/admin/service/systems"
	"basic/apis/common/dto"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Delete 删除等级配置列表
func Delete(c *fiber.Ctx) error {
	params := &dto.DeleteParams{}
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
	data, err := systems.CountryDelete(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
