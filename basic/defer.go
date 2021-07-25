package main

import "fmt"

func main() {
	fmt.Println("--------d2 start ---------")
	fmt.Println(d2())
	d3()
}

func d2() (res int) {
	var i = 1
	var j = 1
	defer func(j int) {
		res++
		j++
	}(j)
	fmt.Println("j", j)
	//res = i
	//rest++
	//return

	return i
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func d3() {
	fmt.Println("--------d3---------")
	var a int
	defer func() {
		fmt.Println("a:", a)
	}()
	a = 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
	// defer 在定义的时候会计算好调用函数的参数，所以会优先输出10、20 两个参数。然后根据定义的顺序倒序执行。
}
