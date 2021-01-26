/*
@Time : 2019-07-18 18:39
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
)

// 常见排序算法
func main() {
	var arr = []int{4, 3, 1, 5, 7, 2, 8, 6}
	fmt.Println("-----maoPaoSort-----")
	maoPao(arr)
	fmt.Println("----selectSort---------")
	selectSort(arr)
	fmt.Println("----quickSort---------")
	fmt.Println(quickSort(arr))
	fmt.Println("----insertSort---------")
	insertSort(arr)

}

func maoPao(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	// 外层循环：确定扫描的次数
	for i := 1; i <= len(arr)-1; i++ {
		// 内层循环：一轮扫描内，两两比较，进行交换
		for j := 0; j <= len(arr)-1-i; j++ { // - i 的原因是后面的元素已经被排序过
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}

// 选择排序
// 选择排序实现思路：
//1.遍历数组/切片，找到最小数的位置
//2.最小数与第0个元素交换位置，此时最小数已归位到位
//3.从第1位开始第二次遍历，找到剩余最小的数字所在位与第1位交换
//4.重复下去，直至数组排序完成
func selectSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	vLen := len(arr)
	// 从0开始查找
	for i := 0; i < vLen; i++ {
		// 从0开始查找，那么需要找出最小值
		min := i
		// 从i右侧的所有元素中找出当前最小值所在的下标
		// {4, 3, 1, 5, 7, 2, 8, 6}
		for j := vLen - 1; j > i; j-- {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
		}
	}
	fmt.Println(arr)

}

func quickSort(values []int) []int {
	if len(values) <= 1 {
		return nil
	}
	mid := values[0]
	tail := len(values) - 1
	head := 0
	i := 1
	for head < tail {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	//  []int{4, 3, 1, 5, 7, 2, 8, 6}
	quickSort(values[:head])
	quickSort(values[head+1:])
	return values
}
func insertSort(arr []int) {
	if arr == nil || len(arr) < 2 {
		fmt.Println("数组不满足要求")
		return
	}
	// 0位置时，只有1个元素，我们认为他已经是有序的
	for i := 1; i <= len(arr)-1; i++ {
		// 若当前元素比前一个元素小，查看 i 位置的元素应该插入到 0-（i-1）位置的何处
		if arr[i] < arr[i-1] {
			temp := arr[i]
			for j := i; j >= 0; j-- {
				if j > 0 && arr[j-1] > temp {
					arr[j] = arr[j-1]
				} else {
					arr[j] = temp
					break
				}
			}
		}
	}
	fmt.Println(arr)
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
