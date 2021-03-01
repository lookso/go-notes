package main

import "fmt"

func main() {
	//fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets3([]int{1, 2, 3}))
}
func subsets3(nums []int) [][]int {
	arr := make([][]int, 0)
	arr = append(arr, []int{})
	for i := 0; i < len(nums); i++ {
		tempArr := make([][]int, 0)
		for _, c := range arr {
			temp := make([]int, 0)
			temp = append(temp, c...)

			temp = append(temp, nums[i])
			tempArr = append(tempArr, temp)
		}
		fmt.Println(tempArr)
		for _, c := range tempArr {
			arr = append(arr, c)
		}
	}
	return arr
}

func subsets(nums []int) [][]int {
	res := [][]int{}

	var dfs func(i int, list []int)
	dfs = func(i int, list []int) {
		if i == len(nums) {
			tmp := make([]int, len(list))
			copy(tmp, list)
			res = append(res, tmp)
			return
		}
		list = append(list, nums[i])
		dfs(i+1, list)
		list = list[:len(list)-1]
		dfs(i+1, list)
	}
	dfs(0, []int{})

	return res
}

func subsets2(nums []int) [][]int {
	var ans [][]int
	ans = append(ans, []int{})

	for _, value := range nums {

		for _, old := range ans {
			var new = make([]int, len(old), (cap(old))+1)
			copy(new, old)
			fmt.Println(ans)
			ans = append(ans, append(new, value))
		}

	}

	return ans
}

// https://leetcode-cn.com/problems/subsets/solution/shou-hua-tu-jie-zi-ji-hui-su-fa-xiang-jie-wei-yun-/
