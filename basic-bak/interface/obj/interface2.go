/*
@Time : 2019-02-17 11:51 
@Author : Tenlu
@File : interface2
@Software: GoLand
*/
package main

import "fmt"

type NewPerson interface {
	Area() string
	run() float32
	retStruct() *NewMan
}

type NewMan struct {
	name string
	age  float32
}

type NewWomen struct {
	name string
	age  float32
}

func (m *NewMan) Area() string {
	return "my name is " + m.name
}
func (m *NewMan) run() float32 {
	return m.age
}

func (m *NewMan) retStruct() *NewMan {
	return &NewMan{name: "jerry", age: 14}
}

//func (m *NewWomen) Area() string {
//	return "my name is " + m.name
//}
//func (m *NewWomen) run() float32 {
//	return m.age
//}
//func (m *NewWomen) retStruct() *NewWomen {
//	return &NewWomen{name: "jeny", age: 14}
//}

func main() {
	man := NewMan{name: "jack", age: 12.02}
	p := NewPerson(&man)

	fmt.Println(p.retStruct().name)
	fmt.Printf("area's value:%f\nrun's value:%s\n", p.run(), p.Area())

	list := []NewPerson{
		&NewMan{name: "toms", age: 12},
		&NewMan{name: "polly", age: 22},
	//	&NewWomen{name:"lucy",age:10},
	}
	for _,l := range list{
		fmt.Println("run:",l.run(),"area:",l.Area())
	}
}
