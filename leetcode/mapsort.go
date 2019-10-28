/*
@Time : 2019-10-26 08:46 
@Author : Tenlu
@File : mapsort
@Software: GoLand
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	mapSort()
}

func mapSort() {
	var mapD = map[int]string{5: "beijing", 2: "china", 3: "money", 4: "fuck"}
	var arrD = make([]int, 0)

	for k, _ := range mapD {
		arrD = append(arrD, k)
	}
	sort.Ints(arrD)
	for _, v := range arrD {
		fmt.Println(v, mapD[v])
	}
}
