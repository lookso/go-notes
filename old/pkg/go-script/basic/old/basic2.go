/*
@Time : 2019/3/18 4:07 PM 
@Author : Tenlu
@File : basic2
@Software: GoLand
*/
package main

import "fmt"

func main() {
	var a, b, c int
	a = 1
	fmt.Println(a, b, c)

	nextInt := intSeq()
	// 通过多次调用 nextInt 来看看闭包的效果。

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// 为了确认这个状态对于这个特定的函数是唯一的，我们重新创建并测试一下。

	newInts := intSeq()
	fmt.Println(newInts())


	f("direct")
	// 使用 go f(s) 在一个 Go 协程中调用这个函数。这个新的 Go 协程将会并行的执行这个函数调用。

	go f("goroutine")
	// 你也可以为匿名函数启动一个 Go 协程。

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	// 现在这两个 Go 协程在独立的 Go 协程中异步的运行，所以我们需要等它们执行结束。这里的 Scanln 代码需要我们在程序退出前按下任意键结束。

	var input string
	fmt.Scanln(&input)
	if(input=="12"){
		print("23")
	}
	fmt.Println("done")

}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}