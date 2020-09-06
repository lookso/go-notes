package defined

import "errors"

var (
	EmptyParam      = errors.New("缺少参数")
	ErrParam        = errors.New("参数错误")
	ErrEmptyVersion = errors.New("client version is empty")
	ErrBusy         = errors.New("现在繁忙，请稍后操作")
)

const (
	ErrCodeBusy = 700 // 现在繁忙，请稍后操作
)
