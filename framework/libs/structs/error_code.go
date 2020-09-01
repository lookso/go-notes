package structs

const (
	// 框架自己定义的错误码，统一API错误码
	// 需要APP进一步处理，比如先去绑定手机等

	//--------------
	// 比较基础的错误
	//--------------
	// 未提供必填参数
	ErrCodeBadRequest = 100400
	// 未登录或登录已过期
	ErrCodeUnauthorized = 100401
	// 无权访问
	ErrCodeForbidden = 100403
	// 服务器内部错误
	ErrCodeServerError = 100500
)