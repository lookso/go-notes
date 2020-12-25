package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i = 10
	ip := &i
	fmt.Println(ip)
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	fmt.Println(fp)
	var name = "jackBai"
	fmt.Println(bytes2str(str2bytes(name)))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
