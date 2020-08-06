package main

import "fmt"

type Duck interface {
	Eat()
}
type Food struct {
	Name string `json:"name"`
}

func (f *Food) Eat() {
	fmt.Println("to eat " + f.Name)
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

func main() {
	var d Duck = &Food{
		Name: "peanut",
	}
	d.Eat()
	interType(321)
}
