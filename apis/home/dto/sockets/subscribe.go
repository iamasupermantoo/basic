package dtosockets

// SubscribeParam 订阅参数
type SubscribeParam struct {
	UUID  string      `json:"uuid" validate:"required"`  //	标识ID
	Event string      `json:"event" validate:"required"` //	订阅事件
	Data  interface{} `json:"data" validate:"required"`  //	订阅消息
}
