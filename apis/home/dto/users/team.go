package dtousers

// TeamDetailsData 用户团队数据
type TeamDetailsData struct {
	InviteNums    int64             `json:"inviteNums"`        //邀请人数
	TodayNums     int64             `json:"todayNums"`         //今日人数
	BuyAmount     float64           `json:"buyAmount"`         //充值金额(今日)
	TodayAmount   float64           `json:"todayAmount"`       //充值金额(昨日)
	TeamEarnings  float64           `json:"teamEarnings"`      //团队总收益
	TodayEarnings float64           `json:"todayEarnings"`     //团队收益(今日)
	CurrentInfo   *TeamIndexInfo    `json:"currentInfo"`       //当前用户信息
	Children      []*WalletBillInfo `json:"children" gorm:"-"` //收益记录列表
}

// WalletBillInfo 团队收益信息
type WalletBillInfo struct {
	Id        int     `json:"id"`                              //订单Id
	Name      string  `json:"name"`                            //收益标题
	Avatar    string  `json:"avatar"`                          //头像
	UserName  string  `json:"username" gorm:"column:username"` //用户名
	NickName  string  `json:"nickname" gorm:"column:nickname"` //用户昵称
	CreatedAt int     `json:"createdAt"`                       //创建时间
	Money     float64 `json:"money"`                           //收益
}

// TeamIndexInfo 团队成员列表
type TeamIndexInfo struct {
	Id        int              `json:"id"`                              //用户ID
	Avatar    string           `json:"avatar"`                          //头像
	UserName  string           `json:"username" gorm:"column:username"` //用户名
	NickName  string           `json:"nickname" gorm:"column:nickname"` //用户昵称
	Depth     int              `json:"depth" gorm:"-"`                  //深度
	Earnings  float64          `json:"earnings" gorm:"column:earnings"` //团队收益
	CreatedAt int              `json:"createdAt"`                       //注册时间
	Children  []*TeamIndexInfo `json:"children" gorm:"-"`               //下级用户列表
}

// TeamIndexParams 团队成员列表参数
type TeamIndexParams struct {
	Id int `json:"id" validate:"required"` //用户Id
}

type TeamDetailsParams struct {
	Id int `json:"id" validate:"required"` //用户Id
}
