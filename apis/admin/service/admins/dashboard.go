package admins

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"time"
)

// Dashboard 工作台
func Dashboard(adminId int) (interface{}, error) {
	//设置时间
	nowTime := time.Now()
	yesterdayTime := time.Now().AddDate(0, 0, -1)
	nowStarTime := time.Date(nowTime.Year(), nowTime.Month(), nowTime.Day(), 0, 0, 0, 0, time.Local)
	yesterdayStarTime := time.Date(yesterdayTime.Year(), yesterdayTime.Month(), yesterdayTime.Day(), 0, 0, 0, 0, time.Local)

	// 活跃人数 今日人数｜昨日人数｜总人数
	var visitorToday, visitorYesterday, visitorSum int64
	if result := database.Db.Model(&models.Access{}).
		Select("COUNT(DISTINCT ip)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.AccessTypeDefault).
		Where("created_at BETWEEN ? AND ?", nowStarTime.Unix(), nowTime.Unix()).Count(&visitorToday); result.Error != nil {
		return nil, result.Error
	}

	if result := database.Db.Model(&models.Access{}).
		Select("COUNT(DISTINCT ip)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.AccessTypeDefault).
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime.Unix(), nowStarTime.Unix()).Count(&visitorYesterday); result.Error != nil {
		return nil, result.Error
	}

	if result := database.Db.Model(&models.Access{}).
		Select("COUNT(DISTINCT ip)").
		Where("type = ?", models.AccessTypeDefault).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Count(&visitorSum); result.Error != nil {
		return nil, result.Error
	}

	// 用户人数 今日人数｜昨日人数｜总人数
	var userToday, userYesterday, userSum int64
	if result := database.Db.Model(&models.User{}).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type > ?", models.UserTypeVirtual).
		Where("created_at BETWEEN ? AND ?", nowStarTime.Unix(), nowTime.Unix()).Count(&userToday); result.Error != nil {
		return nil, result.Error
	}

	if result := database.Db.Model(&models.User{}).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type > ?", models.UserTypeVirtual).
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime.Unix(), nowStarTime.Unix()).Count(&userYesterday); result.Error != nil {
		return nil, result.Error
	}

	if result := database.Db.Model(&models.User{}).
		Where("type > ?", models.UserTypeVirtual).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Count(&userSum); result.Error != nil {
		return nil, result.Error
	}

	//	今日充值｜昨日充值｜总充值
	var depositToday, depositYesterday, depositSum float64
	if result := database.Db.Model(&models.WalletOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.WalletOrderTypeDeposit).
		Where("status = ?", models.WalletOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", nowStarTime.Unix(), nowTime.Unix()).
		Find(&depositToday); result.Error != nil {
		return nil, result.Error
	}

	if result := database.Db.Model(&models.WalletOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.WalletOrderTypeDeposit).
		Where("status = ?", models.WalletOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime.Unix(), nowStarTime.Unix()).
		Find(&depositYesterday); result.Error != nil {
		return nil, result.Error
	}

	if result := database.Db.Model(&models.WalletOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.WalletOrderTypeDeposit).
		Where("status = ?", models.WalletOrderStatusComplete).
		Find(&depositSum); result.Error != nil {
		return nil, result.Error
	}

	// 今日提现｜昨日提现｜总提现
	var withdrawToday, withdrawYesterday, withdrawSum float64

	if result := database.Db.Model(&models.WalletOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.WalletOrderTypeWithdraw).
		Where("status = ?", models.WalletOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", nowStarTime.Unix(), nowTime.Unix()).
		Find(&withdrawToday); result.Error != nil {
		return nil, result.Error
	}

	database.Db.Model(&models.WalletOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.WalletOrderTypeWithdraw).
		Where("status = ?", models.WalletOrderStatusComplete).
		Where("created_at BETWEEN ? AND ?", yesterdayStarTime.Unix(), nowStarTime.Unix()).
		Find(&withdrawYesterday)

	if result := database.Db.Model(&models.WalletOrder{}).
		Select("IFNULL(SUM(money),0)").
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("type = ?", models.WalletOrderTypeWithdraw).
		Where("status = ?", models.WalletOrderStatusComplete).
		Find(&withdrawSum); result.Error != nil {
		return nil, result.Error
	}

	var category []string
	var visitorNumList []any
	var userNumList []any
	var depositNumList []any
	var withdrawNumList []any

	for i := -14; i <= 0; i++ {
		nowTimeTmp := time.Now().AddDate(0, 0, i)
		sourceTime := time.Date(nowTimeTmp.Year(), nowTimeTmp.Month(), nowTimeTmp.Day(), 0, 0, 0, 0, time.Local)
		staTime := sourceTime.Unix()
		category = append(category, sourceTime.Format("01/02"))

		// 访客数
		var visitorNum int64
		if result := database.Db.Model(&models.Access{}).Select("COUNT(DISTINCT ip)").
			Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
			Where("type = ?", models.AccessTypeDefault).
			Where("created_at BETWEEN ? AND ?", staTime, staTime+86399).Count(&visitorNum); result.Error != nil {
			return nil, result.Error
		}
		visitorNumList = append(visitorNumList, visitorNum)

		// 用户数
		var userNum int64
		if result := database.Db.Model(&models.User{}).
			Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
			Where("type > ?", models.UserTypeVirtual).
			Where("created_at BETWEEN ? AND ?", staTime, staTime+86399).Count(&userNum); result.Error != nil {
			return nil, result.Error
		}
		userNumList = append(userNumList, userNum)

		// 充值量
		var depositNum float64
		if result := database.Db.Model(&models.WalletOrder{}).
			Select("IFNULL(SUM(money),0)").
			Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
			Where("type = ?", models.WalletOrderTypeDeposit).
			Where("status = ?", models.WalletOrderStatusComplete).
			Where("created_at BETWEEN ? AND ?", staTime, staTime+86399).
			Find(&depositNum); result.Error != nil {
			return nil, result.Error
		}
		depositNumList = append(depositNumList, depositNum)

		// 提现量
		var withdrawNum float64
		if result := database.Db.Model(&models.WalletOrder{}).
			Select("IFNULL(SUM(money),0)").
			Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
			Where("type = ?", models.WalletOrderTypeWithdraw).
			Where("status = ?", models.WalletOrderStatusComplete).
			Where("created_at BETWEEN ? AND ?", staTime, staTime+86399).
			Find(&withdrawNum); result.Error != nil {
			return nil, result.Error
		}
		withdrawNumList = append(withdrawNumList, withdrawNum)
	}

	data := &dtoadmins.IndexPageData{
		Items: [][]*dtoadmins.Statistics{
			{
				&dtoadmins.Statistics{
					Name:      "访客数",
					Icon:      "sym_o_person_pin_circle",
					Color:     "bg-primary",
					Today:     visitorToday,
					Yesterday: visitorYesterday,
					Total:     visitorSum,
				},
				&dtoadmins.Statistics{
					Name:      "用户数",
					Icon:      "sym_o_person_add",
					Color:     "bg-secondary",
					Today:     userToday,
					Yesterday: userYesterday,
					Total:     userSum,
				},
				&dtoadmins.Statistics{
					Name:      "充值量",
					Icon:      "sym_o_credit_card",
					Color:     "bg-accent",
					Today:     depositToday,
					Yesterday: depositYesterday,
					Total:     depositSum,
				},
				&dtoadmins.Statistics{
					Name:      "提现量",
					Icon:      "sym_o_payments",
					Color:     "bg-dark",
					Today:     withdrawToday,
					Yesterday: withdrawYesterday,
					Total:     withdrawSum,
				},
			},
		},
		EchartsList: &dtoadmins.Echarts{
			Category: category,
			SeriesList: []*dtoadmins.Series{
				{
					Name:   "访客数",
					Type:   "line",
					Smooth: true,
					Data:   visitorNumList,
				}, {
					Name:   "用户数",
					Type:   "line",
					Smooth: true,
					Data:   userNumList,
				}, {
					Name:   "充值量",
					Type:   "line",
					Smooth: true,
					Data:   depositNumList,
				}, {
					Name:   "提现量",
					Type:   "line",
					Smooth: true,
					Data:   withdrawNumList,
				},
			},
		},
	}
	return data, nil
}
