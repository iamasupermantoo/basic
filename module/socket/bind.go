package socket

import "sync"

type BindUserIdMaps struct {
	sync sync.Mutex
	maps map[int]string
}

// Set 设置绑定用户ID
func (_BindUserIdMaps *BindUserIdMaps) Set(userId int, uuid string) {
	_BindUserIdMaps.sync.Lock()
	defer _BindUserIdMaps.sync.Unlock()

	_BindUserIdMaps.maps[userId] = uuid
}

// Get 获取绑定用户的UUID
func (_BindUserIdMaps *BindUserIdMaps) Get(userId int) string {
	_BindUserIdMaps.sync.Lock()
	defer _BindUserIdMaps.sync.Unlock()

	if _, ok := _BindUserIdMaps.maps[userId]; ok {
		return _BindUserIdMaps.maps[userId]
	}

	return ""
}

// GetUserId 获取用户UserId
func (_BindUserIdMaps *BindUserIdMaps) GetUserId(uuid string) int {
	_BindUserIdMaps.sync.Lock()
	defer _BindUserIdMaps.sync.Unlock()

	for k, v := range _BindUserIdMaps.maps {
		if v == uuid {
			return k
		}
	}

	return -1
}

// GetUUIDList 获取用户uuids列表
func (_BindUserIdMaps *BindUserIdMaps) GetUUIDList(userId int) []string {
	_BindUserIdMaps.sync.Lock()
	defer _BindUserIdMaps.sync.Unlock()

	data := make([]string, 0)
	for k, v := range _BindUserIdMaps.maps {
		if k == userId {
			data = append(data, v)
		}
	}
	return data
}

// Del 删除绑定的用户
func (_BindUserIdMaps *BindUserIdMaps) Del(userId int) {
	if userId > 0 {
		_BindUserIdMaps.sync.Lock()
		defer _BindUserIdMaps.sync.Unlock()

		delete(_BindUserIdMaps.maps, userId)
	}
}
