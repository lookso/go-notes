/*
@Time : 2019-08-15 22:19 
@Author : Tenlu
@File : singleNumber
@Software: GoLand
*/

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
package main

import "fmt"

func main() {
	var arr = []int{1, 2, 2, 1, 5}
	fmt.Println(singleNumber(arr))
}
func singleNumber(arr []int) int {
	var value = 0
	for i := 0; i < len(arr); i++ {
		value = value ^ arr[i]
		//利用异或特性，使得出现了两次的数字的二进制位始终为0，最后一个单独的数字与0异或就是它自己
	}
	return value
}
