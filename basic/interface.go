package main

import (
	"fmt"
	"reflect"
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
	interfaceTest(123)

	var d Duck = &Food{
		Name: "peanut",
	}
	d.Eat()
	var m Duck = Mss{"name": "banner"}
	m.Eat()
	interType(321)
}
func interfaceTest(it interface{}) {
	if it, ok := it.(int); ok {
		fmt.Println(it)
	}
	switch v := it.(type) {
	case int:
		fmt.Println("整型", v)
		var s int
		s = v
		fmt.Println(s)
	case string:
		fmt.Println("字符串", v)
	}
	fmt.Println(reflect.TypeOf(it))
}

type Duck interface {
	Eat()
}
type Food struct {
	Name string `json:"name"`
}

type Mss map[string]string

func (f *Food) Eat() {
	fmt.Println("to eat " + f.Name)
}
func (m Mss) Eat() {
	fmt.Println("to eat " + m["name"])
}

func interType(name interface{}) {
	switch t := name.(type) {
	case string:
		fmt.Println("the name type is string and value is ", t)
	case int:
		fmt.Println("the name type is int and value is ", t)
	default:
		fmt.Println("unKnow name type")
	}
	// 或者直接判断
	nv, ok := name.(int)
	if ok {
		fmt.Println("the name type is int and value is ", nv)
	}
	return
}
