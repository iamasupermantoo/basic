package middleware

import "github.com/gofiber/fiber/v2"

type Host struct {
	Fiber *fiber.App
}

const (
	// OtherDomain 其他域名
	OtherDomain = "*"
)

// Hosts 域名对应的App路由
var Hosts map[string]*Host = map[string]*Host{}

// InitDomainRoutingMiddleware 初始化域名路由中间件
func InitDomainRoutingMiddleware(hosts map[string]*Host) fiber.Handler {
	Hosts = hosts
	return func(c *fiber.Ctx) error {
		if _, ok := Hosts[c.Hostname()]; ok {
			//	后端APP对象
			Hosts[c.Hostname()].Fiber.Handler()(c.Context())
		} else {
			//	前台APP对象
			Hosts[OtherDomain].Fiber.Handler()(c.Context())
		}
		return nil
	}
}
