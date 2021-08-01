package main

import (
	"fmt"
	"runtime"
)

func main() {
	defer func() {
		fmt.Printf("cpu-num:%+v,goroutine-num:%+v", runtime.NumCPU(), runtime.NumGoroutine())
	}()
	var done = make(chan int)
	for i := 0; i < 1; i++ {
		go func(done chan int, i int) {
			done <- i
		}(done, i)
		fmt.Println(<-done)
	}
}
