package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

// InitRecoverMiddleware 初始化异常捕获
func InitRecoverMiddleware() fiber.Handler {
	var config = recover.Config{}
	return recover.New(config)
}
