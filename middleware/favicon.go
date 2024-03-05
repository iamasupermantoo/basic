package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

// InitFaviconMiddleware 初始化 favicon 中间
func InitFaviconMiddleware() fiber.Handler {
	return favicon.New(favicon.Config{
		File: "./public/images/favicon.ico",
		URL:  "/favicon.ico",
	})
}
