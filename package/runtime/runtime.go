/*
@Time : 2019-10-02 21:37 
@Author : Tenlu
@File : runtime
@Software: GoLand
*/

package main

import "fmt"

//runtime包中有几个处理goroutine的函数：
//Goexit
//退出当前执行的goroutine，但是defer函数还会继续调用
//
//Gosched
//让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行。
//
//NumCPU
//返回 CPU 核数量
//
//NumGoroutine
//返回正在执行和排队的任务总数
//
//GOMAXPROCS
//用来设置可以并行计算的CPU核数的最大值，并返回之前的值。

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
