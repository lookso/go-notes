package main

import "fmt"

func testDefer(x int) {
	defer println("a")
	defer println("b")
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。 
		}
	}()
	defer println("c")
	panic("panic error!")
}
func main() {
	testDefer(0)
	fmt.Println(foo())
}

func foo() (res int) {
	var i int=1

	defer func() {
		res++
	}()

	return i
}
