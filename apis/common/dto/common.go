package dto

const (
	ActivateStatus = 10
)

// IndexData 列表数据
type IndexData struct {
	Items interface{} `json:"items"` //	列表数据
	Count int64       `json:"count"` //	总数
}

// DeleteParams 删除的参数
type DeleteParams struct {
	Ids []int `json:"ids" validate:"required"`
}

// TranslateOptions 翻译数据
type TranslateOptions struct {
	Label string `json:"label"` //	翻译名
	Value string `json:"value"` //	翻译值
}

// WalletOrderAgreeParams 钱包订单同意参数
type WalletOrderAgreeParams struct {
	AssetId  int     //	资产ID
	AdminId  int     //	管理ID
	ParentId int     //	上级ID
	UserId   int     //	用户ID
	SourceId int     //	来源ID
	BillType int     //	类型
	Balance  float64 //	余额
	Money    float64 //	金额
}
