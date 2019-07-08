/*
@Time : 2019/5/29 11:14 AM 
@Author : Tenlu
@File : sort
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	numbers:=[]int{6,2,7,5,8,9}
	selectOrder:=selectSort(numbers)
	fmt.Println("selectsort:",selectOrder)

	quickSort:=quickSort(numbers)
	fmt.Println("quicksort:",quickSort)
}
// 选择排序
func selectSort(values []int) []int {
	length := len(values)
	if length <= 1 {
		return []int{}
	}
	for i := 0; i < length; i++ {
		min := i // 初始的最小值位置从0开始，依次向右
		// 从i右侧的所有元素中找出当前最小值所在的下标
		for j := length - 1; j > i; j-- {
			if values[j] < values[min] {
				min = j
			}
		}
		//fmt.Printf("i:%d min:%d\n", i, min)
		// 把每次找出来的最小值与之前的最小值做交换
		values[i], values[min] = values[min], values[i]
		//fmt.Println(values)
	}
	return values
}

// 快速排序
func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivot := arr[0]
		var less []int
		var greater []int
		for _, value := range arr[1:] {
			if value <= pivot {
				less = append(less, value)
			} else {
				greater = append(greater, value)
			}
		}
		var result []int
		result = append(result, quickSort(less)...)
		result = append(result, pivot)
		result = append(result, quickSort(greater)...)
		return result
	}
}


