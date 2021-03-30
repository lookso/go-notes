package main

import "fmt"

func main() {
	var nums = []int{-100, -50, 10, 20, -5, 30, 20, -90}
	fmt.Println(maxSubArray(nums))
}
func maxSubArray(nums []int) int {
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		nums[i] = maxA(nums[i], nums[i-1]+nums[i])
		ans = maxA(nums[i], ans)
	}
	return ans
}

func maxA(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//作者：user9779g
//链接：https://leetcode-cn.com/problems/lian-xu-zi-shu-zu-de-zui-da-he-lcof/solution/golang-ti-jie-dong-tai-gui-hua-by-user97-b0qi/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。