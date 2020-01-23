package main

import "fmt"

func main() {
	s := []int{1, 2}
	add(s)
	fmt.Println(s)

	b3 := true
	a3 := true
	if a3 && b3{
		fmt.Println(a3,b3)
	}
}
func add(s []int) {
	s[0] = 3
	fmt.Println(s)
}
