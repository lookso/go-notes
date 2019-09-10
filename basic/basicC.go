package main

import "fmt"

func main() {

	println(DeferFunc1(1)) // 4
	println(DeferFunc2(1)) // 1
	println(DeferFunc3(1)) // 3
}

func DeferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		fmt.Println(t)
		t += 3
	}()
	fmt.Println(t)
	return t
}

func DeferFunc3(i int) (t int) {
	defer func() {
		t += i
	}()
	return 2
}
