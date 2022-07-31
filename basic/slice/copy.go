package main

import "fmt"

func main() {
	var s = []int{1, 2, 3, 4}
	var ss = make([]int, 5)
	copy(ss, s)
	fmt.Println(ss) //[1 2 3 4 0]
}
