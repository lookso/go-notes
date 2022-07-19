package main

import (
	"fmt"
	"syscall"
)

func main() {
	// 按位或运算符"|"是双目运算符。 其功能是参与运算的两数各对应的二进位相或
	fmt.Println(20|2)
	fmt.Println(syscall.LOCK_EX | syscall.LOCK_NB)
}
