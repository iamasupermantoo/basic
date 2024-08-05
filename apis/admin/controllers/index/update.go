package index

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Update 更新当前管理员信息
func Update(c *fiber.Ctx) error {
	params := new(dtoadmins.AdminUpdateParams)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}

	//	从Token用获取管理ID
	params.Id, _ = utils.GetContextClaims(c)

	//	调用服务层更新当前管理员
	data, err := admins.AdminUpdate(params)
	if err != nil {
		return c.JSON(utils.NewResponseJson().Error(err.Error()))
	}
	return c.JSON(utils.NewResponseJson().Success(data))
}
