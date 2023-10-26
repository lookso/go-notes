package main

import (
	"fmt"
)

func main() {
	//fmt.Println(php2go.InArray([]string{"1", "2"}, 1))
	//fmt.Println(php2go.ArrayMerge([]interface{}{"1", "2"}, []interface{}{"3", "4"}))

	var arr=[]int{1,2,3}
	var list=[]int{3,4,5}
	arr=list
	fmt.Println(arr)

}
