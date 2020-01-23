/*
@Time : 2019-10-28 21:19 
@Author : Tenlu
@File : fbnq
@Software: GoLand
*/
package main

import "fmt"

// 斐波那锲数列
func main() {
	num := fbnq(20)
	fmt.Println(num)
}
func fbnq(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	num := fbnq(n-1) + fbnq(n-2)
	return num
}
