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
}
