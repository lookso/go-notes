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
	<-signCh
}
func ctxForChanT() {
	fmt.Println("1当前运行的goroutine: ", runtime.NumGoroutine())
	signCh := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	// 任务A：doSomeThing
	go func(c context.Context) {
		for i := 0; ; i++ {
			if i == 2 { // 模拟出现报错或需要停止的情况
				fmt.Println("任务A准备停止...")
				cancel() // 发出取消指令
				break    // 停止当前任务
			}
			fmt.Println("任务A进行中：", i)
			time.Sleep(time.Second * 1)
		}
	}(ctx)

	go func(c context.Context) { // 专门用来读取停止信号
		for {
			select {
			case <-c.Done():
				fmt.Println("任务A停止。所有goroutine即将结束...")
				signCh <- struct{}{} // 发送信号，解除阻塞main线
				return
			default:
				fmt.Println("2当前运行的goroutine: ", runtime.NumGoroutine())
				time.Sleep(time.Second * 2)
			}
		}
	}(ctx)

	//停止for-select循环方式也可移步https://lan6193.blog.csdn.net/article/details/101208252
	stop := false
	for !stop {
		select {
		case <-signCh:
			stop = true
		default:
		}
	}
}
func main() {
	ctxForChanT()
	defer func() {
		fmt.Println("当前运行的goroutine-num: ", runtime.NumGoroutine())
	}()

}
