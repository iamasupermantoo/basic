package admins

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/common/rds"
	"basic/models"
	"basic/module/database"

	"gorm.io/gorm"
)

// MerchantReset  重置商户配置数据
func MerchantReset(params *dtoadmins.MerchantParams) (interface{}, error) {
	resetTableList := []map[string]string{
		//	跟管理相关表
		{"name": "access", "field": "", "opt": "DELETE"},
		{"name": "admin_logs", "field": "", "opt": "DELETE"},
		{"name": "invite", "field": "", "opt": "DELETE"},
		{"name": "level_order", "field": "", "opt": "DELETE"},
		{"name": "wallet_bill", "field": "", "opt": "DELETE"},
		{"name": "wallet_order", "field": "", "opt": "DELETE"},
		{"name": "wallet_user_account", "field": "", "opt": "DELETE"},
		{"name": "wallet_user_assets", "field": "", "opt": "DELETE"},
		{"name": "notify", "field": "", "opt": "DELETE"},
		{"name": "user", "field": "", "opt": "DELETE"},
		{"name": "real_name_auth", "field": "", "opt": "DELETE"},
		{"name": "setting", "field": "", "opt": "DELETE"},

		//	初始化商户相关数据表
		{"name": "admin_setting", "field": "", "opt": "RESET"},
		{"name": "article", "field": "", "opt": "RESET"},
		{"name": "country", "field": "", "opt": "RESET"},
		{"name": "lang", "field": "", "opt": "RESET"},
		{"name": "level", "field": "", "opt": "RESET"},
		{"name": "translate", "field": "", "opt": "RESET"},
		{"name": "wallet_assets", "field": "", "opt": "RESET"},
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		for _, v := range resetTableList {
			err := merchantSyncFunc(tx, params.Id, v["name"], v["field"], v["opt"])
			if err != nil {
				return err
			}
		}

		//	初始化支付表
		database.Db.Where("admin_id = ?", params.Id).Delete(&models.WalletPayment{})
		err := merchantSyncPaymentFunc(tx, models.SuperAdminId, params.Id, "RESET")

		//	初始化菜单
		database.Db.Where("admin_id = ?", params.Id).Delete(&models.AdminMenu{})
		err = merchantSyncMenuFunc(tx, models.SuperAdminId, params.Id, 0, 0, "RESET")
		return err
	})

	//	删除管理所有缓存
	rds.InitAdminRedis(params.Id)
	return "ok", err
}

// MerchantSync 同步商户数据：配置等级数据 管理设置 文章 国家 语言 等级 翻译 钱包 支付信息
func MerchantSync(params *dtoadmins.MerchantParams) (interface{}, error) {
	//	初始化商户相关数据表
	resetTableList := []map[string]string{
		{"name": "article", "field": "name", "opt": "SYNC"},
		{"name": "admin_setting", "field": "field", "opt": "SYNC"},
		{"name": "country", "field": "alias", "opt": "SYNC"},
		{"name": "lang", "field": "alias", "opt": "SYNC"},
		{"name": "level", "field": "name", "opt": "SYNC"},
		{"name": "translate", "field": "field", "opt": "SYNC"},
		{"name": "wallet_assets", "field": "name", "opt": "SYNC"},
	}

	err := database.Db.Transaction(func(tx *gorm.DB) error {
		for _, v := range resetTableList {
			err := merchantSyncFunc(tx, params.Id, v["name"], v["field"], v["opt"])
			if err != nil {
				return err
			}
		}

		//	初始化支付表
		err := merchantSyncPaymentFunc(tx, models.SuperAdminId, params.Id, "SYNC")

		//	初始化菜单
		err = merchantSyncMenuFunc(tx, models.SuperAdminId, params.Id, 0, 0, "SYNC")
		return err
	})

	//	删除管理所有缓存
	rds.InitAdminRedis(params.Id)
	return "ok", err
}

