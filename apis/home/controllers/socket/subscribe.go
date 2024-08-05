package socket

import (
	"basic/apis/common/validator"
	dtosockets "basic/apis/home/dto/sockets"
	"basic/apis/home/service/sockets"
	"basic/utils"
	"github.com/gofiber/fiber/v2"
)

// Subscribe 订阅消息
// @Summary 订阅消息
// @Tags websocket
// @Param uuid body string ture "用户唯一uuid"
// @Param event body string ture "事件"
// @Param data body interface{} ture "数据"
// @Success 200 {object} dtosockets.SubscribeParam
// @Router /api/v1/article/index [post]
func Subscribe(c *fiber.Ctx) error {
	params := new(dtosockets.SubscribeParam)
	err := c.BodyParser(&params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	//	验证参数
	err = validator.NewValidator(c).Validate(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}

	data, err := sockets.Subscribe(params)
	if err != nil {
		return c.JSON(utils.ErrorJson(err))
	}
	return c.JSON(utils.SuccessJson(data))
}
