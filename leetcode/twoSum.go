/*
@Time : 2019-10-28 16:48 
@Author : Tenlu
@File : twonumsum
@Software: GoLand
*/

// 两数之和
package main

import "fmt"

func main() {
	nums := []int{3, 2, 23, 6, 0}
	existsKeys := twoSum(nums, 6)
	fmt.Println("existsKey", existsKeys)

	fmt.Println("addTwoNum", addTwoNum(10, 90))
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	existsKeys := make([]int, 0)

	for k, v := range nums {
		if _, ok := m[target-v]; ok {
			existsKeys = append(existsKeys, k, m[target-v])
		}
		m[v] = k
	}
	return existsKeys
}

// 两数相加之和

func addTwoNum(a, b int) int {

	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	c := a ^ b        // 不进位情况
	d := (a & b) << 1 // 进位情况，当不为0需要左移一位。为0时，递归结束
	return addTwoNum(c, d)
}

