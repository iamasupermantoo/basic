package sockets

import (
	dtosockets "basic/apis/home/dto/sockets"
	"basic/module/socket"
)

// Subscribe 订阅消息
func Subscribe(params *dtosockets.SubscribeParam) (interface{}, error) {

	// 订阅消息
	socket.Instance.Subscribe(params.UUID, params.Event, params.Data.([]string)...)
	return "ok", nil
}
