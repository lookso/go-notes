/*
@Time : 2019-10-28 21:42 
@Author : Tenlu
@File : square
@Software: GoLand
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(squareRoot(10))
	fmt.Println(mySqrt(9))
	fmt.Println(sprt(9))
	fmt.Println(maopao())
}

func getabs(x float64) float64 {

	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}
	return x
}

// 牛顿迭代法
func squareRoot(x float64) float64 {
	z := 1.0

	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		for getabs(z*z-x) > 0.000000000001 {
			z = (z + x/z) / 2
		}
		return z
	}
}

func mySqrt(x int) int {
	// 二分查找，定位平方根
	i, j := 1, x
	for i <= j {
		mid := (i + j) / 2
		fmt.Println(mid)
		if mid*mid > x { // 去左区间找
			j = mid - 1
		} else if mid*mid < x { // 去右区间找
			i = mid + 1
		} else { // 找到了
			return mid
		}
	}
	return i - 1
}

func sprt(x int) int {
	fmt.Println("------")
	for i := 0; i < x; i++ {
		for j := i; j < x; j++ {
			if i*j == x {
				return j
			}
		}
	}
	return 0
}

func maopao() []int {
	var nums = []int{4, 5, 3, 2, 1, 7, 9, 4, 10}
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		for j := i; j < numsLen; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}
