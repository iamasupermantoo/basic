package users

import (
	commonUsers "basic/apis/common/service/users"
	dtousers "basic/apis/home/dto/users"
	"basic/models"
	"basic/module/database"
)

// InviteInfo 获取邀请列表
func InviteInfo(adminId, userId int) (interface{}, error) {
	data := &dtousers.InviteInfo{}
	if result := database.Db.Model(&models.Invite{}).Where("status > ?", models.UserInviteStatusDelete).Where("admin_id = ?", adminId).Where("user_id = ?", userId).Find(data); result.RowsAffected == 0 {
		userInfo := &models.User{}
		if result = database.Db.Model(userInfo).Where("id = ?", userId).Take(userInfo); result.Error != nil {
			return nil, result.Error
		}

		inviteInfo, err := commonUsers.InviteCreate(adminId, userId, &dtousers.InviteCreateParams{
			UserName: userInfo.UserName,
			Type:     models.UserInviteTypeUser,
		})
		if err != nil {
			return nil, err
		}
		data.Code = inviteInfo.(*models.Invite).Code
	}

	return data, nil
}
