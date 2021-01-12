/*
@Time : 2019-08-15 21:21
@Author : Tenlu
@File : intersection
@Software: GoLand
*/
package main

import "fmt"

// https://leetcode-cn.com/problems/intersection-of-two-arrays/
// 获取两个数组中相同的元素
func main() {
	var num1 = []int{1, 4, 3, 6, 7, 3, 9, 6, 4, 3, 10}
	var num2 = []int{5, 4, 2, 3, 4, 5, 6, 3}

	var num3 = []int{1, 2, 2, 1}
	var num4 = []int{2, 2}
	fmt.Println(intersection(num1, num2))
	fmt.Println(intersection2(num3, num4))
	fmt.Println(intersect3(num3, num4))
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
func intersection2(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	m1 := map[int]int{}
	var s = make([]int, 0)
	for _, v := range nums1 {
		//遍历nums1，初始化map
		m0[v] = v
	}
	for _, v := range nums2 {
		//遍历nums2，初始化map
		m1[v] = v
	}
	if len(m0) > len(m1) {
		m0, m1 = m1, m0
	}
	for v := range m0 {
		if _, has := m1[v]; has {
			s = append(s, v)
		}
	}
	return s
}

// 升级版
//GO
func intersect3(nums1 []int, nums2 []int) []int {
	m0 := map[int]int{}
	for _, v := range nums1 {
		//遍历nums1，初始化map
		m0[v] += 1
	}
	k := 0
	for _, v := range nums2 { // 2,2
		//如果元素相同，将其存入nums2中，并将出现次数减1
		if m0[v] > 0 {
			m0[v] -= 1
			nums2[k] = v
			k++
		}
	}
	return nums2[0:k]
}
