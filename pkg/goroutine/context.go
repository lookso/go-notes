package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int, 5)
	wg.Add(2)
	go send(ctx, ch)
	go receive(ctx, ch)
	time.Sleep(1 * time.Second)
	cancel()
	wg.Wait()
	fmt.Println("stop")
}

func send(ctx context.Context, ch chan int) {
	defer wg.Done()
	i := 0
	for {
		select {
		case v := <-ctx.Done(): // 判断父协程是否退出
			fmt.Println("send Done", v)
			return
		default:
			if i >= 10 {
				fmt.Println("send finished")
				return
			}
			ch <- i
			fmt.Println("send", i)
		}
		i++
		time.Sleep(200 * time.Millisecond)
	}
}

func receive(ctx context.Context, ch chan int) {
	defer wg.Done()
	for {
		select {
		case v := <-ctx.Done():
			fmt.Println("receive Done", v)
			return
		case v := <-ch:
			fmt.Println("receive", v)
		default:
			fmt.Println("receive wait")
			time.Sleep(1 * time.Second)
		}
	}
}