// merchantSyncFunc 管理同步方法
func merchantSyncFunc(tx *gorm.DB, adminId int, tableName string, field string, opt string) error {
	//	删除当前管理数据
	if opt != "SYNC" {
		if result := tx.Exec("DELETE FROM "+tableName+" WHERE admin_id = ?", adminId); result.Error != nil {
			return result.Error
		}
	}
	if opt == "DELETE" {
		return nil
	}

	data := []map[string]interface{}{}
	if result := database.Db.Table(tableName).Where("admin_id = ?", models.SuperAdminId).Find(&data); result.Error != nil {
		return result.Error
	}

	//	插入管理配置
	for _, v := range data {
		v["id"] = 0
		v["admin_id"] = adminId

		syncInsert := false
		if opt == "SYNC" {
			currentData := map[string]interface{}{}
			result := database.Db.Table(tableName).Where(field+" = ?", v[field]).Where("admin_id = ?", adminId).Find(&currentData)
			if result.Error != nil {
				return result.Error
			}
			syncInsert = result.RowsAffected == 0
		}

		if opt == "RESET" || (opt == "SYNC" && syncInsert) {
			if result := tx.Table(tableName).Create(v); result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

// merchantSyncPaymentFunc 同步支付表
func merchantSyncPaymentFunc(tx *gorm.DB, superId, initId int, opt string) error {
	if opt != "SYNC" {
		database.Db.Where("admin_id = ?", initId).Delete(&models.WalletPayment{})
	}

	paymentList := make([]*models.WalletPayment, 0)
	if result := database.Db.Model(&models.WalletPayment{}).Where("admin_id = ?", superId).Find(&paymentList); result.Error != nil {
		return result.Error
	}

	for _, v := range paymentList {
		v.Id = 0
		v.AdminId = initId

		syncInsert := false
		if opt == "SYNC" {
			initPaymentInfo := &models.WalletPayment{}
			result := database.Db.Model(initPaymentInfo).Where("admin_id = ?", initId).Where("name = ?", v.Name).Where("mode = ?", v.Mode).Where("type = ?", v.Type).Find(initPaymentInfo)
			syncInsert = result.RowsAffected == 0
		}

		if v.AssetsId != 0 {
			superAssetsInfo := &models.WalletAssets{}
			if result := database.Db.Model(superAssetsInfo).Where("id = ?", v.AssetsId).Take(superAssetsInfo); result.Error != nil {
				return result.Error
			}

			initAssetInfo := &models.WalletAssets{}
			if result := tx.Model(initAssetInfo).Where("name = ?", superAssetsInfo.Name).Where("admin_id = ?", initId).Take(initAssetInfo); result.Error != nil {
				return result.Error
			}
			v.AssetsId = initAssetInfo.Id
		}

		if opt != "SYNC" || (opt == "SYNC" && syncInsert) {
			if result := tx.Model(&models.WalletPayment{}).Create(v); result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

// merchantSyncMenuFunc 商户同步菜单
func merchantSyncMenuFunc(tx *gorm.DB, superId, initId, parentId, initParentId int, opt string) error {
	menuList := make([]*models.AdminMenu, 0)
	if result := database.Db.Model(&models.AdminMenu{}).
		Where("parent_id = ?", parentId).
		Where("admin_id = ?", superId).
		Where("type <> ?", models.AdminMenuTypeAdminMenu).
		Find(&menuList); result.Error != nil {
		return result.Error
	}

	for _, menu := range menuList {
		// 插入对应的管理员
		parentId = menu.Id
		menu.Id = 0
		menu.AdminId = initId
		menu.ParentId = initParentId

		syncInsert := false
		if opt == "SYNC" {
			menuInfo := &models.AdminMenu{}
			result := database.Db.Model(menuInfo).Where("route = ?", menu.Route).Where("admin_id = ?", initId).Find(menuInfo)
			if result.Error != nil {
				return result.Error
			}
			syncInsert = result.RowsAffected == 0
		}

		if opt == "RESET" || (opt == "SYNC" && syncInsert) {
			if result := tx.Model(&models.AdminMenu{}).Create(menu); result.Error != nil {
				return result.Error
			}
		}

		// 递归插入下级数据
		if err := merchantSyncMenuFunc(tx, superId, initId, parentId, menu.Id, opt); err != nil {
			return err
		}
	}
	return nil
}
