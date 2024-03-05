package users

import (
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	dtousers "basic/apis/home/dto/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"errors"
	"time"
)

// TeamIndex 获取团队成员列表
func TeamIndex(adminId, userId int, lang string, params *dtousers.TeamIndexParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	userChildrenIds := users.GetUserChildren(userId)
	currentUserInfo := &models.User{}
	if result := database.Db.Model(currentUserInfo).
		Where("id = ?", params.Id).
		Where("id IN ?", userChildrenIds).
		Take(currentUserInfo); result.Error != nil {
		return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "incorrectFormat"))
	}
	data := &dtousers.TeamIndexInfo{
		Id:        currentUserInfo.Id,
		Avatar:    currentUserInfo.Avatar,
		UserName:  currentUserInfo.UserName,
		NickName:  currentUserInfo.NickName,
		CreatedAt: currentUserInfo.CreatedAt,
		Depth:     users.GetUserDepth(currentUserInfo.Id, userId, 0),
		Children:  make([]*dtousers.TeamIndexInfo, 0),
	}
	currentUserChildrenIds := users.GetUserChildren(data.Id)
	if result := database.Db.Model(&models.WalletBill{}).
		Select("IFNULL(sum(money), 0)").
		Where("user_id <> ?", params.Id).
		Where("user_id IN ?", currentUserChildrenIds).
		Where("type = ?", models.WalletBillTypeTeamEarnings).
		Find(&data.Earnings); result.Error != nil {
		return nil, result.Error
	}

	// 获取下级用户列表
	if result := database.Db.Table("user as u").
		Select("u.id", "u.avatar", "u.username", "u.nickname", "IFNULL(SUM(wb.money), 0) as earnings", "u.created_at").
		Joins("LEFT JOIN wallet_bill as wb on u.id=wb.user_id AND wb.type = ?", models.WalletBillTypeTeamEarnings).
		Where("u.parent_id = ?", params.Id).
		Order("earnings desc").
		Group("u.id").Limit(50).
		Scan(&data.Children); result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

// TeamDetails 获取用户团队详情数据
func TeamDetails(adminId, userId int, lang string, params *dtousers.TeamDetailsParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	nowTime := time.Now()
	userChildrenIds := users.GetUserChildren(userId)
	todayTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)

	data := &dtousers.TeamDetailsData{}
	currentUserInfo := &models.User{}
	if result := database.Db.Model(currentUserInfo).Where("id = ?", params.Id).Where("id IN ?", userChildrenIds).Take(currentUserInfo); result.Error != nil {
		return nil, errors.New(rds.RedisFindTranslateField(rdsConn, adminId, lang, "incorrectFormat"))
	}
	data.CurrentInfo = &dtousers.TeamIndexInfo{
		Id:        currentUserInfo.Id,
		Avatar:    currentUserInfo.Avatar,
		UserName:  currentUserInfo.UserName,
		NickName:  currentUserInfo.NickName,
		CreatedAt: currentUserInfo.CreatedAt,
		Depth:     users.GetUserDepth(currentUserInfo.Id, userId, 0),
		Children:  make([]*dtousers.TeamIndexInfo, 0),
	}
	currentUserChildrenIds := users.GetUserChildren(data.CurrentInfo.Id)
	if result := database.Db.Model(&models.WalletBill{}).
		Select("IFNULL(sum(money), 0)").
		Where("user_id <> ?", params.Id).
		Where("user_id IN ?", currentUserChildrenIds).
		Where("user_id IN ?", userChildrenIds).
		Where("type = ?", models.WalletBillTypeTeamEarnings).
		Find(&data.CurrentInfo.Earnings); result.Error != nil {
		return nil, result.Error
	}
	data.Children = make([]*dtousers.WalletBillInfo, 0)

	// 总邀请人数
	if result := database.Db.Model(models.User{}).
		Where("id<>?", params.Id).
		Where("id IN ?", currentUserChildrenIds).
		Count(&data.InviteNums); result.Error != nil {
		return nil, result.Error
	}
	// 今日邀请人数
	if result := database.Db.Model(models.User{}).Where("id<>?", params.Id).
		Where("id IN ?", currentUserChildrenIds).
		Where("created_at between ? AND ?", todayTime.Unix(), nowTime.Unix()).
		Count(&data.TodayNums); result.Error != nil {
		return nil, result.Error
	}

	// 总购买
	if result := database.Db.Model(models.WalletBill{}).
		Select("IFNULL(sum(money), 0)").
		Where("user_id <> ?", params.Id).
		Where("user_id IN ?", currentUserChildrenIds).
		Where("type = ?", models.WalletBillTypeBuyProduct).
		Find(&data.BuyAmount); result.Error != nil {
		return nil, result.Error
	}
	// 今日购买
	if result := database.Db.Model(models.WalletBill{}).
		Select("IFNULL(sum(money), 0)").
		Where("user_id <> ?", params.Id).
		Where("user_id IN ?", currentUserChildrenIds).
		Where("created_at between ? AND ?", todayTime.Unix(), nowTime.Unix()).
		Where("type = ?", models.WalletBillTypeBuyProduct).
		Find(&data.TodayAmount); result.Error != nil {
		return nil, result.Error
	}

	// 总团队收益
	data.TeamEarnings = data.CurrentInfo.Earnings
	// 今日团队收益
	if result := database.Db.Model(models.WalletBill{}).
		Select("IFNULL(sum(money), 0)").
		Where("user_id <> ?", params.Id).
		Where("user_id IN ?", currentUserChildrenIds).
		Where("created_at between ? AND ?", todayTime.Unix(), nowTime.Unix()).
		Where("type = ?", models.WalletBillTypeTeamEarnings).
		Find(&data.TodayEarnings); result.Error != nil {
		return nil, result.Error
	}

	//团队收益列表
	if result := database.Db.Table("wallet_bill as wb").
		Select("wb.id", "wb.name", "wb.money", "wb.created_at", "u.avatar", "u.username", "u.nickname").
		Joins("LEFT JOIN user as u on wb.user_id = u.id").
		Where("wb.user_id IN ?", currentUserChildrenIds).
		Where("wb.type = ? ", models.WalletBillTypeTeamEarnings).
		Where("wb.user_id <> ?", params.Id).
		Order("wb.created_at desc").
		Limit(50).
		Scan(&data.Children); result.Error != nil {
		return nil, result.Error
	}
	for _, child := range data.Children {
		child.Name = rds.RedisFindTranslateField(rdsConn, adminId, lang, child.Name)
	}

	return data, nil
}
