package models

import (
	"basic/module/database"
	"basic/utils"
	"encoding/json"
	"net"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

// AdminRouterList 后台路由列表
var AdminRouterList = map[string]string{}

// AdminLogs 管理日志记录
type AdminLogs struct {
	Id        int    `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int    `gorm:"type:int unsigned not null;comment:管理ID"`
	Name      string `gorm:"type:varchar(120) not null;comment:名称"`
	Ip        uint32 `gorm:"type:int unsigned not null;comment:IP4"`
	Headers   string `gorm:"type:text;comment:头信息"`
	Route     string `gorm:"type:varchar(60) not null;comment:路由"`
	Body      string `gorm:"type:text;comment:内容"`
	Data      string `gorm:"type:text;comment:数据"`
	CreatedAt int    `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

// InitAdminRouter 载入后端所有路由
func InitAdminRouter(app *fiber.App) {
	routerList := app.GetRoutes()
	for i := 0; i < len(routerList); i++ {
		if routerList[i].Name != "" {
			AdminRouterList[routerList[i].Path] = routerList[i].Name
		}
	}
}

// AdminLogsFunc 管理日志方法
func AdminLogsFunc(c *fiber.Ctx, adminId, userId int) {
	//	过滤不需要的日志
	pathList := strings.Split(c.Path(), "/")
	filterList := []string{"audio", "index", "views", "init", "info", "fields", "google"}
	if utils.ArrayStringIndexOf(filterList, pathList[len(pathList)-1]) > -1 {
		return
	}

	headersBytes, _ := json.Marshal(c.GetReqHeaders())
	database.Db.Create(&AdminLogs{
		AdminId:   adminId,
		Name:      AdminRouterList[c.Path()],
		Ip:        utils.Ip4ToInt(net.ParseIP(c.IP())),
		Headers:   string(headersBytes),
		Route:     c.Path(),
		Body:      string(c.Body()),
		CreatedAt: int(time.Now().Unix()),
	})
}
