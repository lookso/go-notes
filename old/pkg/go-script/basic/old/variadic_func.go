/*
@Time : 2019/1/17 9:57 AM
@Author : Tenlu
@File : variadic_func
@Software: GoLand
*/
package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)

	// 等同于下面
	// sum(1,2,3,4)
}
