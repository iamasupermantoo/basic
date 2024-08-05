package database

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/models"
	"basic/utils"
)

// AdminUserInit 初始化管理表
func AdminUserInit() []*models.AdminUser {

	//	插入初始化数据
	admPriKey, admPubKey := utils.GenerateRsaKey()
	homePriKey, homePubKey := utils.GenerateRsaKey()
	superData := &models.AdminUserSuperData{
		RsaList: map[string]*models.RsaSetting{
			models.ServiceAdminName: {
				PriKey: string(admPriKey),
				PubKey: string(admPubKey),
			},
			models.ServiceHomeName: {
				PriKey: string(homePriKey),
				PubKey: string(homePubKey),
			},
		},
	}
	password := utils.PasswordEncrypt("taozijun")
	otherPassword := utils.PasswordEncrypt("abc123")

	adminData := &dtoadmins.AdminInfoData{
		Template: "default",
		Nums:     5,
	}

	return []*models.AdminUser{
		{Id: models.SuperAdminId, ParentId: 0, UserName: "admin", NickName: "八戒网络科技", Email: "muiprosperyls15@gmail.com", Password: password, SecurityKey: password, Data: utils.JsonMarshal(superData)},
		{Id: models.AdminMerchat, ParentId: 1, UserName: "merchant", Domains: "127.0.0.1:8089,localhost:8089,basic.ainn.us", NickName: "八戒网络科技", Email: "muiprosperyls15@gmail.com", Password: otherPassword, SecurityKey: otherPassword, Data: utils.JsonMarshal(adminData)},
	}
}
