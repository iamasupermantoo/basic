package menu

import (
	"basic/apis/admin/service/settings"
	"basic/apis/common/dto"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Delete 删除菜单
func Delete(c *fiber.Ctx) error {
	params := &dto.DeleteParams{}
	err := c.BodyParser(params)
	if err != nil {
		return err
	}
	adminId, userId := utils.GetContextClaims(c)

	//调用服务层逻辑
	data, err := settings.AdminSettingMenuDelete(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
