package utils

import (
	"strings"

	"github.com/goccy/go-json"
)

// ArrayStringIndexOf 数组查找字符串是否存在
func ArrayStringIndexOf(arr []string, name string) int {
	for i, v := range arr {
		if v == name {
			return i
		}
	}
	return -1
}

// ArrayStringContainsOf 字符串模糊查找
func ArrayStringContainsOf(arr []string, name string) int {
	for i, v := range arr {
		if strings.Contains(v, name) {
			return i
		}
	}
	return -1
}

// StringToBoolMaps 值转换成 Bool maps
func StringToBoolMaps(value string) map[string]bool {
	data := map[string]bool{}
	_ = json.Unmarshal([]byte(value), &data)
	return data
}
