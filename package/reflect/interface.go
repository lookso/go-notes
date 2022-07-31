package main

import (
	"fmt"
	"reflect"
)

func main() {
	var info interface{}
	info = map[int]string{1: "jack"}
	fmt.Println(reflect.TypeOf(info))
	v := reflect.ValueOf(info)
	fmt.Println(v.MapKeys())
}
