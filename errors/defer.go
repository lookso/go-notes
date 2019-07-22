package main

import (
	"fmt"
	"os"
)

type number int

func (n number) print() {
	fmt.Println("print:", n)
}
func (n *number) pprint() {
	fmt.Println("pprint:", *n)
}

func main() {

	var n number
	defer n.print()  //0
	defer n.pprint() // 3
	defer func() {
		n.print() // 3
	}()
	n = 3

	var whatever [3]struct{}
	for i := range whatever {
		defer func() {
			fmt.Println(i)
		}()
	}

	defer func() {
		if p := recover(); p != nil {
			fmt.Println(p)
		}
	}()
	errHandle()
	defer d1()
	defer d2()
	defer d3()
	deferFuncParameter()
	fmt.Println("test:", test())
}

func d1() {
	fmt.Println("this is d1")
}
func d2() {
	fmt.Println("this is d2")
}
func d3() int {

	defer func() {
		fmt.Println("this is d3 return1")
	}()
	defer func() {
		fmt.Println("this is d3 return2")
	}()
	fmt.Println("this is d3")
	return 10
}

func deferFuncParameter() int {
	var aInt = 1

	defer fmt.Println(aInt)
	aInt = 2
	return 10
}
func errHandle() {
	f, err := os.Open("defer.txt")
	if err != nil {
		fmt.Println("file is not exists")
	}
	defer f.Close()
}

func test() (r int) { // 12

	t := 5
	defer func() {
		t = t + 5
		r = 12
	}()
	return t
}
