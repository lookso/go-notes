/*
@Time : 2019/1/8 1:53 PM
@Author : Tenlu
@File : basic
@Software: GoLand
*/
// Go 语言支持并发，我们只需要通过 go 关键字来开启 goroutine 即可。
// Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。
// 同一个程序中的所有 goroutine 共享同一个地址空间。

package main

import (
	"fmt"
	"time"
)

//func main() {
//	fmt.Println("In main()")
//	go longWait()
//	go shortWait()
//	fmt.Println("About to sleep in main()")
//	// sleep works with a Duration in nanoseconds (ns) !
//	time.Sleep(10 * 1e9)  // sleep for 10 seconds
//	fmt.Println("At the end of main()")
//}
//
//func longWait() {
//	fmt.Println("Beginning longWait()")
//	time.Sleep(5 * 1e9) // sleep for 5 seconds
//	fmt.Println("End of longWait()")
//}
//
//func shortWait() {
//	fmt.Println("Beginning shortWait()")
//	time.Sleep(2 * 1e9) // sleep for 2 seconds
//	fmt.Println("End of shortWait()")
//}

//In main()
//About to sleep in main()
//Beginning shortWait()
//Beginning longWait()
//End of shortWait()
//End of longWait()
//At the end of main()

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}

// 执行以上代码，你会看到输出的 hello 和 world 是没有固定先后顺序

//world
//hello
//world
//hello
//hello
//world
//world
//hello
//world
//hello
