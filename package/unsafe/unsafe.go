package main

import (
	"fmt"
	"unsafe"
)

type NextYear struct {
	Year int8 `json:"year"`
	Moon int8 `json:"moon"`
	Day  int  `json:"day"`
}

var ny = NextYear{}

func main() {
	var i = 10
	ip := &i
	fmt.Println(ip)
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	fmt.Println(fp)
	var name = "jackBai"
	fmt.Println(bytes2str(str2bytes(name)))
	// Sizeof函数返回操作数在内存中的字节大小（返回该类型所占用的内存大小）
	fmt.Println(unsafe.Sizeof("true"))                // 16
	fmt.Println(unsafe.Sizeof(true))                  // 1
	fmt.Println(unsafe.Sizeof(int(0)))                //8
	fmt.Println(unsafe.Sizeof(int8(0)))               // 1
	fmt.Println(unsafe.Sizeof(int16(10)))             // 2
	fmt.Println(unsafe.Sizeof(int32(10000000)))       // 4
	fmt.Println(unsafe.Sizeof(int64(10000000000000))) // 8
	fmt.Println(unsafe.Sizeof(int(1)))                // 8
	fmt.Println(unsafe.Sizeof(float64(0)))            // 8

	fmt.Println("ny", unsafe.Alignof(ny))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
