package socket

import (
	"basic/apis/common/validator"
	dtosockets "basic/apis/home/dto/sockets"
	"basic/apis/home/service/sockets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// BindUserId
// @Summary 绑定用户
// @Tags websocket
// @Param uuid body string ture "用户唯一uuid"
// @Success 200 {object} dtosockets.BindParams
// @Router /api/v1/article/index [post]
func BindUserId(c *fiber.Ctx) error {
	params := new(dtosockets.BindParams)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	adminId, userId := utils.GetContextClaims(c)
	data, err := sockets.BindUserId(adminId, userId, params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
