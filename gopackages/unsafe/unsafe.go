package main

import (
	"fmt"
	"unsafe"
)
// 如何得到一个对象所占内存大小？
// https://www.cnblogs.com/-wenli/p/12681044.html

// Go 普通指针类型、unsafe.Pointer、uintptr之间的关系
// https://www.cnblogs.com/-wenli/p/12682477.html

// 各种类型所占的字节数
// https://books.studygolang.com/gopl-zh/ch13/ch13-01.html

//Sizeof函数返回的大小只包括数据结构中固定的部分，例如字符串对应结构体中的指针和字符串长度部分，但是并不包含指针指向的字符串的内容。
//Go语言中非聚合类型通常有一个固定的大小，尽管在不同工具链下生成的实际大小可能会有所不同。
//考虑到可移植性，引用类型或包含引用类型的大小在32位平台上是4个字节，在64位平台上是8个字节。
//
//计算机在加载和保存数据时，如果内存地址合理地对齐的将会更有效率。例如2字节大小的int16类型的变量地址应该是偶数，
//一个4字节大小的rune类型变量的地址应该是4的倍数，一个8字节大小的float64、uint64或64-bit指针类型变量的地址应该是8字节对齐的。但是对于再大的地址对齐倍数则是不需要的，即使是complex128等较大的数据类型最多也只是8字节对齐。
//
//由于地址对齐这个因素，一个聚合类型（结构体或数组）的大小至少是所有字段或元素大小的总和，或者更大因为可能存在内存空洞。
//内存空洞是编译器自动添加的没有被使用的内存空间，用于保证后面每个字段或元素的地址相对于结构或数组的开始地址能够合理地对齐（译注：内存空洞可能会存在一些随机数据，可能会对用unsafe包直接操作内存的处理产生影响）。

type NextYear struct {
	Nums   []int  `json:"nums"` // 24
	Year   int8   `json:"year"` // 1
	Moon   int8   `json:"moon"` // 1
	Day    int    `json:"day"` //8
	Second string `json:"second"` //16 2个机器字

	A byte  // 1
	B int64 // 8
	C byte  // 1
}

var ny = NextYear{}

func main() {
	var i = 10
	ip := &i
	fmt.Println(ip)
	var fp *float64 = (*float64)(unsafe.Pointer(ip))
	fmt.Println("fp", fp)
	var name = "jackBai"
	fmt.Println(bytes2str(str2bytes(name)))
	// Sizeof函数返回操作数在内存中的字节大小（返回该类型所占用的内存大小）
	fmt.Println(unsafe.Sizeof("true"))                // 16
	fmt.Println(unsafe.Sizeof(true))                  // 1
	fmt.Println(unsafe.Sizeof(int(0)))                //8
	fmt.Println("int8", unsafe.Sizeof(int8(10)))      // 1
	fmt.Println(unsafe.Sizeof(int16(10)))             // 2
	fmt.Println(unsafe.Sizeof(int32(10000000)))       // 4
	fmt.Println(unsafe.Sizeof(int64(10000000000000))) // 8
	fmt.Println(unsafe.Sizeof(int(1)))                // 8
	fmt.Println(unsafe.Sizeof(float64(0)))            // 8
	fmt.Println("string", unsafe.Sizeof("1"))         // 16
	fmt.Println("ny-sizeof", unsafe.Sizeof(ny))
	fmt.Println("ny", unsafe.Alignof(ny))
	fmt.Println("ny-nums", unsafe.Offsetof(ny.Nums))
}

func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
