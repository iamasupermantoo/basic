package socket

import (
	"github.com/gofiber/contrib/websocket"
	"sync"
)

// ConnMaps 客户端Maps
type ConnMaps struct {
	sync sync.Mutex
	maps map[string]*websocket.Conn
}

// Set 设置客户端
func (_ConnMaps *ConnMaps) Set(uuid string, conn *websocket.Conn) {
	_ConnMaps.sync.Lock()
	defer _ConnMaps.sync.Unlock()

	_ConnMaps.maps[uuid] = conn
}

// Get 获取客户端
func (_ConnMaps *ConnMaps) Get(uuid string) *websocket.Conn {
	_ConnMaps.sync.Lock()
	defer _ConnMaps.sync.Unlock()

	if _, ok := _ConnMaps.maps[uuid]; ok {
		return _ConnMaps.maps[uuid]
	}
	return nil
}

// Del 删除客户端
func (_ConnMaps *ConnMaps) Del(uuid string) {
	_ConnMaps.sync.Lock()
	defer _ConnMaps.sync.Unlock()

	delete(_ConnMaps.maps, uuid)
}
