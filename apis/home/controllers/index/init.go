package index

import (
	"basic/apis/admin/service/users"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Init 用户初始化数据
// @Summary 用户初始化数据
// @Description 获取用户初始化数据
// @Tags Basic
// @Param domain query string true "域名"
// @Param lang query string false "语言"
// @Success 200 {object} dtousers.UserInitData
// @Failure 400 {object} error
// @Router /api/v1/init [Get]
func Init(c *fiber.Ctx) error {
	data := users.UserInit(c.Query("domain"), c.Query("lang"))
	return c.JSON(utils.SuccessJson(data))
}
