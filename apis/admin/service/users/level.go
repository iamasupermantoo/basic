package users

import (
	dtousers "basic/apis/admin/dto/users"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"errors"
	"time"
)

// LevelIndex 获取会员列表
func LevelIndex(adminId, userId int, params *dtousers.LevelOrderIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtousers.LevelOrderIndexData, 0)}

	//	过滤where条件参数
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		EqInt("level_id = ?", params.LevelId).
		EqInt("status = ?", params.Status).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt).
		BetweenDate("expired_at BETWEEN ? AND ?", params.ExpiredAt)

	// 查询数据库
	result := database.Db.Table("level_order as l").
		Select("l.id", "l.admin_id", "l.user_id", "l.level_id", "l.status", "l.data", "l.updated_at", "l.created_at", "l.expired_at", "a.username as adminName", "u.username as username").
		Joins("left join admin_user as a on l.admin_id=a.id").
		Joins("left join user as u on l.user_id=u.id").
		Where("l.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("l.status > ?", models.UserLevelOrderStatusDelete).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)
	return data, result.Error
}

// LevelCreate 新增会员
func LevelCreate(adminId, userId int, params *dtousers.LevelOrderCreateParams) (interface{}, error) {
	//	获取等级信息
	levelInfo := &models.Level{}
	if result := database.Db.Model(levelInfo).
		Where("id = ?", params.LevelId).
		Where("status = ?", models.LevelStatusActive).
		Take(levelInfo); result.Error != nil {
		return nil, result.Error
	}

	//	用户信息
	userInfo := &models.User{}
	if result := database.Db.Model(userInfo).
		Where("username = ?", params.UserName).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("status = ?", models.UserStatusActive).
		Take(userInfo); result.Error != nil {
		return nil, result.Error
	}

	//	当前用户必须是所属管理的用户等级
	if levelInfo.AdminId != rds.GetAdminSettingId(userInfo.AdminId) {
		return nil, errors.New("incorrectFormat")
	}

	//	过期时间
	nowTime := time.Now()
	expiredTime := nowTime.AddDate(0, 0, levelInfo.Days).Unix()
	if levelInfo.Days == -1 {
		expiredTime = nowTime.AddDate(10, 0, 0).Unix()
	}

	currentUserLevelInfo := &models.LevelOrder{}
	result := database.Db.Model(currentUserLevelInfo).
		Where("user_id = ?", userInfo.Id).
		Order("id desc").
		Find(currentUserLevelInfo)
	if result.RowsAffected == 0 {
		//	新增用户等级会员
		if result = database.Db.Model(&models.LevelOrder{}).Create(map[string]interface{}{
			"admin_id":   userInfo.AdminId,
			"user_id":    userInfo.Id,
			"level_id":   levelInfo.Id,
			"status":     models.UserLevelOrderStatusActive,
			"expired_at": expiredTime,
		}); result.Error != nil {
			return nil, result.Error
		}
	} else {
		//	更新用户等级会员
		if result = database.Db.Model(&models.LevelOrder{}).Where("id = ?", currentUserLevelInfo.Id).Updates(map[string]interface{}{
			"level_id":   levelInfo.Id,
			"status":     models.UserLevelOrderStatusActive,
			"expired_at": expiredTime,
		}); result.Error != nil {
			return nil, result.Error
		}
	}
	return "ok", nil
}

// LevelDelete 删除会员
func LevelDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.LevelOrder{}).
		Where("id IN ?", params.Ids).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Update("status", models.UserLevelOrderStatusDelete); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// LevelUpdate 更新会员
func LevelUpdate(adminId, userId int, params *dtousers.LevelOrderUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.LevelOrder{}).Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Updates(params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}
