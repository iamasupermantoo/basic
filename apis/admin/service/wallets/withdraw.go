package wallets

import (
	dtoadmins "basic/apis/admin/dto/admins"
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
	"basic/utils"
	"errors"
	"time"

	"github.com/goccy/go-json"
	"gorm.io/gorm"
)

// WithdrawIndex 用户提现列表
func WithdrawIndex(adminId, userId int, params *dtowallets.WithdrawIndexParams) (interface{}, error) {
	data := &dto.IndexData{Items: make([]*dtowallets.WithdrawIndexData, 0)}
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE ?", params.UserName+"%")).
		InInt("source_id IN ?", models.FindTableColumnIntn("wallet_user_account", "id", "real_name LIKE ?", params.RealName+"%")).
		InInt("source_id IN ?", models.FindTableColumnIntn("wallet_user_account", "id", "number LIKE ?", params.Number+"%")).
		InInt("source_id IN ?", models.FindTableColumnIntn("wallet_user_account", "id", "code LIKE ?", params.Code+"%")).
		Like("order_sn LIKE ?", params.OrderSn+"%").
		Like("data LIKE ?", params.Data).
		EqInt("status = ?", params.Status).
		EqInt("type = ?", params.Type).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	result := database.Db.Table("wallet_order AS wo").
		Select("wo.id", "wo.admin_id", "au.username AS adminName", "wua.real_name AS realName", "wua.number  AS number", "wua.code", "u.username", "wo.type", "wo.order_sn", "wo.money", "wo.data", "wo.fee", "wo.status", "wo.data", "wo.updated_at", "wo.created_at").
		Joins("left join admin_user AS au on wo.admin_id = au.id").
		Joins("left join user AS u on u.id = wo.user_id").
		Joins("left join wallet_user_account AS wua on wua.id = wo.source_id").
		Where("wo.status > ?", models.WalletAssetsStatusDelete).
		Where("wo.type IN ?", []int{models.WalletOrderTypeWithdraw, models.WalletOrderTypeAssetsWithdraw}).
		Where("wo.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items)

	return data, result.Error
}

// WithdrawCreate 用户提现创建
func WithdrawCreate(adminId, userId int, params *dtowallets.WithdrawCreateParams) (interface{}, error) {
	rdsConn := cache.Rds.Get()
	defer rdsConn.Close()

	// 查询用户是否存在
	userInfo := &models.User{}
	if result := database.Db.Model(userInfo).
		Where("username = ?", params.UserName).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Take(userInfo); result.Error != nil {
		return nil, result.Error
	}

	//	支付信息
	paymentInfo := &models.WalletPayment{}
	if result := database.Db.Model(paymentInfo).
		Where("id = ?", params.PaymentId).
		Where("mode in ?", []int{models.WalletPaymentModeWithdraw, models.WalletPaymentModeAssetsWithdraw}).
		Take(paymentInfo); result.Error != nil {
		return nil, result.Error
	}

	//	检查支付方式是否匹配
	if paymentInfo.AssetsId > 0 && rds.GetAdminSettingId(userInfo.AdminId) != paymentInfo.AdminId {
		return nil, errors.New("支付类型不匹配")
	}

	//	对应支付信息的卡片
	accountInfo := &models.WalletUserAccount{}
	if result := database.Db.Model(accountInfo).
		Where("payment_id = ?", paymentInfo.Id).
		Where("user_id = ?", userInfo.Id).
		Where("status = ?", models.WalletUserAccountStatusActive).
		Take(accountInfo); result.Error != nil {
		return nil, errors.New("提现账号未绑定")
	}

	//	获取管理提现配置信息
	nowTime := time.Now()
	withdrawSetting := &dtoadmins.AdminSettingWithdrawSetting{}
	_ = json.Unmarshal([]byte(rds.RedisFindAdminSettingField(rdsConn, adminId, "walletWithdrawSetting")), &withdrawSetting)

	walletOrderType := models.WalletOrderTypeWithdraw
	if paymentInfo.Mode == models.WalletOrderTypeAssetsWithdraw {
		walletOrderType = models.WalletOrderTypeAssetsWithdraw
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		walletOrderInfo := &models.WalletOrder{
			OrderSn:   utils.NewRandom().OrderSn(),
			AdminId:   userInfo.AdminId,
			UserId:    userInfo.Id,
			Type:      walletOrderType,
			SourceId:  accountInfo.Id,
			AssetsId:  paymentInfo.AssetsId,
			Money:     params.Money,
			Fee:       withdrawSetting.Fee * params.Money / 100,
			UpdatedAt: int(nowTime.Unix()),
			CreatedAt: int(nowTime.Unix()),
		}

		if result := tx.Create(walletOrderInfo); result.Error != nil {
			return result.Error
		}

		switch paymentInfo.Mode {
		case models.WalletPaymentModeWithdraw:
			//	余额提现
			if params.Money > userInfo.Money {
				return errors.New("insufficientBalance")
			}

			if err := users.UserSpending(tx, &dto.WalletOrderAgreeParams{
				AdminId:  userInfo.AdminId,
				UserId:   userInfo.Id,
				ParentId: userInfo.ParentId,
				SourceId: walletOrderInfo.Id,
				BillType: models.WalletBillTypeWithdraw,
				Balance:  userInfo.Money,
				Money:    walletOrderInfo.Money,
			}); err != nil {
				return err
			}

		case models.WalletPaymentModeAssetsWithdraw:
			if err := users.UserAssetsSpending(tx, &dto.WalletOrderAgreeParams{
				AssetId:  paymentInfo.AssetsId,
				AdminId:  userInfo.AdminId,
				UserId:   userInfo.Id,
				ParentId: userInfo.ParentId,
				SourceId: walletOrderInfo.Id,
				BillType: models.WalletBillTypeAssetsWithdraw,
				Money:    walletOrderInfo.Money,
			}); err != nil {
				return err
			}
		}

		return nil
	})

	return "ok", err
}

