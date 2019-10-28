/*
@Time : 2019-10-26 08:13 
@Author : Tenlu
@File : nums
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
)
// 不通过乘法和循环获取 1-n之间的输出
func main() {
	run(10, 0)
	fmt.Println("-------getnums---------")
	getNums(11, 0)
}
func run(max int, current int) {

	if current <= max {
		fmt.Println(current)
		current++
		run(max, current)
	} else {
		os.Exit(110)
	}
}

func getNums(max int, current int) {
loop:
	if current <= max {
		fmt.Println(current)
		current++
		goto loop
	} else {
		os.Exit(110)
	}
}
//
