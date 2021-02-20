package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Phones []interface{} `json:"phone"`
}

var res = make(map[string]*User, 0)

func main() {
	phones := []interface{}{1234567890}
	res["user"] = &User{
		Name:   "jack",
		Age:    10,
		Phones: phones,
	}
	rt:=reflect.TypeOf(res)
	fmt.Println(rt) // map[string]*main.User
	rv:=reflect.ValueOf(res)
	fmt.Println(rv) //map[user:0xc000072180]

}
