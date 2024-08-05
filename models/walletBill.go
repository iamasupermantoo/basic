package models

// WalletBill 钱包账单
type WalletBill struct {
	Id        int     `gorm:"type:int unsigned primary key auto_increment;comment:主键;"`
	AdminId   int     `gorm:"type:int unsigned not null;comment:管理ID"`
	UserId    int     `gorm:"type:int unsigned not null;comment:用户ID"`
	AssetsId  int     `gorm:"type:int unsigned not null;comment:资产ID"`
	SourceId  int     `gorm:"type:int unsigned not null;comment:来源ID"`
	Type      int     `gorm:"type:smallint not null default 1;default:1;comment:类型 ..."`
	Name      string  `gorm:"type:varchar(60) not null;comment:名称"`
	Money     float64 `gorm:"type:decimal(12,2) not null;comment:金额"`
	Balance   float64 `gorm:"type:decimal(12,2) not null;comment:余额"`
	Data      string  `gorm:"type:text;comment:数据"`
	CreatedAt int     `gorm:"type:int unsigned not null;autoCreateTime;comment:创建时间"`
}

// WalletBillTypeOptionsInfo 钱包账单Options
type WalletBillTypeOptionsInfo struct {
	Label string `json:"label"` //	账单标题
	Value int    `json:"value"` //	账单值
}

// WalletBillDepositTypeList 入金类型列表
var WalletBillDepositTypeList = map[int]string{
	WalletBillTypeDeposit:              "walletBillTypeDeposit",
	WalletBillTypeWithdrawRefuse:       "walletBillTypeWithdrawRefuse",
	WalletBillTypeAssetsWithdrawRefuse: "walletBillTypeAssetsWithdrawRefuse",
	WalletBillTypeEarnings:             "walletBillTypeEarnings",
	WalletBillTypeAward:                "walletBillTypeAward",
	WalletBillTypeAssetsDeposit:        "walletBillTypeAssetsDeposit",
	WalletBillTypeAssetsEarnings:       "walletBillTypeAssetsEarnings",
	WalletBillTypeAssetsAward:          "walletBillTypeAssetsAward",
	WalletBillTypeRegisterAward:        "walletBillTypeRegisterAward",
	WalletBillTypeShareAward:           "walletBillTypeShareAward",
	WalletBillTypeSystemDeposit:        "walletBillTypeSystemDeposit",
	WalletBillTypeSystemAssetsDeposit:  "walletBillTypeSystemAssetsDeposit",
	WalletBillTypeTeamEarnings:         "walletBillTypeTeamEarnings",
	WalletBillTypeTeamAssetsEarnings:   "walletBillTypeTeamAssetsEarnings",
}

// WalletBillSpendTypeList 花费类型列表
var WalletBillSpendTypeList = map[int]string{
	WalletBillTypeWithdraw:             "walletBillTypeWithdraw",
	WalletBillTypeBuyProduct:           "walletBillTypeBuyProduct",
	WalletBillTypeBuyLevel:             "walletBillTypeBuyLevel",
	WalletBillTypeAssetsWithdraw:       "walletBillTypeAssetsWithdraw",
	WalletBillTypeAssetsBuyProduct:     "walletBillTypeAssetsBuyProduct",
	WalletBillTypeSystemWithdraw:       "walletBillTypeSystemWithdraw",
	WalletBillTypeSystemAssetsWithdraw: "walletBillTypeSystemAssetsWithdraw",
}

const (
	// WalletBillAccountDefault 账户账单
	WalletBillAccountDefault = -1

	// WalletBillAccountAssets 资产账单
	WalletBillAccountAssets = -2

	// WalletBillTypeDeposit 充值类型
	WalletBillTypeDeposit = 1

	// WalletBillTypeAssetsDeposit 资产充值类型
	WalletBillTypeAssetsDeposit = 101

	// WalletBillTypeSystemDeposit 余额系统充值
	WalletBillTypeSystemDeposit = 3

	// WalletBillTypeSystemAssetsDeposit 资产系统充值
	WalletBillTypeSystemAssetsDeposit = 103

	// WalletBillTypeWithdraw 提现类型
	WalletBillTypeWithdraw = 11

	// WalletBillTypeAssetsWithdraw 资产提现类型
	WalletBillTypeAssetsWithdraw = 111

	// WalletBillTypeSystemWithdraw 余额系统扣款
	WalletBillTypeSystemWithdraw = 13

	// WalletBillTypeSystemAssetsWithdraw 资产系统扣款
	WalletBillTypeSystemAssetsWithdraw = 113

	// WalletBillTypeWithdrawRefuse 余额提现拒绝
	WalletBillTypeWithdrawRefuse = 15

	// WalletBillTypeAssetsWithdrawRefuse 资产提现拒绝
	WalletBillTypeAssetsWithdrawRefuse = 115

	// WalletBillTypeBuyProduct 购买
	WalletBillTypeBuyProduct = 21

	// WalletBillTypeAssetsBuyProduct 资产购买产品
	WalletBillTypeAssetsBuyProduct = 121

	// WalletBillTypeBuyLevel 购买等级
	WalletBillTypeBuyLevel = 23

	// WalletBillTypeEarnings 收益
	WalletBillTypeEarnings = 51

	// WalletBillTypeAssetsEarnings 资产收益
	WalletBillTypeAssetsEarnings = 151

	// WalletBillTypeAward 奖励
	WalletBillTypeAward = 61

	// WalletBillTypeAssetsAward 资产奖励
	WalletBillTypeAssetsAward = 161

	// WalletBillTypeRegisterAward 注册奖励
	WalletBillTypeRegisterAward = 66

	// WalletBillTypeShareAward 邀请奖励
	WalletBillTypeShareAward = 67

	// WalletBillTypeTeamEarnings 团队收益
	WalletBillTypeTeamEarnings = 71

	// WalletBillTypeTeamAssetsEarnings 团队资产收益
	WalletBillTypeTeamAssetsEarnings = 171
)

// GetBillTypeList 获取账单列表
func GetBillTypeList() map[int]string {
	billOptions := map[int]string{}
	for k, v := range WalletBillDepositTypeList {
		billOptions[k] = v
	}
	for k, v := range WalletBillSpendTypeList {
		billOptions[k] = v
	}
	return billOptions
}

// GetBillAccountTypeList 获取钱包账单账户类型数组
func GetBillAccountTypeList(accountType int) []*WalletBillTypeOptionsInfo {
	billList := GetBillTypeList()
	data := make([]*WalletBillTypeOptionsInfo, 0)

	for k, v := range billList {
		if (accountType == WalletBillAccountDefault && k < 100) || (accountType == WalletBillAccountAssets && k > 100) {
			data = append(data, &WalletBillTypeOptionsInfo{Label: v, Value: k})
		}
	}
	return data
}

// GetBillMoney 获取账单金额
func GetBillMoney(billType int, money float64) float64 {
	if _, ok := WalletBillDepositTypeList[billType]; ok {
		return money
	}
	return 0 - money
}
