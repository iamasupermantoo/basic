package users

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils/views"
	"errors"
	"strconv"
	"time"

	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

// UserDeposit 用户资金入账
func UserDeposit(tx *gorm.DB, params *dto.WalletOrderAgreeParams) error {
	nowTime := time.Now()

	//	检查类型是否存在
	if _, ok := models.WalletBillDepositTypeList[params.BillType]; !ok {
		return errors.New("incorrectFormat")
	}

	//	更新用户金额
	result := tx.Model(&models.User{}).Where("id = ?", params.UserId).Updates(map[string]interface{}{"money": params.Balance + params.Money, "updated_at": nowTime.Unix()})
	if result.Error != nil {
		return result.Error
	}

	//	团队收益操作
	err := TeamEarnings(tx, params.AdminId, params.ParentId, params.SourceId, params.AssetId, models.WalletBillSpendTypeList[params.BillType], params.Money)
	if err != nil {
		return err
	}

	//	写入账单
	walletBill := &models.WalletBill{AdminId: params.AdminId, UserId: params.UserId, AssetsId: 0, SourceId: params.SourceId, Type: params.BillType, Name: models.WalletBillDepositTypeList[params.BillType], Money: params.Money, Balance: params.Balance, CreatedAt: int(nowTime.Unix())}
	result = tx.Create(&walletBill)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UserAssetsDeposit 用户资产入金
func UserAssetsDeposit(tx *gorm.DB, params *dto.WalletOrderAgreeParams) error {
	nowTime := time.Now()

	if _, ok := models.WalletBillDepositTypeList[params.BillType]; !ok {
		return errors.New("incorrectFormat")
	}

	userAssetsInfo := &models.WalletUserAssets{}
	result := database.Db.Model(userAssetsInfo).Where("user_id = ?", params.UserId).Where("wallet_assets_id = ?", params.AssetId).Where("status > ?", models.WalletUserAssetsStatusDelete).Find(userAssetsInfo)
	if result.Error != nil {
		return result.Error
	}

	//	如果查不到数据, 那么创建用户资产数据
	var err error
	if result.RowsAffected == 0 {
		params.Balance = 0
		userAssetsInfo = &models.WalletUserAssets{AdminId: params.AdminId, UserId: params.UserId, WalletAssetsId: params.AssetId, Money: params.Money, UpdatedAt: int(nowTime.Unix()), CreatedAt: int(nowTime.Unix())}
		result = tx.Create(userAssetsInfo)
		err = result.Error
	} else {
		params.Balance = userAssetsInfo.Money
		result = tx.Model(userAssetsInfo).Where("id = ?", userAssetsInfo.Id).Updates(map[string]interface{}{"money": params.Balance + params.Money, "updated_at": nowTime.Unix()})
		err = result.Error
	}
	if err != nil {
		return err
	}

	//	团队收益操作
	err = TeamEarnings(tx, params.AdminId, params.ParentId, params.SourceId, params.AssetId, models.WalletBillSpendTypeList[params.BillType], params.Money)
	if err != nil {
		return err
	}

	//	写入账单
	walletBill := &models.WalletBill{AdminId: params.AdminId, UserId: params.UserId, AssetsId: userAssetsInfo.WalletAssetsId, SourceId: params.SourceId, Type: params.BillType, Name: models.WalletBillDepositTypeList[params.BillType], Money: params.Money, Balance: params.Balance, CreatedAt: int(nowTime.Unix())}
	result = tx.Create(walletBill)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UserSpending 用户资金花费
func UserSpending(tx *gorm.DB, params *dto.WalletOrderAgreeParams) error {
	nowTime := time.Now()
	//	检查类型是否存在
	if _, ok := models.WalletBillSpendTypeList[params.BillType]; !ok {
		return errors.New("incorrectFormat")
	}

	//	检查余额不足
	if params.Money <= 0 || params.Money > params.Balance {
		return errors.New("insufficientBalance")
	}

	//	团队收益操作
	err := TeamEarnings(tx, params.AdminId, params.ParentId, params.SourceId, params.AssetId, models.WalletBillSpendTypeList[params.BillType], params.Money)
	if err != nil {
		return err
	}

	//	更新用户金额
	result := tx.Model(&models.User{}).Where("id = ?", params.UserId).Updates(map[string]interface{}{"money": params.Balance - params.Money, "updated_at": nowTime.Unix()})
	if result.Error != nil {
		return result.Error
	}

	//	写入账单
	walletBill := &models.WalletBill{AdminId: params.AdminId, UserId: params.UserId, AssetsId: 0, SourceId: params.SourceId, Type: params.BillType, Name: models.WalletBillSpendTypeList[params.BillType], Money: params.Money, Balance: params.Balance, CreatedAt: int(nowTime.Unix())}
	result = tx.Create(&walletBill)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// UserAssetsSpending 用户资产出金
func UserAssetsSpending(tx *gorm.DB, params *dto.WalletOrderAgreeParams) error {
	nowTime := time.Now()
	//	检查类型是否存在
	if _, ok := models.WalletBillSpendTypeList[params.BillType]; !ok {
		return errors.New("incorrectFormat")
	}

	userAssetsInfo := &models.WalletUserAssets{}
	result := database.Db.Model(userAssetsInfo).Where("user_id = ?", params.UserId).Where("wallet_assets_id = ?", params.AssetId).Where("status = ?", models.WalletUserAssetsStatusActive).Take(userAssetsInfo)
	if result.Error != nil {
		return result.Error
	}

	//	检查余额不足
	if params.Money <= 0 || params.Money > userAssetsInfo.Money {
		return errors.New("insufficientBalance")
	}

	//	团队收益操作
	err := TeamEarnings(tx, params.AdminId, params.ParentId, params.SourceId, params.AssetId, models.WalletBillSpendTypeList[params.BillType], params.Money)
	if err != nil {
		return err
	}

	//	更新用户资产金额
	params.Balance = userAssetsInfo.Money
	result = tx.Model(&models.WalletUserAssets{}).Where("id = ?", userAssetsInfo.Id).Updates(map[string]interface{}{"money": params.Balance - params.Money, "updated_at": nowTime.Unix()})
	if result.Error != nil {
		return result.Error
	}

	//	写入账单
	walletBill := &models.WalletBill{AdminId: params.AdminId, UserId: params.UserId, AssetsId: userAssetsInfo.WalletAssetsId, SourceId: params.SourceId, Type: params.BillType, Name: models.WalletBillSpendTypeList[params.BillType], Money: params.Money, Balance: params.Balance, CreatedAt: int(nowTime.Unix())}
	result = tx.Create(&walletBill)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// TeamEarnings 团队邀请收益
func TeamEarnings(tx *gorm.DB, adminId, parentId, sourceId, assetsId int, billTypeStr string, money float64) error {
	//	过滤不需要的查询
	if parentId > 0 {
		rdsConn := cache.Rds.Get()
		defer rdsConn.Close()
		teamEarningsSetting := make([]*dtoadmins.AdminSettingTeamEarningSetting, 0)
		_ = json.Unmarshal([]byte(rds.RedisFindAdminSettingField(rdsConn, adminId, "earningsSetting")), &teamEarningsSetting)

		for i := 0; i < len(teamEarningsSetting); i++ {
			//	并且设置有当前账单类型 并且 上级ID 大于 0
			if _, ok := teamEarningsSetting[i].Options[billTypeStr]; ok && parentId > 0 {

				//	值必须为真 , 收益率必须大于 0
				if teamEarningsSetting[i].Options[billTypeStr] && teamEarningsSetting[i].Rate > 0 {

					//	如果资产ID设置了,  那么增加用户资产
					currentUserInfo := &models.User{}
					result := database.Db.Model(currentUserInfo).Where("id = ?", parentId).Take(currentUserInfo)
					if result.Error != nil {
						return result.Error
					}

					//	当前等级收益金额
					earningsMoney := money * teamEarningsSetting[i].Rate / 100
					if assetsId > 0 {
						err := UserAssetsDeposit(tx, &dto.WalletOrderAgreeParams{
							AssetId:  assetsId,
							AdminId:  currentUserInfo.AdminId,
							ParentId: currentUserInfo.ParentId,
							UserId:   currentUserInfo.Id,
							SourceId: sourceId,
							BillType: models.WalletBillTypeTeamAssetsEarnings,
							Money:    earningsMoney,
						})
						if err != nil {
							return err
						}
					} else {
						err := UserDeposit(tx, &dto.WalletOrderAgreeParams{
							AdminId:  currentUserInfo.AdminId,
							ParentId: currentUserInfo.ParentId,
							UserId:   currentUserInfo.Id,
							SourceId: sourceId,
							BillType: models.WalletBillTypeTeamEarnings,
							Balance:  currentUserInfo.Money,
							Money:    earningsMoney,
						})
						if err != nil {
							return err
						}
					}

					//	下一级用户ID
					parentId = currentUserInfo.ParentId
				}
			}
		}
	}
	return nil
}

// GetLevelOptions 获取等级Options
func GetLevelOptions(adminId int) []*views.InputOptions {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	data := make([]*views.InputOptions, 0)
	var levelList []*models.Level
	database.Db.Model(&models.Level{}).Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Where("status = ?", models.LevelStatusActive).Find(&levelList)

	for i := 0; i < len(levelList); i++ {
		data = append(data, &views.InputOptions{
			Label: rds.RedisFindTranslateField(rdsConn, adminId, "zh-CN", levelList[i].Name) + "【 管理ID:" + strconv.Itoa(levelList[i].AdminId) + " 】",
			Value: levelList[i].Id,
		})
	}
	return data
}

// GetUserChildren 获取用户下级Ids 组
func GetUserChildren(userId int) []int {
	userIds := []int{userId}

	ids := make([]int, 0)
	database.Db.Model(&models.User{}).Select("id").Where("parent_id = ?", userId).Where("status > ?", models.UserStatusDelete).Find(&ids)
	for _, id := range ids {
		userIds = append(userIds, GetUserChildren(id)...)
	}
	return userIds
}

// GetUserDepth 获取用户深度
func GetUserDepth(userId, index int, depth int) int {
	userInfo := &models.User{}
	database.Db.Model(userInfo).Where("id = ?", userId).Find(userInfo)
	if userInfo.Id == 0 || userInfo.ParentId == 0 {
		return depth
	}
	depth++
	if userInfo.ParentId == index {
		return depth
	}
	return GetUserDepth(userInfo.ParentId, index, depth)
}
