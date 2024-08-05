package database

import (
	dtoadmins "basic/apis/admin/dto/admins"
	"basic/apis/admin/service/admins"
	"basic/models"
	"basic/module/cache"
	"basic/module/database"
)

// TableInfo 创建表参数
type TableInfo struct {
	TableModel interface{} // 表模型
	Data       interface{} // 插入数据
	Command    string      // 修改命令
}

// InitDatabase 初始化数据库
func InitDatabase() {
	//	清理Redis缓存数据
	rds := cache.Rds.Get()
	defer rds.Close()
	_, _ = rds.Do("flushall")

	// 创建表
	createTable([]*TableInfo{
		{TableModel: &models.AdminSetting{}, Data: AdminSettingInit(), Command: "COMMENT '管理设置表';"},
		{TableModel: &models.Article{}, Data: SystemArticleInit(), Command: "COMMENT '文章表';"},
		{TableModel: &models.Notify{}, Data: nil, Command: "COMMENT '系统通知表';"},
		{TableModel: &models.AdminMenu{}, Data: AdminMenuInit(), Command: "COMMENT '管理菜单表';"},
		{TableModel: &models.WalletPayment{}, Data: WalletPaymentInit(), Command: "COMMENT '钱包支付表';"},
		{TableModel: &models.WalletAssets{}, Data: WalletAssetsInit(), Command: "COMMENT '钱包资产管理表';"},
		{TableModel: &models.Country{}, Data: SystemCountryInit(), Command: "COMMENT '系统国家表';"},
		{TableModel: &models.Level{}, Data: SystemLevelInit(), Command: "COMMENT '系统等级表';"},
		{TableModel: &models.Lang{}, Data: SystemLangInit(), Command: "COMMENT '系统语言表';"},
		{TableModel: &models.Translate{}, Data: SystemTranslateInit(), Command: "COMMENT '系统语言字典表';"},
		{TableModel: &models.WalletUserAccount{}, Data: nil, Command: "COMMENT '钱包卡片表';"},
		{TableModel: &models.LevelOrder{}, Data: nil, Command: "COMMENT '用户等级订单表';"},
		{TableModel: &models.WalletOrder{}, Data: nil, Command: "COMMENT '钱包订单表';"},
		{TableModel: &models.WalletBill{}, Data: nil, Command: "COMMENT '钱包账单表';"},
		{TableModel: &models.WalletUserAssets{}, Data: nil, Command: "COMMENT '钱包用户资产表';"},
		{TableModel: &models.Invite{}, Data: nil, Command: "COMMENT '邀请表';"},
		{TableModel: &models.RealNameAuth{}, Data: nil, Command: "COMMENT '用户实名认证表';"},
		{TableModel: &models.Setting{}, Data: nil, Command: "COMMENT '用户设置表';"},
		{TableModel: &models.AdminLogs{}, Data: nil, Command: "COMMENT '日志表';"},
		{TableModel: &models.Access{}, Data: nil, Command: "COMMENT '钱包资产管理表';"},
		{TableModel: &models.AdminUser{}, Data: AdminUserInit(), Command: "COMMENT '管理表';"},
		{TableModel: &models.User{}, Data: nil, Command: "COMMENT '用户表' AUTO_INCREMENT = 24587;"},
	})

	// 依赖管理员表，创建管理表之后再设置权限
	authItems, authChildren, authAssignments := LoadAdminAuth()
	createTable([]*TableInfo{
		{TableModel: &models.AuthItem{}, Data: authItems, Command: "COMMENT '权限目录表';"},
		{TableModel: &models.AuthChild{}, Data: authChildren, Command: "COMMENT '权限目录子集表';"},
		{TableModel: &models.AuthAssignment{}, Data: authAssignments, Command: "COMMENT '权限分配表';"},
	})

	//	重置默认商户配置
	if _, err := admins.MerchantReset(&dtoadmins.MerchantParams{Id: models.AdminMerchat}); err != nil {
		return
	}

	// 配置初始化用户数据
	if result := database.Db.Create(UserInit()); result.Error != nil {
		return
	}
	if result := database.Db.Create(WalletsOrderInit()); result.Error != nil {
		return
	}
	if result := database.Db.Create(WalletsBillInit()); result.Error != nil {
		return
	}
	if result := database.Db.Create(WalletsUserAssetsInit()); result.Error != nil {
		return
	}
	if result := database.Db.Create(WalletsUserAccountInit()); result.Error != nil {
		return
	}
}

// createTable 创建表
func createTable(params []*TableInfo) {
	for _, v := range params {
		if !database.Db.Migrator().HasTable(v.TableModel) {
			if err := database.Db.AutoMigrate(v.TableModel); err != nil {
				panic(err)
				return
			}

			// 获取表名
			if result, err := database.Db.Migrator().TableType(v.TableModel); err == nil {
				database.Db.Exec("ALTER TABLE " + result.Name() + " " + v.Command)
			} else {
				panic(err)
				return
			}

			// 初始化数据
			if v.Data != nil {
				database.Db.Create(v.Data)
			}
		}
	}
}
