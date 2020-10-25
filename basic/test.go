package main

import (
	"fmt"
	"sync"
)

/*
for i := 1; i <= 10; i++ {
	sum += i
	go func() {
		total += i
	}()
}
fmt.Printf("sum:%d,total:%d\n", sum, total)
上面代码有啥问题？
*/
func main() {
	var wg = new(sync.WaitGroup)
	var lc = new(sync.Mutex)
	var (
		total = 0
		sum   = 0
	)
	for i := 1; i <= 10; i++ {
		sum += i
		lc.Lock()
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer lc.Unlock()
			total += i
		}(i)
	}
	wg.Wait()
	fmt.Printf("sum:%d,total:%d\n", sum, total)

}
