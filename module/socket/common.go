package socket

import (
	"github.com/gofiber/contrib/websocket"
)

// SetClientConn 设置客户端连接
func (_Socket *Socket) SetClientConn(uuid string, conn *websocket.Conn) {
	_Socket.clientConnMaps.Set(uuid, conn)
}

// DelClientConn 删除客户端连接
func (_Socket *Socket) DelClientConn(uuid string) {
	_Socket.clientConnMaps.Del(uuid)
}

// BindUserId 绑定用户ID
func (_Socket *Socket) BindUserId(uuid string, userId int) {
	_Socket.clientBindUserId.Set(userId, uuid)

	//	绑定用户事件
	if _Socket.onBindFunc != nil {
		_Socket.onBindFunc(uuid, userId)
	}
}

// UnbindUserId 解绑用户ID
func (_Socket *Socket) UnbindUserId(userId int) {
	_Socket.clientBindUserId.Del(userId)
}

// OnMessage 消息事件
func (_Socket *Socket) OnMessage(uuid string, msg []byte) {
	if _Socket.onMessageFunc != nil {
		_Socket.onMessageFunc(uuid, msg)
	}
}

// OnClose 关闭事件
func (_Socket *Socket) OnClose(uuid string) {
	// 删除客户端
	_Socket.DelClientConn(uuid)
	// 删除用户绑定
	_Socket.UnbindUserId(_Socket.clientBindUserId.GetUserId(uuid))

	// 删除订阅数据
	_Socket.clientSubscribe.Delete(uuid)

	//	关闭消息事件
	if _Socket.onCloseFunc != nil {
		_Socket.onCloseFunc(uuid)
	}
}
