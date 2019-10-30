package main

import "fmt"

func test() {
	//defer func() {fmt.Println("11111")}()
	//defer func() {fmt.Println("22222")}()
	//panic("system error")
	//defer func() {fmt.Println("3333")}()
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func t1() int{
	fmt.Println("1")
	return 2
}
func t2() int{
	fmt.Println("2")
	return 2
}

func main() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

	//a := 1
	//defer calc("1", a, t1())
	//a = 0
	//defer calc("2", a, t2())
	fmt.Println("---------")
	deferDemo()
	fmt.Println("---------")

	fmt.Println("---deferDemoB()--start-----")
	deferDemoB()
	fmt.Println("---deferDemoB()--end------")


}

// 需要注意的是defer执行顺序和值传递 index:1肯定是最后执行
// 但是index:1的第三个参数是一个函数,所以最先被调用calc("10", 1, 2)==>10,1,2,3
// 执行index：2时，与之前一样，需要先调用calc("20",0,2)==>20,0,2,2 执行到b=1时开始调用,
// index:2==>calc("2",0,2)==>2,0,2,2 最后执行index:1==>calc("1",1,3)==>1,1,3,4

// 10 1 2 3
// 20 0 2 2
// 2  0 2 2
// 1  1 3 4

func deferDemo()  {
	t:=1
	defer func(t int) {
		fmt.Println("aa:",t)
	}(t)
	t=2
	defer func(t int) {
		fmt.Println("bb:",t)
	}(t)
	t=3
	defer func(t int) {
		fmt.Println("cc:",t)
	}(t)
	t=4
	defer func(t int) {
		fmt.Println("dd:",t)
	}(t)
	t=5
}

func deferDemoB()  {
	t:=1
	defer func() {
		fmt.Println("aa:",t)
	}()
	t=2
	defer func() {
		fmt.Println("bb:",t)
	}()
	t=3
	defer func() {
		fmt.Println("cc:",t)
	}()
	t=4
	defer func() {
		t=200
		fmt.Println("dd:",t)
	}()
	t=5
}

