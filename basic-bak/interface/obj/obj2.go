package main

import "fmt"

type Human interface {
	run() string
	eat() string
	Post() Human
}

// 接口里方法必须都被实现

type OurPeople struct {
	country string
}

func (o *OurPeople) run() string {
	return "run:" + o.country
}
func (o *OurPeople) eat() string {
	return "eat:" + o.country
}
func (o *OurPeople) Post() Human{
	return o
}

// 通过结构体,返回接口类型
func (o *OurPeople) Use() Human {
	return o
}

//
func main() {
	ourP := OurPeople{country: "china"}
	var human Human
	human = &ourP
	fmt.Println(human.run())
	fmt.Println(human.eat())
	fmt.Println(ourP.Use().run())
}
