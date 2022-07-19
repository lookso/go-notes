package wire

import (
	"fmt"
	"os"
	"testing"
)

// 测试 不使用wire注入代码的使用
func TestOld(t *testing.T) {
	message := NewMessage()
	greeter := NewGreeter(message)
	event := NewEvent(greeter)

	event.Start()
}

// 测试 使用wire生成的代码
func TestWireGen(t *testing.T) {
	e := InitGreeteEvent1()
	e.Start()
}

// 测试 错误处理
// 根据时间内戳的奇偶性产生不结果，需要点耐心
func TestOldHandleError(t *testing.T) {
	message := NewMessage()
	greeter := NewGreeter2(message)
	e, err := NewEvent2(greeter)

	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}

// 测试 wire 中错误处理
// 错误具体处理肯定是程序员自己来，这里主要是看下 wire 生成代码里对错误的处理
func TestHandleError(t *testing.T) {
	e, err := InitGreeteEvent2()
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}

// 测试 wire 参数传入生成代码
func TestInputParams(t *testing.T) {
	e, err := InitGreeteEvent3("hello world")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()
}

