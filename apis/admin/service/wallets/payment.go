package wallets

import (
	dtowallets "basic/apis/admin/dto/wallets"
	"basic/apis/common/dto"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"
	"basic/utils"
	"errors"
	"github.com/goccy/go-json"
)

// PaymentIndex 支付列表
func PaymentIndex(adminId, userId int, params *dtowallets.PaymentIndexParams) (interface{}, error) {
	var data = dto.IndexData{Items: make([]*dtowallets.PaymentIndexData, 0)}
	filterParams := utils.NewFilterParams().
		InInt("admin_id IN ?", models.FindTableColumnIntn("admin_user", "id", "username LIKE  ?", params.AdminName+"%")).
		Like("name LIKE ?", params.Name+"%").
		Like("data LIKE ?", params.Data+"%").
		Like("desc LIKE ?", params.Desc+"%").
		EqInt("status = ?", params.Status).
		EqInt("assets_id = ?", params.AssetsId).
		EqInt("mode = ?", params.Mode).
		EqInt("type = ?", params.Type).
		BetweenDate("created_at BETWEEN ? AND ?", params.CreatedAt)

	adminChildrenIds := rds.RedisFindAdminChildrenIds(adminId)
	if result := database.Db.Table("wallet_payment AS wp").
		Select("wp.id", "wp.admin_id", "au.username AS adminName", "wp.assets_id", "wp.name", "wp.icon", "wp.type", "wp.mode", "wp.level", "wp.status", "wp.data", "wp.desc", "wp.updated_at", "wp.created_at").
		Joins("LEFT JOIN admin_user AS au on wp.admin_id = au.id").
		Joins("LEFT JOIN wallet_assets AS wa on wp.assets_id = wa.id").
		Where("wp.status > ?", models.WalletPaymentStatusDelete).
		Where("wp.admin_id IN ?", adminChildrenIds).
		Scopes(filterParams.Scopes()).
		Count(&data.Count).
		Scopes(utils.Paginate(params.Pagination)).
		Scan(&data.Items); result.Error != nil {
		return nil, result.Error
	}

	//	处理资产ID 如果管理Id等于0, 那么使用 -1 提现类型
	for _, v := range data.Items.([]*dtowallets.PaymentIndexData) {
		if v.AdminId == 0 {
			v.AssetsId = -1
		}
		// 如果是管理员设置, 并且是提现, 那么转化数据
		if v.AdminId > 0 && (v.Mode == models.WalletPaymentModeDeposit || v.Mode == models.WalletPaymentModeAssetsDeposit) {
			v.DataJson = &models.WalletPaymentData{}
			_ = json.Unmarshal([]byte(v.Data), &v.DataJson)
		}
	}

	return data, nil
}

// PaymentCreate 支付创建
func PaymentCreate(adminId, userId int, params *dtowallets.PaymentCreateParams) (interface{}, error) {
	if params.AssetsId > 0 {
		assetsInfo := &models.WalletAssets{}
		if result := database.Db.Model(assetsInfo).Where("id = ?", params.AssetsId).Take(assetsInfo); result.Error != nil {
			return nil, result.Error
		}
		if adminId != assetsInfo.AdminId {
			return nil, errors.New("incorrectFormat")
		}
	}

	if result := database.Db.Create(&models.WalletPayment{
		Icon:     params.Icon,
		AdminId:  adminId,
		Name:     params.Name,
		Mode:     params.Mode,
		Type:     params.Type,
		AssetsId: params.AssetsId,
	}); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// PaymentDelete 支付删除
func PaymentDelete(adminId, userId int, params *dto.DeleteParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletPayment{}).
		Where("admin_id in ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id in ?", params.Ids).
		Updates(map[string]interface{}{"status": models.WalletPaymentStatusDelete}); result.Error != nil {
		return nil, result.Error
	}

	return "ok", nil
}

// PaymentUpdate 支付修改
func PaymentUpdate(adminId, userId int, params *dtowallets.PaymentUpdateParams) (interface{}, error) {
	if result := database.Db.Model(&models.WalletPayment{}).
		Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).
		Where("id = ?", params.Id).
		Updates(&params); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}

// PaymentUpdateDesc 更新支付数据
func PaymentUpdateDesc(adminId, userId int, params *dtowallets.PaymentDescParams) (interface{}, error) {
	dataJsonBytes, _ := json.Marshal(params.DataJson)
	if result := database.Db.Model(&models.WalletPayment{}).Where("id=?", params.Id).Where("admin_id IN ?", rds.RedisFindAdminChildrenIds(adminId)).Update("data", string(dataJsonBytes)); result.Error != nil {
		return nil, result.Error
	}
	return "ok", nil
}
