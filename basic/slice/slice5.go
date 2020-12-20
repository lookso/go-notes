package main

import "fmt"

func main() {

	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		fmt.Println(&val)
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
	var a = [5]int{1, 2}
	var b = [5]int{1, 2}
	if a == b {
		fmt.Println("123")
	}
}
