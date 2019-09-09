/*
@Time : 2019-08-27 09:30 
@Author : Tenlu
@File : struct
@Software: GoLand
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string // 不可比较
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//if sm1 == sm2 {
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}
}

// go run -gcflags -m struct.go
