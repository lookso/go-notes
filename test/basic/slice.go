package main

import "fmt"

var oldNames = make([]int, 1)
var names = []int{1, 2, 3, 4, 5, 7}

type MyName struct {
}

func main() {
	fmt.Println("oldNames", len(oldNames), cap(oldNames))
	newNames := append(oldNames, names...)
	fmt.Println("names", len(names), cap(names))
	fmt.Println("newNames", len(newNames), cap(newNames))

	x := [3]int{1, 2, 3}

	func(arr [3]int) {
		arr[0] = 7
		fmt.Println(arr)
	}(x)

	fmt.Println(x)

}
