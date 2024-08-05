package socket

import (
	"sync"
)

// SubscribeMaps 订阅数据
type SubscribeMaps struct {
	sync      sync.Mutex
	eventMaps map[string][]*Subscribe
}

// Subscribe 订阅参数
type Subscribe struct {
	Event string
	Args  []string
}

// Subscribe 订阅消息
func (_SubscribeMaps *SubscribeMaps) Subscribe(uuid string, event string, args ...string) {
	subscribe := _SubscribeMaps.GetSubscribe(uuid, event)
	if subscribe != nil {
		for _, arg := range args {
			if _SubscribeMaps.stringIndexOf(subscribe.Args, arg) == -1 {
				subscribe.Args = append(subscribe.Args, arg)
			}
		}
	}
}

// UnSubscribe 取消订阅
func (_SubscribeMaps *SubscribeMaps) UnSubscribe(uuid string, event string, args ...string) {
	subscribe := _SubscribeMaps.GetSubscribe(uuid, event)
	if subscribe != nil {
		for _, arg := range args {
			indexOf := _SubscribeMaps.stringIndexOf(subscribe.Args, arg)
			if indexOf > -1 {
				subscribe.Args = append(subscribe.Args[:indexOf], subscribe.Args[indexOf+1:]...)
			}
		}
	}
}

// Delete 删除订阅数据者
func (_SubscribeMaps *SubscribeMaps) Delete(uuid string) {
	_SubscribeMaps.sync.Lock()
	defer _SubscribeMaps.sync.Unlock()
	delete(_SubscribeMaps.eventMaps, uuid)
}

// GetSubscribe 获取当前订阅事件
func (_SubscribeMaps *SubscribeMaps) GetSubscribe(uuid string, event string) *Subscribe {
	_SubscribeMaps.sync.Lock()
	defer _SubscribeMaps.sync.Unlock()

	if _, ok := _SubscribeMaps.eventMaps[uuid]; ok {
		for _, subscribe := range _SubscribeMaps.eventMaps[uuid] {
			if subscribe.Event == event {
				return subscribe
			}
		}
	}
	return nil
}

// GetSubscribes 获取事件所有订阅者
func (_SubscribeMaps *SubscribeMaps) GetSubscribes(event string) map[string]*Subscribe {
	_SubscribeMaps.sync.Lock()
	defer _SubscribeMaps.sync.Unlock()

	data := map[string]*Subscribe{}
	for uuid, subscribes := range _SubscribeMaps.eventMaps {
		for _, subscribe := range subscribes {
			if subscribe.Event == event {
				data[uuid] = subscribe
			}
		}
	}
	return data
}

func (_SubscribeMaps *SubscribeMaps) stringIndexOf(argsList []string, arg string) int {
	for k, v := range argsList {
		if v == arg {
			return k
		}
	}
	return -1
}
