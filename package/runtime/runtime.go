/*
@Time : 2019-10-02 21:37 
@Author : Tenlu
@File : runtime
@Software: GoLand
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {

	fmt.Println(runtime.GOROOT())  // /usr/local/go
	fmt.Println(runtime.Version()) // go1.12.7
	fmt.Println(runtime.NumCPU())  // NumCPU返回本地机器的逻辑CPU个数
	fmt.Println("-----------------")

	for i := 0; i < 5; i++ {
		ch := make(chan int)
		go func(i int, ch chan int) {
			ch <- i
			if i == 2 {
				runtime.Goexit()
			}
			if i == 1 {
				runtime.Gosched()
			}
			fmt.Println(i)
		}(i, ch)
		<-ch
	}
	fmt.Println("当前存在的协程个数:", runtime.NumGoroutine()) // NumGoroutine返回当前存在的Go程数
}

//https://studygolang.com/articles/12598?fr=sidebar