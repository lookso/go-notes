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
	"time"
)

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

	var pipline = make(chan bool)
	var mch = make(chan bool)
	go func() {
		fmt.Println(time.Now().Format("2006-01-01 15:03:04"))
		pipline <- true
	}()

	go func() {
		<-pipline
		fmt.Println(time.Now().Format("2006-01-01 15:03:04"))
		mch <- true
	}()
	<-mch

	fmt.Println("------------")
	gosched()
}

func gosched() {
	go output("goroutine 2")
	
	runtime.Gosched()
	output("goroutine 1")
}

func output(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
	}
}
