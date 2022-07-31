package main

import (
	"fmt"
	"runtime"
)

func foo() *int {
	t := 3
	return &t
}

func main() {
	x := foo()
	fmt.Println(*x)
	runtime.GC()
}
