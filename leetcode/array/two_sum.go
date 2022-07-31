package main

import "fmt"

// https://leetcode-cn.com/problems/two-sum/
func main(){
	fmt.Println(twoSum([]int{1,2,3,4,5},6))
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
