package wire

import "testing"

// 测试 wire对 error 返回值的处理
func TestError(t *testing.T) {
	_, e := InitService()
	t.Log(e)
}

// 测试 带清理资源的函数的方法
func TestServiceWithClean(t *testing.T) {
	_, cleanFunc, e := NewServiceClean()
	defer cleanFunc()
	if e != nil {
		t.Log(e)
	}
}
