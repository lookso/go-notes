package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	fmt.Println("各int类型的大小：")
	var i1 int = 1
	var i2 int8 = 2
	var i3 int16 = 3
	var i4 int32 = 4
	var i5 int64 = 5
	var i6 uint64 = 6
	fmt.Printf("int    : %v\n", unsafe.Sizeof(i1))
	fmt.Printf("int8   : %v\n", unsafe.Sizeof(i2))
	fmt.Printf("int16  : %v\n", unsafe.Sizeof(i3))
	fmt.Printf("int32  : %v\n", unsafe.Sizeof(i4))
	fmt.Printf("int64  : %v\n", unsafe.Sizeof(i5))
	fmt.Printf("uint64 : %v\n\n", unsafe.Sizeof(i6))

	// 输出各int类型的取值范围
	fmt.Println("各int类型的取值范围：")
	//fmt.Println("int:", math.MinInt, "~", math.MaxInt) 报错，没有 math.MinInt math.MaxInt
	fmt.Println("int8:", math.MinInt8, "~", math.MaxInt8)
	fmt.Println("int16:", math.MinInt16, "~", math.MaxInt16)
	fmt.Println("int32:", math.MinInt32, "~", math.MaxInt32)
	fmt.Println("int64:", math.MinInt64, "~", math.MaxInt64)
	fmt.Println()

	// n是自动推导类型
	n := 1234567890
	fmt.Printf("n := 1234567890 的默认类型为：%T\n", n)
	fmt.Printf("int类型的字节数为：%v\n\n", unsafe.Sizeof(n))

	// 初始化一个32位整型值
	var a int32 = 987654321

	fmt.Println("var a int32 = 987654321")
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 十六进制为0x%x，十进制为%d\n", a, a)

	// 将a转换为int8类型, 发生数值截断
	b := int8(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int8:  十六进制为0x%x，十进制为%d\n", b, b)

	// 将a转换为int16类型, 发生数值截断
	c := int16(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 十六进制为0x%x，十进制为%d\n", c, c)

	// 将a转换为int64类型
	d := int64(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int64: 十六进制为0x%x，十进制为%d\n", d, d)
	// 将a转换为uint64类型
	e := uint64(a)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("uint64: 十六进制为0x%x，十进制为%d\n", e, e)
}
