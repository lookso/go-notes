/*
@Time : 2019/4/21 3:54 PM 
@Author : Tenlu
@File : sort
@Software: GoLand
@https://www.jishuchi.com/read/sorting-algorithm/1092
*/
package main

import (
	"fmt"
)

func main() {
	var arr = []int{2, 3, 1, 4, 5, 7, 10, 9}
	bubbleArr := bubbleSort(arr)
	fmt.Printf("bubblesort arr:%v\n", bubbleArr)

	quickArr:=quickSort(arr)
	fmt.Printf("quicksort arr:%v\n",quickArr)
}
// 冒泡排序
func bubbleSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1-i; j++ {
			if arr[j]>arr[j+1]{
				arr[j],arr[j+1] = arr[j+1],arr[j]
			}
		}
	}
	return arr
}

// 快排
func quickSort(arr []int) []int {
	return runQuick(arr, 0, len(arr)-1)
}
func runQuick(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		runQuick(arr, left, partitionIndex-1)
		runQuick(arr, partitionIndex+1, right)
	}
	return arr
}
func partition(arr []int, left, right int) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, i, index)
			index += 1
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


