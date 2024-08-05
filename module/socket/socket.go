package socket

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"sync"
	"time"
)

const (
	HeartbeatTime = 30 * time.Second
)

// Instance 初始化socket实例
var Instance = &Socket{
	clientConnMaps:   &ConnMaps{maps: make(map[string]*websocket.Conn), sync: sync.Mutex{}},
	clientBindUserId: &BindUserIdMaps{maps: make(map[int]string), sync: sync.Mutex{}},
	clientSubscribe:  &SubscribeMaps{eventMaps: make(map[string][]*Subscribe), sync: sync.Mutex{}},
}

// Socket 套接字
type Socket struct {
	sync             sync.Mutex                    //	锁操作
	clientConnMaps   *ConnMaps                     //	客户端连接Maps
	clientBindUserId *BindUserIdMaps               //	uuid绑定用户
	clientSubscribe  *SubscribeMaps                // uuid订阅数据
	onMessageFunc    func(uuid string, msg []byte) //	消息事件
	onCloseFunc      func(uuid string)             //	关闭事件
	onBindFunc       func(uuid string, userId int) //	绑定事件
}

// NewSocketConn 创建socket 连接
func NewSocketConn() fiber.Handler {
	return websocket.New(func(conn *websocket.Conn) {
		// 设置当前连接对象
		uuidStr := uuid.New().String()
		Instance.SetClientConn(uuidStr, conn)
		_ = conn.WriteMessage(websocket.TextMessage, []byte(uuidStr))

		defer func() {
			Instance.OnClose(uuidStr)
		}()

		var (
			msg []byte
			err error
		)

		// 心跳包设置
		go sendHeartbeat(conn)

		// 处理业务
		for {
			// 读取消息
			if _, msg, err = conn.ReadMessage(); err != nil {
				break
			}

			// 消息事件
			Instance.OnMessage(uuidStr, msg)
		}
	})
}

// Subscribe 订阅数据
func (_Socket *Socket) Subscribe(uuid string, event string, args ...string) {
	_Socket.clientSubscribe.Subscribe(uuid, event, args...)
}

// UnSubscribe 取消订阅
func (_Socket *Socket) UnSubscribe(uuid string, event string, args ...string) {
	_Socket.clientSubscribe.UnSubscribe(uuid, event, args...)
}

// GetSubscribes 获取事件所有者
func (_Socket *Socket) GetSubscribes(event string) map[string]*Subscribe {
	return _Socket.clientSubscribe.GetSubscribes(event)
}

// sendHeartbeat 心跳包处理
func sendHeartbeat(conn *websocket.Conn) {
	ticker := time.NewTicker(HeartbeatTime)
	defer ticker.Stop()

	for range ticker.C {
		if err := conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
			_ = conn.Close()
			return
		}
	}
}
