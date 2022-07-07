package wire

import "testing"
// 测试 结构体值注入
func TestSetter(t *testing.T) {
	p := InitializeFooBar()
	t.Log(p)
}
