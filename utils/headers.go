package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GetContextClaims 获取上下文Claims 管理ID 用户ID
func GetContextClaims(c *fiber.Ctx) (int, int) {
	user := c.Locals("token").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims == nil {
		panic("Authorization Failed")
	}

	return int(claims["adminId"].(float64)), int(claims["userId"].(float64))
}

// GetHeadersToken 获取头信息Token值
func GetHeadersToken(c *fiber.Ctx) string {
	headers := c.GetReqHeaders()
	if _, ok := headers["Authorization"]; ok {
		if len(headers["Authorization"]) > 0 {
			authorizationList := strings.Split(headers["Authorization"][0], " ")
			if len(authorizationList) == 2 {
				return authorizationList[1]
			}
		}
	}

	return ""
}
