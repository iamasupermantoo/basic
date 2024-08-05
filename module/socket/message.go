package socket

// WriteJSON 写入JSON数据
func (_Socket *Socket) WriteJSON(uuid string, data interface{}) error {
	conn := _Socket.clientConnMaps.Get(uuid)
	if conn != nil {
		return conn.WriteJSON(data)
	}
	return nil
}

// WriteUserJSON 写入用户客户端
func (_Socket *Socket) WriteUserJSON(userId int, data interface{}) {
	uuids := _Socket.clientBindUserId.GetUUIDList(userId)

	for _, uuid := range uuids {
		_ = _Socket.WriteJSON(uuid, data)
	}
}
