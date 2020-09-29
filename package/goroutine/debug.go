package main

import (
	"context"
	"fmt"
	"runtime"
)

func sayHello(ctx context.Context, done chan bool) {
	fmt.Println("Hello goroutine")
	select {
	case <-ctx.Done():
		fmt.Println("主 goroutine 结束")
		return
	default:
		done <- true
		fmt.Println("监控 goroutine 中...")
	}
}

func main() {

	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()
	var done = make(chan bool, 10)
	ctx, cancel := context.WithCancel(context.Background())
	go sayHello(ctx, done)
	<-done
	cancel()
	fmt.Println("Hello main")
}
