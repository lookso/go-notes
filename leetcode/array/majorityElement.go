package main

import "fmt"

// https://leetcode-cn.com/problems/majority-element/

func main() {
	var num = []int{2, 3, 1, 2, 1, 6, 5, 9, 10, 3, 3}
	fmt.Println(majorityElement(num))
}
func majorityElement(nums []int) int {
	var count, res = 0, 0
	for _, v := range nums {
		if count == 0 {
			res = v
		}
		if res == v {
			count += 1
		} else {
			count -= 1
		}
	}
	return res
}
