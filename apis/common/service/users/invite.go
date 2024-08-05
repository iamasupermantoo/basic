package users

import (
	"basic/apis/common/rds"
	dtousers "basic/apis/home/dto/users"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"time"
)

// InviteCreate 新增邀请
func InviteCreate(adminId, userId int, params *dtousers.InviteCreateParams) (interface{}, error) {
	var inviteAdminId, inviteUserId int
	switch params.Type {
	case models.UserInviteTypeAdmin:
		adminInfo := &models.AdminUser{}
		result := database.Db.Model(adminInfo).Where("username = ?", params.UserName).Where("id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Where("status = ?", models.AdminUserStatusActive).Take(adminInfo)
		if result.Error != nil {
			return nil, result.Error
		}

		inviteAdminId = adminInfo.Id
	case models.UserInviteTypeUser:
		//	用户信息
		userInfo := &models.User{}
		result := database.Db.Model(userInfo).Where("username = ?", params.UserName).Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Where("status = ?", models.UserStatusActive).Take(userInfo)
		if result.Error != nil {
			return nil, result.Error
		}
		inviteAdminId = userInfo.AdminId
		inviteUserId = userInfo.Id
	}

	currentInviteInfo := &models.Invite{}
	result := database.Db.Model(currentInviteInfo).Where("admin_id = ?", inviteAdminId).Where("user_id = ?", inviteUserId).Find(currentInviteInfo)
	if result.RowsAffected == 0 {
		//	新增用户邀请码
		result = database.Db.Model(&models.Invite{}).Create(map[string]interface{}{
			"admin_id":   inviteAdminId,
			"user_id":    inviteUserId,
			"code":       utils.NewRandom().SetNumberRunes().String(6),
			"status":     models.UserInviteStatusActive,
			"updated_at": time.Now().Unix(),
			"created_at": time.Now().Unix(),
		})
		if result.Error != nil {
			return nil, result.Error
		}
	} else {
		result = database.Db.Model(&models.Invite{}).Where("id = ?", currentInviteInfo.Id).Updates(map[string]interface{}{
			"status":     models.UserInviteStatusActive,
			"updated_at": time.Now().Unix(),
		})
		if result.Error != nil {
			return nil, result.Error
		}
	}
	return "ok", nil
}
