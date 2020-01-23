/*
@Time : 2019-10-29 13:23
@Author : Tenlu
@File : consumer
@Software: GoLand
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	//  找出数组里面3个数的和等于指定值。
	var nums = []int{2, 2, 6, 1, 4, 5, -1, -1}
	getThreeNums(nums, 10)

}

func getThreeNums(array []int, target int) {
	sort.Ints(array)
	var result [][3]int
	for i := 0; i < len(array)-2; i++ {
		fmt.Println(array[i])
		if i > 0 && array[i-1] == array[i] {
			continue
		}
		j := i + 1
		k := len(array) - 1

		for j < k {
			if array[i]+array[j]+array[k] > target {
				k--
				//for array[k] == array[k+1] && j < k {
				//	k--
				//}
			} else if array[i]+array[j]+array[k] < target {
				j++
				//for array[j] == array[j-1] && j < k {
				//	j++
				//}
			} else {
				result = append(result, [3]int{array[i], array[j], array[k]})
				j++
				k--
				//for array[k] == array[k+1] && j < k {
				//	k--
				//}
				//for array[j] == array[j-1] && j < k {
				//	j++
				//}
			}
		}
	}
	fmt.Println(result)

}
