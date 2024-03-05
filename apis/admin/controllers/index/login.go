package index

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Login 管理登录
func Login(c *fiber.Ctx) error {
	params := &dtoadmins.AdminLoginParams{}
	err := c.BodyParser(params)
	if err != nil {
		panic(err.Error())
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	//	执行服务层
	data, err := admins.AdminLogin(params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	return c.JSON(utils.NewResponseJson().Success(data))
}
