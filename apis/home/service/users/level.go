package users

import (
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	dtousers "basic/apis/home/dto/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"errors"
	"time"

	"gorm.io/gorm"
)

// LevelIndex 等级信息列表
func LevelIndex(adminId int, userId int, lang string) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	levelList := make([]*dtousers.LevelIndexData, 0)
	adminSettingId := rds.RedisFindAdminIdToSettingAdminId(rdsConn, adminId)

	result := database.Db.Model(&models.Level{}).
		Where("status=?", models.LevelStatusActive).
		Where("admin_id=?", adminSettingId).Order("level ASC").
		Scan(&levelList)

	for _, level := range levelList {
		level.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, level.Name)
		level.Desc = rds.RedisFindTranslateField(rdsConn, adminId, lang, level.Desc)
	}

	return levelList, result.Error
}

// LevelOrderCreate 当前用户购买等级
func LevelOrderCreate(adminId int, userId int, lang string, params *dtousers.LevelOrderParams) (interface{}, error) {
	//获取当前用户信息
	userInfo := &models.User{}
	result := database.Db.Model(userInfo).Where("id=?", userId).Take(userInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	//获取购买等级信息
	levelInfo := &models.Level{}
	result = database.Db.Model(levelInfo).
		Where("id=?", params.Id).
		Where("status=?", models.LevelStatusActive).
		Take(levelInfo)
	if result.Error != nil {
		return nil, result.Error
	}

	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	//获取当前用户等级信息
	nowTime := time.Now()
	buyLevelMoney := levelInfo.Money
	userLevelOrderInfo := &models.LevelOrder{}
	database.Db.Model(userLevelOrderInfo).
		Where("user_id=?", userInfo.Id).Where("status=?", models.UserLevelOrderStatusActive).Where("expired_at>?", nowTime.Unix()).
		Last(userLevelOrderInfo)

	//如果已经购买过会员，那么升级会员
	if userLevelOrderInfo.Id != 0 {
		userLevelInfo := &models.Level{}
		result = database.Db.Model(userLevelInfo).Where("id=?", userLevelOrderInfo.LevelId).Take(userLevelInfo)
		if result.Error != nil {
			return nil, result.Error
		}

		//如果用户已有会员等级大于购买会员等级,无法购买
		if userLevelInfo.Level >= levelInfo.Level {
			return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "incorrectFormat"))
		}

		//	如果已有会员, 并且设置差价购买, 那么扣掉之前的价格
		buyLevelMode := rds.RedisFindAdminSettingField(rdsConn, adminId, "buyLevelMode")
		if buyLevelMode == models.UserBuyLevelModePriceDifference {
			buyLevelMoney = buyLevelMoney - userLevelInfo.Money
		}
	}

	//	过期时间
	expiredTime := nowTime.AddDate(0, 0, levelInfo.Days).Unix()
	if levelInfo.Days == -1 {
		expiredTime = nowTime.AddDate(10, 0, 0).Unix()
	}

	//事务开始
	result.Error = database.Db.Transaction(func(tx *gorm.DB) error {
		//插入等级订单列表
		levelOrder := &models.LevelOrder{
			AdminId:   userInfo.AdminId,
			UserId:    userInfo.Id,
			LevelId:   levelInfo.Id,
			ExpiredAt: int(expiredTime),
			UpdatedAt: int(nowTime.Unix()),
			CreatedAt: int(nowTime.Unix()),
		}
		result = tx.Create(levelOrder)
		if result.Error != nil {
			return result.Error
		}

		// 其他会员状态改为-1
		result = tx.Model(&models.LevelOrder{}).Where("user_id = ?", userInfo.Id).Where("id <> ?", levelOrder.Id).Update("status", models.UserLevelOrderStatusDisable)
		if result.Error != nil {
			return result.Error
		}

		//写入账单
		result.Error = users.UserSpending(tx, &dto.WalletOrderAgreeParams{
			AdminId:  userInfo.AdminId,
			ParentId: userInfo.ParentId,
			UserId:   userInfo.Id,
			SourceId: levelOrder.Id,
			BillType: models.WalletBillTypeBuyLevel,
			Balance:  userInfo.Money,
			Money:    buyLevelMoney,
		})
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	return "ok", result.Error

}
