package cache

import (
	"github.com/gomodule/redigo/redis"
)

// InitConsume 初始化订阅消费
func (_SubScribe *SubScribe) InitConsume() {
	for {
		switch n := _SubScribe.subConn.Receive().(type) {
		case error:
			// 重新初始化订阅消息
			_SubScribe.subConn = redis.PubSubConn{Conn: Rds.Get()}
			for name := range _SubScribe.clientMaps {
				_ = _SubScribe.subConn.Subscribe(redis.Args{}.Add(name)...)
			}

			go _SubScribe.InitConsume()
			return
		case redis.Message:
			if _, ok := _SubScribe.clientMaps[n.Channel]; ok {
				_SubScribe.clientMaps[n.Channel](n.Data)
			}
		}
	}
}
