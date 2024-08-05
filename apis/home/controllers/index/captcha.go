package index

import (
	"strconv"

	"github.com/dchest/captcha"
	"github.com/gofiber/fiber/v2"
)

// NewCaptcha 创建验证码
// @Summary 创建验证码
// @Tags Basic
// @Success 200 {object} string
// @Router /api/v1/captcha/create [get]
func NewCaptcha(c *fiber.Ctx) error {
	c.SendString(captcha.NewLen(4))
	return nil
}

// Captcha
// @Summary 显示验证码
// @Tags Basic
// @Param id path string true "验证码ID"
// @Param width query string true "验证码宽-高，例：/api/v1/captcha/dhgghwrgsdfafwqt/400-200"
// @Success 200 {object} string
// @Router /api/v1/captcha/:id/:width-:height [get]
func Captcha(c *fiber.Ctx) error {
	width, _ := strconv.Atoi(c.Params("width"))
	height, _ := strconv.Atoi(c.Params("height"))
	c.Type("png")
	return captcha.WriteImage(c.Context().Response.BodyWriter(), c.Params("id"), width, height)
}
