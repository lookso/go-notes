package wire

import (
	"errors"
	"fmt"
	"time"
)

// Message string 的别名
// wire 中使用基本类型最好取个别名
type Message string

// Greeter 打招呼的人
type Greeter struct {
	Message Message
	Grumpy  bool // 演示错误处理时用的属性。
}

// Event 打招呼事件
type Event struct {
	Greeter Greeter
}

// Start 开始打招呼
func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

// Greet 打招呼的方法
func (g Greeter) Greet() Message {
	return g.Message
}

// NewEvent 创建打招呼事件
func NewEvent(g Greeter) Event {
	return Event{Greeter: g}
}

// NewMessage 创建打招呼内容
func NewMessage() Message {
	return Message("Hi there!")
}

// NewMessage2 创建任意打招呼内容
func NewMessage2(phrase string) Message {
	return Message(phrase)
}

// NewGreeter 创建打招呼对象
func NewGreeter(m Message) Greeter {
	return Greeter{Message: m}
}

// NewGreeter2 50%的机会使打招呼的人变得易恕
func NewGreeter2(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	// grumpy = true // 等不及了就打开这句吧
	return Greeter{Message: m, Grumpy: grumpy}
}

// NewEvent2 打招呼的人易恕则出现错误
func NewEvent2(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

// Greet2 打招呼的人易恕则粗暴回应
func (g Greeter) Greet2() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}
