package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Start print 75")
	var ch = make(chan int)
	for i := 0; i < 75; i++ {
		go func(i int, ch chan int) {
			ch <- i
		}(i, ch)
		fmt.Println(<-ch)
	}
	fmt.Println("End print 75")
	defer func() {
		fmt.Println("goroutine-num", runtime.NumGoroutine())
	}()
}
