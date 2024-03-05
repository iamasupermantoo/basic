package utils

import (
	"gorm.io/gorm"
)

// Pagination 分页
type Pagination struct {
	SortBy      string `json:"sortBy"`      //	排序字段
	Descending  bool   `json:"descending"`  //	排序 真DESC 假ASC
	Page        int    `json:"page"`        //	当前页数
	RowsPerPage int    `json:"rowsPerPage"` //	每页显示条数
}

// Paginate 分页处理
func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := db.Statement.Table

		if pagination == nil || pagination.RowsPerPage <= 0 {
			return db
		}

		//	页数不对的话
		if pagination.Page <= 0 {
			pagination.Page = 1
		}

		if pagination.SortBy != "" && pagination.Descending {
			orderStr := tableName + "." + pagination.SortBy + " DESC"
			if !pagination.Descending {
				orderStr = tableName + "." + pagination.SortBy + " ASC"
			}
			db.Order(orderStr)
		}

		return db.Offset((pagination.Page - 1) * pagination.RowsPerPage).Limit(pagination.RowsPerPage)
	}
}
