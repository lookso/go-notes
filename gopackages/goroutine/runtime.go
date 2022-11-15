

package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 通过channel保证协程顺序输出
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
		ch <- k
		fmt.Println("NumGoroutine", runtime.NumGoroutine())
	}
	wg.Wait()
}

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println("------------------------")
	runtimeMain()
}

