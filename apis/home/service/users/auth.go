package users

import (
	"basic/apis/common/rds"
	dtousers "basic/apis/home/dto/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"errors"
	"time"
)

// AuthInfo 获取认证列表
func AuthInfo(adminId, userId int) (interface{}, error) {
	data := &dtousers.HomeAuthIndexData{}
	database.Db.Model(models.RealNameAuth{}).Where("user_id = ?", userId).Find(data)
	return data, nil
}

// AuthCreate 新增认证
func AuthCreate(adminId, userId int, lang string, params *dtousers.AuthCreateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//	用户信息
	userInfo := &models.User{}
	if result := database.Db.Model(userInfo).Where("id = ?", userId).Where("status = ?", models.UserStatusActive).Take(userInfo); result.Error != nil {
		return nil, result.Error
	}

	currentAuthInfo := &models.RealNameAuth{}
	if result := database.Db.Model(currentAuthInfo).Where("user_id = ?", userInfo.Id).Where("status > ?", models.UserRealNameStatusDelete).Find(currentAuthInfo); result.RowsAffected == 0 {
		params.UserId = userId
		params.AdminId = adminId
		params.Type = models.UserRealNameTypeIdCard
		params.Status = models.UserRealNameStatusActive
		params.UpdatedAt = int(time.Now().Unix())
		params.CreatedAt = int(time.Now().Unix())
		if result = database.Db.Table("real_name_auth").Create(params); result.Error != nil {
			return nil, result.Error
		}
	} else {
		// 审核中
		if currentAuthInfo.Status == models.UserRealNameStatusActive || currentAuthInfo.Status == models.UserRealNameStatusComplete {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "pendingRealName"))
		}

		// 已实名
		if currentAuthInfo.Status == models.UserRealNameStatusComplete {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "realNameFailed"))
		}

		params.Status = models.UserRealNameStatusActive
		if result = database.Db.Model(&models.RealNameAuth{}).Where("id = ?", currentAuthInfo.Id).Updates(params); result.Error != nil {
			return nil, result.Error
		}

	}
	return "ok", nil
}
