package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// InitLimiterMiddleware 初始化速率限制
func InitLimiterMiddleware() fiber.Handler {
	var config = limiter.Config{
		Max:        50,
		Expiration: 10 * time.Second,
	}
	return limiter.New(config)
}
