package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func chanT() {
	fmt.Println("1当前运行的goroutine: ", runtime.NumGoroutine())
	fmt.Println("start")
	var ch = make(chan bool)
	go func() {
		fmt.Println("2当前运行的goroutine: ", runtime.NumGoroutine())
		ch <- true
		// 通道关闭后&读完后 range才会结束，后面的代码才会执行,不然会出现死锁
		defer close(ch)
	}()
	for value := range ch {
		fmt.Println("ch", value)
	}
	defer fmt.Println("end")
	fmt.Println("3当前运行的goroutine: ", runtime.NumGoroutine())
}

func ctxChanT() {
	signCh := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(c context.Context) {
		for {
			select {
			case <-c.Done():
				fmt.Println("任务A停止")
				signCh <- struct{}{}
				return
			default:
				for i := 0; ; i++ {
					fmt.Println("任务A：", i)
					if i == 2 {
						cancel()
						break
					}
					time.Sleep(time.Second * 1)
					fmt.Println("2当前运行的goroutine: ", runtime.NumGoroutine())
				}
			}
		}
	}(ctx)
}
func main()  {
	ctxChanT()
}