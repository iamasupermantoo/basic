package dtousers

// LevelIndexData 等级列表数据
type LevelIndexData struct {
	Id    int    `json:"id"`    //Id
	Name  string `json:"name"`  //等级名称
	Icon  string `json:"icon"`  //等级图标
	Money string `json:"money"` //金额
	Level int    `json:"level"` //等级
	Days  int    `json:"days"`  //过期时间
	Data  string `json:"data"`  //数据
	Desc  string `json:"desc"`  //详情
}

// LevelOrderParams 用户购买等级参数
type LevelOrderParams struct {
	Id int `json:"id" validate:"required"` //等级Id
}
