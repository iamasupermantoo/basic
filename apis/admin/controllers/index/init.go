package index

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"

	"github.com/gofiber/fiber/v2"
)

// Init 管理初始化数据
func Init(c *fiber.Ctx) error {
	adminId, _ := utils.GetContextClaims(c)
	adminInfo := &dtoadmins.AdminInfo{}
	result := database.Db.Model(&models.AdminUser{}).Where("id = ?", adminId).Take(adminInfo)
	if result.Error != nil {
		return c.JSON(utils.ErrorJson(result.Error))
	}

	//	获取管理路由列表 获取管理菜单
	routerList, menuList := admins.GetAdminRouterMenuList(adminId)

	//	管理配置
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	config := &dtoadmins.AdminInitConfig{
		Name: rds.RedisFindAdminSettingField(rdsConn, adminId, "siteName"),
		Logo: rds.RedisFindAdminSettingField(rdsConn, adminId, "siteLogo"),
	}
	return c.JSON(utils.SuccessJson(&dtoadmins.AdminInitData{
		Config:     config,
		RouterList: routerList,
		MenuList:   menuList,
		AdminInfo:  adminInfo,
	}))
}
