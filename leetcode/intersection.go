/*
@Time : 2019-08-15 21:21 
@Author : Tenlu
@File : intersection
@Software: GoLand
*/
package main

import "fmt"

func main() {
	var num1 = []int{1, 4, 3, 6, 7, 3,9,6,4,3,10}
	var num2 = []int{5, 4, 2, 3, 4, 5, 6, 3}
	fmt.Println(intersection(num1, num2))
}
func intersection(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	res := make([]int, 0)
	for _, v1 := range nums1 {
		set[v1] = true
	}
	for _, v2 := range nums2 {
		if m, ok := set[v2]; ok && m { //nums2里面包含nums1里的元素
			res = append(res, v2)
			set[v2] = false //防止重复输出
		}
	}
	return res
}
