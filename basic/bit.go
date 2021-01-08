package main

import "fmt"

func main() {
	bitRun()
}

// 位运算
func bitRun() {
	a := 60
	b := 13
	fmt.Printf("binary: %#b, decimal: %b\n", a, b) // 二进制
	fmt.Printf("octal: 0o%o, decimal: %b\n", a, b)// 8进制
	num := 0.1234
	fmt.Printf("floating point: %x, decimal: %f\n", num, num) // 16进制
}
// https://studygolang.com/articles/25288