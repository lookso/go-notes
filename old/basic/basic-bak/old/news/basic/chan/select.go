/*
@Time : 2019/5/15 3:19 PM 
@Author : Tenlu
@File : select
@Software: GoLand
*/
package main

import (
	"time"
	"fmt"
)

func main()  {
	selectChannel()
	// select{}阻塞需要有线程在运行,
	// select{}阻塞main函数和for range 做定时器服务
	for i := 0; i < 20; i++ { //启动20个协程处理消息队列中的消息
		go thrind(i)
	}
	select {} // 阻塞
}
func thrind( i int){
	for range time.Tick(1000 * time.Millisecond) {
		fmt.Println("线程:",i)
	}
}
func selectChannel()  {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(4*time.Second)
		close(c)
	}()
	go func() {
		time.Sleep(3*time.Second)
		ch1 <- 3
	}()
	go func() {
		time.Sleep(3*time.Second)
		ch2 <- 5
	}()
	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <-ch1:
		fmt.Printf("ch1 case...")
	case <-ch2:
		fmt.Printf("ch1 case...")
	default:
		fmt.Printf("default go...")
	}
}

