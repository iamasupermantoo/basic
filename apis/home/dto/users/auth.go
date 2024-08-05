package dtousers

// AuthCreateParams 新增认证
type AuthCreateParams struct {
	AdminId   int    `json:"-"`
	UserId    int    `json:"-"`
	RealName  string `json:"realName" validate:"required"`  //	证件姓名
	Number    string `json:"number" validate:"required"`    //	证件号码
	Photo1    string `json:"photo1" validate:"required"`    //	证件照1
	Photo2    string `json:"photo2" validate:"required"`    //	证件照2
	Photo3    string `json:"photo3"`                        //	证件照3
	Type      int    `json:"type"`                          //	类型
	Status    int    `json:"-"`                             //	状态
	Telephone string `json:"telephone" validate:"required"` //	手机号码
	Address   string `json:"address" validate:"required"`   //	详细地址
	UpdatedAt int    `json:"-"`                             //	更新时间
	CreatedAt int    `json:"-"`                             //	创造时间
}

// AuthIndexParams 认证参数
type AuthIndexParams struct {
	RealName  string `json:"realName"  validate:"required"`  //	真实姓名
	Number    string `json:"number"  validate:"required"`    //	卡号
	Telephone string `json:"telephone"  validate:"required"` //	电话号码
}

// HomeAuthIndexData 认证数据
type HomeAuthIndexData struct {
	RealName  string `json:"realName"`  //	真实姓名
	Number    string `json:"number"`    //	卡号
	Photo1    string `json:"photo1"`    //	证件照1
	Photo2    string `json:"photo2"`    //	证件照2
	Photo3    string `json:"photo3"`    //	证件照3
	Telephone string `json:"telephone"` //	手机号码
	Address   string `json:"address"`   //	详细地址
	Type      int    `json:"type"`      //	类型
	Status    int    `json:"status"`    //	状态
	Data      string `json:"data"`      //	拒绝理由
	UpdatedAt int    `json:"updatedAt"` //	更新时间
}
