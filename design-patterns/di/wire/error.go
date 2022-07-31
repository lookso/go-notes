package wire

import (
	"errors"
	"fmt"
)

// Service 模拟一个服务
type Service struct {
}

// NewService 模拟创建服务
// 为了演示，此函数必定会出错
func NewService() (Service, error) {
	return Service{}, errors.New("出错演示")
}

// NewServiceClean 模拟创建服务，并清理相关资源
//
// wire对provider的返回值个数和顺序有所规定：
// 	1. 第一个参数是需要生成的依赖对象
// 	2. 如果返回2个返回值，第二个参数必须是func()或者error
// 	3. 如果返回3个返回值，第二个参数必须是func()，第三个参数则必须是error
func NewServiceClean() (Service, func(), error) {
	s := Service{}
	e := errors.New("出错演示")
	fn := func() {
		fmt.Println("执行资源清理函数")
	}
	return s, fn, e
}
