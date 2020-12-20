package main

import (
	"fmt"
)

//接口是一个规范
type Usber interface {
	// 最好以 er 结尾表示接口
	start()
	stop()
}

// 如果接口里有方法的话，必须要通过结构体或者通过自定义类型实现这个接口。

type Phone struct {
	Name string
}

// 手机要实现 usb 接口的话必须得实现 usb 接口中的所有方法

func (p *Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p *Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {
	var p Usber = &Phone{Name: "华为手机"} // Golang 中接口就是一个数据类型,表示手机实现 Usb 接口
	p.start()
	p.stop()
}
