package main

import "fmt"

func main() {
	var num1 = []int{1, 2, 3, 0, 0, 0}
	var num2 = []int{2, 5, 6}
	fmt.Println(merge(num1, 3, num2, 3))
	fmt.Println(merge2(num1, 3, num2, 3))
}
func merge2(nums1 []int, m int, nums2 []int, n int) []int {
	index1 := m - 1
	index2 := n - 1
	tail := m + n - 1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[tail] = nums1[index1]
			index1--
		} else {
			nums1[tail] = nums2[index2]
			index2--
		}
		tail--
	}
	for tail >= 0 && index1 >= 0 {
		nums1[tail] = nums1[index1]
		index1--
		tail--
	}
	for tail >= 0 && index2 >= 0 {
		nums1[tail] = nums2[index2]
		index2--
		tail--
	}
	return nums1
}
func merge(nums1 []int, m int, nums2 []int, n int) []int {
	n1 := m - 1
	n2 := n - 1
	l := m + n - 1
	for n2 >= 0 {
		if n1 >= 0 && nums1[n1] > nums2[n2] {
			nums1[l] = nums1[n1]
			n1--
		} else {
			nums1[l] = nums2[n2]
			n2--
		}
		l--
	}
	return nums1
	//链接：https://leetcode-cn.com/problems/merge-sorted-array/solution/go-shuang-100shuang-zhi-zhen-cong-hou-wang-qian-bi/
}
