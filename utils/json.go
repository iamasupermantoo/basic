package utils

import "github.com/goccy/go-json"

// JsonMarshal 对象转字符串
func JsonMarshal(data interface{}) string {
	if data == nil {
		return ""
	}
	dataBytes, _ := json.Marshal(data)
	return string(dataBytes)
}
