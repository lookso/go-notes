/*
@Time : 2019-08-09 20:04 
@Author : Tenlu
@File : defer_call
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
)

func deferCall() {
	defer func() { fmt.Println("print before recover") }()
	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	defer func() { fmt.Println("print after recover") }()

	panic("panic info paipai sble")
}
func main() {
	deferCall()
	
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	defer mycall(yourcall())
	
}
func mycall(num int) int {

	fmt.Println(strconv.Itoa(num)+"-mycall")
	return 111
}
func yourcall() int {
	fmt.Println("222-yourcall")
	return 222
}
func calc(index string, a, b int) int {
	res := a + b
	fmt.Println(index, a, b, res)
	return res
}

/**
10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4

*/
