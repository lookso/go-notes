/*
@Time : 2019-07-24 07:42
@Author : Tenlu
@File : binarySearch
@Software: GoLand
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{8, 9, 6, 7, 10, 1, 5}
	sort.Ints(arr)
	findVal := 9
	key := binarySearch(arr, findVal, 0, len(arr)-1)
	if key < 0 {
		fmt.Println("no find the key")
	} else {
		fmt.Println("find key", key)
	}
	if binarySearch2(arr, 8) {
		fmt.Println("find key", key)
	} else {
		fmt.Println("no find the key")
	}
}
func binarySearch(arr []int, findVal int, start int, end int) (key int) {
	if start > end {
		return -1
	}
	mid := int(math.Round(float64(start+end) / 2))
	if arr[mid] == findVal {
		return mid
	}
	if arr[mid] > findVal {
		return binarySearch(arr, findVal, start, mid-1)
	} else {
		return binarySearch(arr, findVal, mid+1, end)
	}
}

func binarySearch2(arr []int, key int) bool {
	low := 0             // 定义最低下标为记录首位
	high := len(arr) - 1 // 定义最高下标为记录末位
	if key < arr[low] || key > arr[high] {
		return false
	}
	for low <= high {
		mid := (low + high) / 2
		mid = int(math.Round(float64(mid)))
		guess := arr[mid]
		if guess == key {
			return true
		}
		if guess < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// 对半查找,时间复杂度是O(logn)
