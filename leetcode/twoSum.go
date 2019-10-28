/*
@Time : 2019-10-28 16:48 
@Author : Tenlu
@File : twonumsum
@Software: GoLand
*/
package main

import "fmt"

func main() {
	nums := []int{3, 2, 23, 6,0}
	existsKeys := twoSum(nums, 6)
	fmt.Println("existsKey", existsKeys)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	existsKeys := make([]int, 0)

	for k, v := range nums {
		if _, ok := m[target-v]; ok {
			existsKeys = append(existsKeys, k,m[target-v])
		}
		m[v] = k
	}
	return existsKeys
}
