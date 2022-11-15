package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup
	var num int64
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			atomic.AddInt64(&num, 10)
			defer wg.Done()
		}()
	}
	wg.Wait()
	wg.Add(10)
	for i := 1; i <= 10; i++ {
		go func() {
			atomic.AddInt64(&num, 10)
			defer wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("num", atomic.LoadInt64(&num))
}
