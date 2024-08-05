package utils

// Response 返回信息
type ResponseJson struct {
	Code int         `json:"code"` //	错误代码
	Data interface{} `json:"data"` //	内容
	Msg  string      `json:"msg"`  //	错误信息
}

// ErrorJson 错误信息
func ErrorJson(err error) *ResponseJson {
	return NewResponseJson().Error(err.Error())
}

// SuccessJson 成功信息
func SuccessJson(data interface{}) *ResponseJson {
	return NewResponseJson().Success(data)
}

// NewResponseJson 创建返回Json对象
func NewResponseJson() *ResponseJson {
	return &ResponseJson{}
}

// Success 成功
func (_ResponseJson *ResponseJson) Success(data interface{}) *ResponseJson {
	_ResponseJson.Data = data
	return _ResponseJson
}

// Error 错误
func (_ResponseJson *ResponseJson) Error(msg string) *ResponseJson {
	_ResponseJson.Msg = msg
	_ResponseJson.Code = -1
	return _ResponseJson
}
