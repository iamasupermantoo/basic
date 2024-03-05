package middleware

import (
	"basic/apis/common/rds"
	"basic/models"
	"basic/utils"
	"crypto/rsa"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

const (
	// AuthScheme 标头
	AuthScheme = "Bearer"
)

// InitJwtMiddleware 初始化Jwt中间件
func InitJwtMiddleware(privateKey *rsa.PrivateKey) fiber.Handler {
	var config = jwtware.Config{
		ContextKey: "token",
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    privateKey.Public(),
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			adminId, userId, isRedisToken := models.RedisTokenParams(c)
			if isRedisToken {
				//	如果是后台APP
				if c.App().Config().AppName == models.ServiceAdminName {
					//	如果是后台路由, 那么验证当前管理路由是否有权限
					adminRouteList := rds.RedisFindAdminRouter(adminId)
					if utils.ArrayStringIndexOf(adminRouteList, "*") > -1 || utils.ArrayStringIndexOf(adminRouteList, c.Path()) > -1 {
						//	记录后台管理日志
						models.AdminLogsFunc(c, adminId, userId)
					}
				}

				return c.Next()
			}

			//	缓存中不存在Token
			return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired JWT")
		},
	}

	return jwtware.New(config)
}
