package wallets

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/apis/common/service/users"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"errors"
	"gorm.io/gorm"
)

// DepositCreate 用户充值新增
func DepositCreate(adminId, userId int, params *dtowallets.DepositCreateParams) (interface{}, error) {
	// 判断用户是否存在
	userInfo := &models.User{}
	if result := database.Db.Model(&userInfo).
		Where("username = ?", params.UserName).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Take(&userInfo); result.Error != nil {
		return nil, result.Error
	}

	// 判断是否有充值方式
	walletPaymentInfo := &models.WalletPayment{}
	if result := database.Db.Model(&walletPaymentInfo).Where("id = ?", params.PaymentId).
		Where("mode in ?", []int{models.WalletPaymentModeDeposit, models.WalletPaymentModeAssetsDeposit}).
		Take(&walletPaymentInfo); result.Error != nil {
		return nil, result.Error
	}

	//	必须属于用户的充值方式
	if rds.GetAdminSettingId(userInfo.AdminId) != walletPaymentInfo.AdminId {
		return nil, errors.New("incorrectFormat")
	}

	// 插入订单数据
	walletOrderType := models.WalletOrderTypeDeposit
	if walletPaymentInfo.Mode == models.WalletPaymentModeAssetsDeposit {
		walletOrderType = models.WalletOrderTypeAssetsDeposit
	}

	if result := database.Db.Create(&models.WalletOrder{
		OrderSn:  utils.NewRandom().OrderSn(),
		AdminId:  userInfo.AdminId,
		UserId:   userInfo.Id,
		Type:     walletOrderType,
		SourceId: walletPaymentInfo.Id,
		AssetsId: walletPaymentInfo.AssetsId,
		Money:    params.Money,
		Status:   models.WalletOrderStatusActive,
	}); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// DepositDelete 用户充值删除
func DepositDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletOrder{}).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id IN ?", params.Ids).
		Where("type IN ?", []int{models.WalletOrderTypeDeposit, models.WalletOrderTypeAssetsDeposit}).
		Update("status", models.WalletOrderStatusDelete); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// DepositUpdate 用户充值更新处理
func DepositUpdate(adminId, userId int, params *dtowallets.DepositUpdateParams) (interface{}, error) {
	walletOrderInfo := &models.WalletOrder{}
	if result := database.Db.Model(walletOrderInfo).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id = ?", params.Id).Take(walletOrderInfo); result.Error != nil {
		return nil, result.Error
	}

	if result := database.Db.Model(&models.WalletOrder{}).
		Where("id = ?", walletOrderInfo.Id).
		Update("data", params.Data); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// DepositIndex 用户充值列表
func DepositIndex(adminId, userId int, params *dtowallets.DepositIndexParams) (interface{}, error) {
	data := dto.IndexData{Items: make([]*dtowallets.DepositIndexData, 0)}

	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE  ?", params.AdminName+"%")).
		InInt("user_id IN ?", models.FindTableColumnIntn("user", "id", "username LIKE  ?", params.UserName+"%")).
		EqInt("source_id = ?", params.SourceId).
		EqInt("status = ?", params.Status).
		EqInt("type = ?", params.Type).
		Like("data LIKE ?", params.Data+"%").
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	if result := database.Db.Table("wallet_order AS wo").
		Select("wo.id", "wo.admin_id", "au.username AS adminName", "wo.user_id", "u.username", "wo.source_id", "wo.type", "wo.order_sn", "wo.money", "wo.data", "wo.voucher", "wo.fee", "wo.status", "wo.data", "wo.updated_at", "wo.created_at").
		Joins("left join admin_user AS au on wo.admin_id = au.id").
		Joins("left join user AS u on u.id = wo.user_id").
		Where("wo.status > ?", models.WalletAssetsStatusDelete).
		Where("wo.type IN ?", []int{models.WalletOrderTypeAssetsDeposit, models.WalletOrderTypeDeposit}).
		Where("wo.admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}

// DepositStatus 充值审核
func DepositStatus(adminId, userId int, params *dtowallets.DepositStatusParams) (interface{}, error) {
	walletOrderInfo := &models.WalletOrder{}

	if result := database.Db.Model(&walletOrderInfo).
		Where("id = ?", params.Id).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("status = ?", models.WalletOrderStatusActive).
		Take(&walletOrderInfo); result.Error != nil {
		return nil, result.Error
	}

	// 是否存在用户
	userInfo := &models.User{}
	if result := database.Db.Model(&userInfo).
		Where("id = ?", walletOrderInfo.UserId).
		Take(&userInfo); result.Error != nil {
		return nil, result.Error
	}

	//	事务开始
	err := database.Db.Transaction(func(tx *gorm.DB) error {
		//	更新充值订单状态
		if result := tx.Model(models.WalletOrder{}).
			Where("id = ?", walletOrderInfo.Id).
			Updates(&params); result.Error != nil {
			return result.Error
		}

		// 如果状态是完成, 那么用户余额｜资产增加
		if params.Status == models.WalletOrderStatusComplete {
			//	如果是余额充值
			if walletOrderInfo.Type == models.WalletOrderTypeDeposit {
				if err := users.UserDeposit(tx, &dto.WalletOrderAgreeParams{
					AdminId:  adminId,
					ParentId: userInfo.ParentId,
					UserId:   userInfo.Id,
					SourceId: walletOrderInfo.Id,
					BillType: models.WalletBillTypeDeposit,
					Balance:  userInfo.Money,
					Money:    walletOrderInfo.Money,
				}); err != nil {
					return err
				}
			}

			//	如果是用户资产
			if walletOrderInfo.Type == models.WalletOrderTypeAssetsDeposit {
				if err := users.UserAssetsDeposit(tx, &dto.WalletOrderAgreeParams{
					AdminId:  adminId,
					ParentId: userInfo.ParentId,
					UserId:   userInfo.Id,
					AssetId:  walletOrderInfo.AssetsId,
					SourceId: walletOrderInfo.Id,
					BillType: models.WalletOrderTypeAssetsDeposit,
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
