package main

import "fmt"

func main() {
	s := " "
	fmt.Println("s-len", len(s))
	if s == " " {
		fmt.Println("s is empty")
	}

	n := BinaryLeftPlus([]int{1, 2})
	fmt.Println(n)
}

func BinaryLeftPlus(nums []int) int {
	var res int
	for _, num := range nums {
		if num > 0 {
			res += 1 << (num - 1)
		}
	}
	return res
}
