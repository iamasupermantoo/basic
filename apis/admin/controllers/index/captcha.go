package index

import (
	"strconv"

	"github.com/dchest/captcha"
	"github.com/gofiber/fiber/v2"
)

// NewCaptcha 创建验证码
func NewCaptcha(c *fiber.Ctx) error {
	c.SendString(captcha.NewLen(4))
	return nil
}

// Captcha 显示验证码
func Captcha(c *fiber.Ctx) error {
	width, _ := strconv.Atoi(c.Params("width"))
	height, _ := strconv.Atoi(c.Params("height"))

	c.Type("png")
	return captcha.WriteImage(c.Context().Response.BodyWriter(), c.Params("id"), width, height)
}
