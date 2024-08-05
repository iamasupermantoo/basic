package index

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// UpdatePassword 当前管理员更新密码
func UpdatePassword(c *fiber.Ctx) error {
	params := new(dtoadmins.AdminUpdatePasswordParams)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	params.Id, _ = utils.GetContextClaims(c)
	data, err := admins.AdminUpdatePassword(params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	return c.JSON(utils.NewResponseJson().Success(data))
}
