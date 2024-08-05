package menu

import (
	"basic/apis/admin/service/admins"
	"basic/apis/common/dto"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Delete 后台菜单删除
func Delete(c *fiber.Ctx) error {
	params := &dto.DeleteParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}
	adminId, userId := utils.GetContextClaims(c)

	//调用服务层逻辑
	data, err := admins.AdminMenuDelete(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
