package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// InitLoggerMiddleware 初始化日志中间件
func InitLoggerMiddleware() fiber.Handler {
	var config = logger.Config{}

	return logger.New(config)
}
