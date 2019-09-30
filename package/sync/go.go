package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)



func endWork(fromFunction string, ch chan int) {
	defer func() { fmt.Println(fromFunction, "sleepRandom complete") }()
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	randomNumber := r.Intn(100)
	sleepTime := randomNumber + 100
	if ch != nil {
		ch <- sleepTime
	}
}
func nextWork(ctx context.Context) {
	var ch = make(chan bool, 10)
	defer func() {
		fmt.Println("sleepRandomContext complete")
		ch <- true
	}()
//	go endWork("nextWork",chan int)

	select {
	case <-ctx.Done():
		fmt.Println("sleepRandomContext: Time to return")
	default:
		fmt.Println("123")
	}
}
func doWork(ctx context.Context) {
	cancelTimeOut, cancelFunc := context.WithTimeout(ctx, time.Duration(100)*time.Millisecond)
	ch := make(chan bool)
	select {
	case <-ctx.Done():
		fmt.Println("exit")
	case <-ch:
		fmt.Println("exit")
	}
	go nextWork(cancelTimeOut)
	defer cancelFunc()
}
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	doWork(ctx)
	time.Sleep(time.Second * 4)

}
