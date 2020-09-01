package structs

import "net/http"

type Response struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Data  interface{} `json:"data"`
}

func Data(obj interface{}) (int, Response) {
	resp := Response{
		Code: 0,
		Msg:  "success",
		Data: obj,
	}
	return 200, resp
}
// BadRequest 请求不完整，缺少必要参数等
func BadRequest(msg string) (int, Response) {
	if msg == "" {
		msg = "缺少参数"
	}
	resp := Response{
		Code: ErrCodeBadRequest,
		Msg:  msg,
		Data: nil,
	}
	return http.StatusBadRequest, resp
}

// UnAuthorized 需要登录
func UnAuthorized(msg string, data interface{}) (int, Response) {

	if msg == "" {
		msg = "请先登录"
	}

	resp := Response{
		Code: ErrCodeUnauthorized,
		Msg:  msg,
		Data: data,
	}
	return http.StatusUnauthorized, resp
}

// Forbidden 无权限访问
func Forbidden(msg string) (int, Response) {

	if msg == "" {
		msg = "无权访问"
	}

	resp := Response{
		Code: ErrCodeForbidden,
		Msg:  msg,
		Data: nil,
	}
	return http.StatusForbidden, resp
}