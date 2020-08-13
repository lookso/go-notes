package main

import (
	"fmt"
	"sync"
)

func runtimeMain() {
	var wg sync.WaitGroup
	var ch = make(chan int)
	for k := 1; k <= 10; k++ {
		wg.Add(1)
		go func(k int) {
			<-ch
			fmt.Println(k)
			defer wg.Done()
		}(k)
		//fmt.Println("NumGoroutine", runtime.NumGoroutine())
		ch <- k
	}

	wg.Wait()
}

func main() {
	runtimeMain()
}
