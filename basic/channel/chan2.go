package main

import (
	"fmt"
	"time"
)

func main() {
	var pool = make(chan int, 2)
	for i := 0; i < 2; i++ {
		pool <- i
	}
	//close(pool)
	fmt.Println(<-pool)
	fmt.Println(<-pool)
	//fmt.Println(<-pool) // 如果close了chan,当通道里没有值得时候获取到的都是空值,但是不会报错
	// 如果没有通过close(pool),关闭在第三次消费chan的时候会报错: chanfatal error: all goroutines are asleep - deadlock!

	// 初始化chan
	c := make(chan int)
	go func() {
		fmt.Println("go sub")
		time.Sleep(time.Second)
		c <- 1
	}()
	// 处理其他事务
	fmt.Println("main")
	// 读取chan消息
	<-c

	select {
	case e := <-time.After(time.Second):
		fmt.Println(e.Unix())
	}
}
