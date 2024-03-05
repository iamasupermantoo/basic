package role

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Update 更新角色
func Update(c *fiber.Ctx) error {
	params := &dtoadmins.RoleUpdateParams{}
	c.BodyParser(params)

	//	验证参数
	err := validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	调用服务层流程
	data, err := admins.RoleUpdate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))
}
