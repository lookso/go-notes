package main

import (
	"fmt"
	"reflect"
)

func main() {
	var userInfo = map[string]string{"name": "toms"}
	var num int
	type Address struct {
		Name string
	}
	var address Address
	fmt.Println(reflect.TypeOf(userInfo).Kind())
	fmt.Println(reflect.TypeOf(num).Kind())
	// Kind返回v持有的值的分类，如果v是Value零值，返回值为Invalid
	if reflect.Struct == reflect.TypeOf(address).Kind() {
		fmt.Println("right")
	}
	var a = 1024
	name:=reflect.ValueOf(a).String()
	fmt.Println(name) // <int Value>
}
