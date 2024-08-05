package models

import (
	"basic/module/database"
)

// FindTableColumnIntn 查询表列数据
func FindTableColumnIntn(table string, field string, query string, where string) []int {
	data := make([]int, 0)
	if where == "%" || where == "%%" || where == "" {
		return data
	}
	database.Db.Table(table).Select(field).Where(query, where).Scan(&data)
	if len(data) == 0 {
		data = append(data, -1)
	}
	return data
}

// FindTableColumnString 查询表列字符串数组
func FindTableColumnString(table string, field string, query string, where string) []string {
	data := make([]string, 0)
	if where == "%" || where == "%%" {
		return data
	}
	database.Db.Table(table).Select(field).Where(query, where).Scan(&data)
	if len(data) == 0 {
		data = append(data, "")
	}
	return data
}
