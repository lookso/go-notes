package main

import (
	"fmt"
)

type InterfacePeople interface {
	Speak(string) string
}

type Stduent struct{}

func (stu Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}
func main() {
	var peo InterfacePeople = Stduent{}
	think := "bitch"
	fmt.Println(peo.Speak(think))

	var mynum interface{}
	mynum = 12
	fmt.Println(mynum)
	y, ok := mynum.(int)
	if ok {
		fmt.Println(y)
	}
	switch mynum.(type) {
	case int:
		fmt.Println("int")
		break
	case string:
		fmt.Println("string")
		break
	default:
		fmt.Println("other")
		break
	}

	var x *int = nil
	Foo(x)

	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)

}

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}
