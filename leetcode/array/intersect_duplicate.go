package main

import (
	"fmt"
	"sort"
)

// 两个数组中重复的元素
// https://blog.csdn.net/qq_46595591/article/details/107313619
func main() {
	var arr1 = []int{1, 4, 5, 6, 8, 10}
	var arr2 = []int{3, 5, 7, 8, 11}
	fmt.Println(intersectDuplicate(arr1, arr2))
	fmt.Println(intersectDuplicate2(arr1, arr2))
}

func intersectDuplicate(a, b []int) []int {
	sort.Ints(a)
	sort.Ints(b)                         // 对两数组排序
	res := make([]int, 0, len(a)+len(b)) // 创建第三个数组
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] { // 如果两个元素不相同，对小元素所在数组后一位遍历
			i++
		} else if a[i] > b[j] {
			j++ //如果两个元素不相同，对大元素所在数组后一位遍历
		} else {
			res = append(res, a[i]) //判断两个数组中的相同元素，将相同元素添加到第三个数组中，并将两个数组向右一位移动
			i++
			j++
		}
	}
	return res
}

func intersectDuplicate2(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersectDuplicate2(nums2, nums1) //获取两个数组中的全部记录
	}
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++ //遍历较长数组，并记录各元素出现次数
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if m[num] > 0 {
			res = append(res, num) //遍历较短数组，如果出现的元素次数大于0，将其放到第三个数组中
			m[num]--
		}
	}
	return res
}
