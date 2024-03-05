package database

import (
	"basic/models"
	"basic/utils"
)

// UserInit 用户表初始化
func UserInit() []*models.User {
	//	插入用户表初始化
	password := utils.PasswordEncrypt("abc123")
	return []*models.User{
		{Id: 1, AdminId: models.AdminMerchat, UserName: "ceshi", NickName: "Jsona", Email: "muiprosperyls15@gmail.com", Password: password, SecurityKey: password},
		{Id: 2, AdminId: models.AdminMerchat, ParentId: 1, UserName: "ceshi1", Money: 100, NickName: "Jsona", Password: password, SecurityKey: password},
		{Id: 3, AdminId: models.AdminMerchat, ParentId: 1, UserName: "ceshi2", Money: 100, NickName: "Jsona", Password: password, SecurityKey: password},
		{Id: 4, AdminId: models.AdminMerchat, ParentId: 1, UserName: "ceshi3", Money: 100, NickName: "Jsona", Password: password, SecurityKey: password},
		{Id: 5, AdminId: models.AdminMerchat, ParentId: 1, UserName: "ceshi4", Money: 100, NickName: "Jsona", Password: password, SecurityKey: password},
	}
}
