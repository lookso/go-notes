package main

import (
	"fmt"
)

// 适配器模式用于转换一种接口适配另一种接口。
//实际使用中Adaptee一般为接口，并且使用工厂函数生成实例。

//Adaptee 是被适配的目标接口
type Adaptee interface {
	SpecificRequest() string
}

//AdapteeImpl 是被适配的目标类
type adapteeImpl struct{}

//NewAdaptee 是被适配接口的工厂函数
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

//SpecificRequest 是目标类的一个方法
func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

// 重点,开始适配
//Target 是适配的目标接口
type Target interface {
	Request() string
}

//Adapter 是转换Adaptee为Target接口的适配器
type adapter struct {
	Adaptee
}

//NewAdapter 是Adapter的工厂函数
func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

//Request 实现Target接口
func (a *adapter) Request() string {
	return a.SpecificRequest()
}

func main() {
	adaptee := NewAdaptee()
	target := NewAdapter(adaptee)
	res := target.Request()  // 在不改动原接口方法的情况下,适配目标接口方法
	fmt.Println(res)
}
