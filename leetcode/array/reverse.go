package main

import "fmt"

//给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。
//示例 1:
//
//输入: [1,2,3,4,5,6,7] 和 k = 3
//输出: [5,6,7,1,2,3,4]

func main() {
	var nums = []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(rotate(nums, 3))
}

//GO
func rotate(nums []int, k int)[]int {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
	return nums
}

func reverse(arr []int) []int{
	//fmt.Println(len(arr)/2)
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-i-1] = arr[len(arr)-i-1], arr[i]
	}
	fmt.Println(arr)
	return arr
}