// WithdrawDelete 用户提现删除
func WithdrawDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	result := database.Db.Model(models.WalletOrder{}).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id IN ?", params.Ids).
		Where("status = ?", models.WalletOrderStatusComplete).
		Update("status", models.WalletOrderStatusDelete)

	return "ok", result.Error
}

// WithdrawUpdate 用户提现更新
func WithdrawUpdate(adminId, userId int, params *dtowallets.WithdrawUpdateParams) (interface{}, error) {
	result := database.Db.Model(models.WalletOrder{}).
		Where("admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id = ? ", params.Id).
		Updates(params)

	return "ok", result.Error
}

// WithdrawStatus 用户提现审核
func WithdrawStatus(adminId, userId int, params *dtowallets.WithdrawStatusParams) (interface{}, error) {
	walletOrderInfo := &models.WalletOrder{}
	if result := database.Db.Model(models.WalletOrder{}).
		Where("id = ?", params.Id).
		Where("admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Take(walletOrderInfo); result.Error != nil {
		return nil, result.Error
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Model(models.WalletOrder{}).
			Where("id = ?", params.Id).
			Updates(params); result.Error != nil {
			return result.Error
		}

		userInfo := &models.User{}
		if result := database.Db.Model(userInfo).Where("id=?", walletOrderInfo.UserId).Take(userInfo); result.Error != nil {
			return result.Error
		}

		// 如果拒绝，那么新增用户金额
		if params.Status == models.WalletOrderStatusRefuse {
			switch walletOrderInfo.Type {
			case models.WalletOrderTypeWithdraw:
				//	拒绝提现, 增加用户账户资金
				if err := users.UserDeposit(tx, &dto.WalletOrderAgreeParams{
					AdminId:  userInfo.AdminId,
					ParentId: userInfo.ParentId,
					UserId:   userInfo.Id,
					SourceId: walletOrderInfo.Id,
					BillType: models.WalletBillTypeWithdrawRefuse,
					Balance:  userInfo.Money,
					Money:    walletOrderInfo.Money,
				}); err != nil {
					return err
				}

			case models.WalletOrderTypeAssetsWithdraw:
				assetsInfo := &models.WalletAssets{}
				if result := database.Db.Model(assetsInfo).Where("id = ?", walletOrderInfo.AssetsId).Take(assetsInfo); result.Error != nil {
					return result.Error
				}

				//	拒绝提现 - 返回用户资产数量
				if err := users.UserAssetsDeposit(tx, &dto.WalletOrderAgreeParams{
					AssetId:  assetsInfo.Id,
					AdminId:  userInfo.AdminId,
					ParentId: userInfo.ParentId,
					UserId:   userInfo.Id,
					SourceId: walletOrderInfo.Id,
					BillType: models.WalletBillTypeAssetsWithdrawRefuse,
					Money:    walletOrderInfo.Money,
				}); err != nil {
					return err
				}
			}
		}

		return nil
	})

	return "ok", err
}
