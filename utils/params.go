package utils

import (
	"time"

	"gorm.io/gorm"
)

const (
	// WhereInInt IN操作 整型
	WhereInInt = 1
	// WhereInStr IN操作 字符串
	WhereInStr = 2
	// WhereLike LIKE操作
	WhereLike = 11
	// WhereEqInt 全等 整型
	WhereEqInt = 21
	// WhereEqStr 全等 字符串
	WhereEqStr = 22
	// WhereRangeDate 范围时间
	WhereBetweenDate = 31
)

// FilterParams 过滤参数对象
type FilterParams struct {
	params []*Filter
}

// RangeDatePicker 时间范围
type RangeDatePicker struct {
	From string `json:"from"` //	开始时间
	To   string `json:"to"`   //	结束时间
}

// Filter 过滤参数
type Filter struct {
	Type  int         //	过滤类型
	Field string      //	字段参数
	Value interface{} //	字段值
}

// NewFilterParams 创建过滤器
func NewFilterParams() *FilterParams {
	return &FilterParams{}
}

// InInt IN 整型
func (_FilterParams *FilterParams) InInt(field string, value interface{}) *FilterParams {
	_FilterParams.params = append(_FilterParams.params, &Filter{
		Type:  WhereInInt,
		Field: field,
		Value: value,
	})
	return _FilterParams
}

// InStr IN 字符串
func (_FilterParams *FilterParams) InStr(field string, value interface{}) *FilterParams {
	_FilterParams.params = append(_FilterParams.params, &Filter{
		Type:  WhereInStr,
		Field: field,
		Value: value,
	})
	return _FilterParams
}

// Like 模糊查询
func (_FilterParams *FilterParams) Like(field string, value interface{}) *FilterParams {
	_FilterParams.params = append(_FilterParams.params, &Filter{
		Type:  WhereLike,
		Field: field,
		Value: value,
	})
	return _FilterParams
}

// EqInt 全等整型
func (_FilterParams *FilterParams) EqInt(field string, value interface{}) *FilterParams {
	_FilterParams.params = append(_FilterParams.params, &Filter{
		Type:  WhereEqInt,
		Field: field,
		Value: value,
	})
	return _FilterParams
}

// EqStr 全等字符串
func (_FilterParams *FilterParams) EqStr(field string, value interface{}) *FilterParams {
	_FilterParams.params = append(_FilterParams.params, &Filter{
		Type:  WhereEqStr,
		Field: field,
		Value: value,
	})
	return _FilterParams
}

// BetweenDate 时间范围
func (_FilterParams *FilterParams) BetweenDate(field string, value interface{}) *FilterParams {
	_FilterParams.params = append(_FilterParams.params, &Filter{
		Type:  WhereBetweenDate,
		Field: field,
		Value: value,
	})
	return _FilterParams
}

func (_FilterParams *FilterParams) Scopes() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		tableName := db.Statement.Table
		for _, v := range _FilterParams.params {
			switch v.Type {
			case WhereInInt:
				//	IN 整型
				if len(v.Value.([]int)) > 0 {
					db.Where(tableName+"."+v.Field, v.Value)
				}
			case WhereInStr:
				//	IN 字符串
				if len(v.Value.([]string)) > 0 {
					db.Where(tableName+"."+v.Field, v.Value)
				}
			case WhereLike:
				//	LIKE 模糊查询
				value := v.Value.(string)
				if value != "" && value != "%" && value != "%%" {
					db.Where(tableName+"."+v.Field, v.Value)
				}
			case WhereEqInt:
				//	相等整型
				if v.Value.(int) != 0 {
					db.Where(tableName+"."+v.Field, v.Value)
				}
			case WhereEqStr:
				//	相当字符串
				if v.Value.(string) != "" {
					db.Where(tableName+"."+v.Field, v.Value)
				}
			case WhereBetweenDate:
				//	时间范围
				value := v.Value.(*RangeDatePicker)
				if value != nil {
					fromLen := len(value.From)
					toLen := len(value.To)

					if (fromLen == 10 || fromLen == 19) && (toLen == 10 || toLen == 19) {

						var starTime, endTime time.Time
						if len(value.From) == 19 {
							starTime, _ = time.ParseInLocation("2006/01/02 15:04:05", value.From, time.Local)
						} else {
							starTime, _ = time.ParseInLocation("2006/01/02", value.From, time.Local)
						}

						if len(value.To) == 19 {
							endTime, _ = time.ParseInLocation("2006/01/02 15:04:05", value.To, time.Local)
						} else {
							endTime, _ = time.ParseInLocation("2006/01/02", value.To, time.Local)
						}

						db.Where(tableName+"."+v.Field, starTime.Unix(), endTime.Unix()+86399)
					}
				}
			default:
				db.Where(v.Field, v.Value)
			}
		}
		return db
	}
}
