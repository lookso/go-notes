package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
	fmt.Println("-------")
	doSelect()
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func doSelect() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		chan1 <- 1
		time.Sleep(5 * time.Second)
	}()

	go func() {
		chan2 <- 1
		time.Sleep(5 * time.Second)
	}()

	select {
	case <-chan1:
		fmt.Println("select chan1 ready.")
	case <-chan2:
		fmt.Println("select chan2 ready.")
	default:
		fmt.Println("select default")
	}

	fmt.Println("select  main exit.")
}
