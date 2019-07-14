package main

import "fmt"

func main() {
	defer d1()
	defer d2()
	defer d3()
	deferFuncParameter()
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

func deferFuncParameter() int{
	var aInt = 1

	defer fmt.Println(aInt)
	aInt = 2
	return 10
}
