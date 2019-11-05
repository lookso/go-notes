/*
@Time : 2019/3/1 1:49 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
	"runtime"
)

func main() {
	fmt.Println("--------strconv--------")
	fmt.Printf("%T\n", strconv.Itoa(10)+",hello") // strconv.Itoa 将int类型转换为string类型
	num, _ := strconv.Atoi("12")                  // strconv.Atoi 将string类型转换为int类型
	fmt.Println(num)

	runtime.GC()

}
