package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

var ch chan int

func work(i int) {
	fmt.Println("work" + strconv.Itoa(i))
	time.Sleep(time.Second)
	ch <- i
}

func main() {
	ch = make(chan int, 10)
	for i := 0; i < 100; i++ {
		go work(i)
		<-ch
	}
	defer fmt.Println("NumGoroutine", runtime.NumGoroutine())
}
