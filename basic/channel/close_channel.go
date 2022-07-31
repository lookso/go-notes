package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		fmt.Println("goroutine-num", runtime.NumGoroutine())
	}()
	const NumSenders = 50
	dataCh := make(chan int)
	stopCh := make(chan struct{})
	defer close(dataCh)
	// senders
	for i := 0; i < NumSenders; i++ {
		go func(i int) {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- i:
				}
			}
		}(i)
	}
	// the receiver
	for value := range dataCh {
		if value == 10 {
			fmt.Println("send stop signal to senders.")
			close(stopCh)
			break
		}
		fmt.Println(value)
	}
	select {
	case <-time.After(time.Second * 6):
		fmt.Println("5s after....")
	}

}
