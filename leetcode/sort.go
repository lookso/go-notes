/*
@Time : 2019-07-18 18:39 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("-----quicksort-----")
	arr := []int{5, 1, 10, 3, 6}
	sort.Ints(arr)
	fmt.Println("arr:", arr)

	QSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
	fmt.Println("----------")
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)

	quick3sortArr := Quick3Sort(arr)
	fmt.Println("quicksort3", quick3sortArr)

	fmt.Println("-----maopaosort-----")
	maoPao([]int{4, 3, 1, 5, 7, 6})
}

func maoPao(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i + 1; j < len(values)-1; j++ {
			if values[i] > values[j] {
				values[i], values[j] = values[j], values[i]
			}
		}
	}
	fmt.Println(values)

}

// quickSort 使用快速排序算法，排序整型数组
func QSort(arr []int, start int, end int) {
	var (
		flag int = arr[start]
		low  int = start
		high int = end
	)
	for {
		for low < high {
			if arr[high] < flag {
				arr[low] = arr[high]
				break
			}
			high--
		}
		for low < high {
			if arr[low] > flag {
				arr[high] = arr[low]
				break
			}
			low++
		}
		if low >= high {
			fmt.Printf("low:%d,high:%d\n", low, high)
			arr[low] = flag
			break
		}
	}
	if low-1 > start {
		QSort(arr, start, low-1)
	}
	if high+1 < end {
		QSort(arr, high+1, end)
	}
}

/**
快速排序：分治法+递归实现
随意取一个值A，将比A大的放在A的右边，比A小的放在A的左边；然后在左边的值AA中再取一个值B，将AA中比B小的值放在B的左边，将比B大的值放在B的右边。以此类推
*/
func quickSort(arr []int, firstLen, lastLen int) {
	flag := firstLen
	left := firstLen
	right := lastLen

	if firstLen >= lastLen {
		return
	}
	// 将大于arr[flag]的都放在右边，小于的，都放在左边
	for firstLen < lastLen {
		// 如果flag从左边开始，那么是必须先从有右边开始比较，也就是先在右边找比flag小的
		for firstLen < lastLen {
			if arr[lastLen] >= arr[flag] {
				lastLen--
				continue
			}
			// 交换数据
			arr[lastLen], arr[flag] = arr[flag], arr[lastLen]
			flag = lastLen
			break
		}
		for firstLen < lastLen {
			if arr[firstLen] <= arr[flag] {
				firstLen++
				continue
			}
			arr[firstLen], arr[flag] = arr[flag], arr[firstLen]
			flag = firstLen
			break
		}
	}
	quickSort(arr, left, flag-1)
	quickSort(arr, flag+1, right)
}

func Quick3Sort(values []int) []int {
	//arr := []int{5, 1, 10, 3, 6}
	if len(values) <= 1 {
		return values
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
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
	fmt.Println("vvv",values)
	Quick3Sort(values[:head])
	fmt.Println("ttt",values)
	Quick3Sort(values[head+1:])
	fmt.Println("qqq",values)

	//vvv [1 5 6 10 3]
	//ttt [1 5 6 10 3]
	//vvv [3 5 10 6]
	//ttt [3 5 10 6]
	//vvv [6 10]
	//ttt [6 10]
	//qqq [6 10]
	//qqq [3 5 6 10]
	//qqq [1 3 5 6 10]
	//quicksort3 [1 3 5 6 10]
	return values
}
