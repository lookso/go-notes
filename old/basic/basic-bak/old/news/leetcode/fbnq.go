/*
@Time : 2019/5/26 3:50 PM 
@Author : Tenlu
@File : fbnq
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	fmt.Println("fbnq num:",fbnq(6))
}
// 索引从0开始
func fbnq(n int) int {
	// 1 1 2 3 5 8 13 21 34 55`
	if n<2 {
		return 1
	}
	fib:=fbnq(n-1)+fbnq(n-2)
	return fib
}
