package main

import "fmt"

func main() {
	var s = make([]int, 1, 1024)
	fmt.Println(len(s), cap(s), s)
	for i := 0; i < 1024; i++ {
		s = append(s, i)
	}
	fmt.Println(len(s), cap(s))
	s = append(s, []int{2, 3, 4}...)
	fmt.Println(len(s), cap(s))
	var s2 = make([]int,3000)
	fmt.Println("s2 len", copy(s2, s))
	fmt.Println(s2)

}
