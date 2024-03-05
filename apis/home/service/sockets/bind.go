package sockets

import (
	dtosockets "basic/apis/home/dto/sockets"
	"basic/module/socket"
)

// BindUserId 绑定用户ID
func BindUserId(adminId, userId int, params *dtosockets.BindParams) (interface{}, error) {

	// 绑定用户ID
	socket.Instance.BindUserId(params.UUID, userId)

	return "ok", nil
}
