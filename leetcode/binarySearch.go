/*
@Time : 2019-07-24 07:42 
@Author : Tenlu
@File : binarySearch
@Software: GoLand
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{8, 9, 6, 7, 10, 7, 1}
	sort.Ints(arr)
	findVal := 9
	key := binarySearch(arr, findVal, 0, len(arr)-1)
	fmt.Println("find key", key)
}
func binarySearch(arr []int, findVal int, start int, end int) (key int) {
	if start > end {
		return -1
	}
	mid := (start + end) / 2
	if arr[mid] == findVal {
		return mid
	}
	if arr[mid] > findVal {
		return binarySearch(arr, findVal, start, mid-1)
	} else {
		return binarySearch(arr, findVal, mid+1, end)

	}
}
