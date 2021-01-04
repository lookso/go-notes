package main

import "fmt"

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\n", fbnq(i))
	}
}
func fbnq(n int) int {
	if n < 2 {
		return n
	}
	return fbnq(n-2) + fbnq(n-1)
}
