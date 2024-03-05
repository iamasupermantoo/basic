package role

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/apis/common/validator"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Delete 删除角色
func Delete(c *fiber.Ctx) error {
	params := &dtoadmins.RoleDeleteParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//调用服务层流程
	data, err := admins.RoleDelete(params)

	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	return c.JSON(utils.SuccessJson(data))

}
