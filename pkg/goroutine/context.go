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
	fmt.Println("-------------------------")
	mc()
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
//----------------------------------------
var swg sync.WaitGroup
func gen(ctx context.Context) <-chan int {
	// 创建子context
	subCtx, _ := context.WithCancel(ctx)
	go sub(subCtx)  // 这里使用ctx，也能给goroutine通知
	dst := make(chan int)
	n := 1
	go func() {
		defer swg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("end")
				return // return，防止goroutine泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func sub(ctx context.Context) {
	defer swg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end too")
			return // returning not to leak the goroutine
		default:
			fmt.Println("test")
		}
	}
}

func mc() {
	swg.Add(2)
	ctx, cancel := context.WithCancel(context.Background())
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	cancel()
	swg.Wait()
}