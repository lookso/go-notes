package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"unsafe"
)

type User struct {
	Name   string        `json:"name"`
	Age    int           `json:"age"`
	Phones []interface{} `json:"phone" id:"100"`
}

func main() {
	phones := []interface{}{1234567890}
	user := &User{
		Name:   "jack",
		Age:    10,
		Phones: phones,
	}
	rt := reflect.TypeOf(*user)
	fmt.Println(rt.Name(), rt.Kind()) // User struct

	rv := reflect.ValueOf(*user)

	fmt.Println(rv) //{jack 10 [1234567890]}

	// 遍历结构体所有成员
	for i := 0; i < rt.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		fieldType := rt.Field(i)
		// 输出成员名和tag
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}
	// 通过字段名, 找到字段类型信息
	if catType, ok := rt.FieldByName("Phones"); ok {
		// 从tag中取出需要的tag
		fmt.Println("tag_type", catType.Tag.Get("json"), catType.Tag.Get("id"))
	}

	it := reflect.TypeOf(user.Phones[0])
	fmt.Println(it.Name(), it.Kind()) // int int
	fmt.Println("------")
	ObjectToInterface()
	fmt.Println("------")
	SetObject()
	ubt, _ := json.Marshal(&user)
	fmt.Println(bytes2string(ubt))
}

func ObjectToInterface() {
	var x float64 = 3.4

	v := reflect.ValueOf(x) //v is reflect.Value

	var y float64 = v.Interface().(float64)
	fmt.Println("value:", y)
}

func SetObject() {
	var x float64 = 3.4
	v := reflect.ValueOf(&x)
	v.Elem().SetFloat(7.1)
	fmt.Println("x :", v.Elem().Interface())
}

func string2bytes(s string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))

	bh := reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}

	return *(*[]byte)(unsafe.Pointer(&bh))
}

func bytes2string(b []byte) string {
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))

	sh := reflect.StringHeader{
		Data: sliceHeader.Data,
		Len:  sliceHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&sh))
}
