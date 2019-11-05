/*
@Time : 2019-05-13 22:58 
@Author : Tenlu
@File : atomic
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var nums uint64
func incr(wg *sync.WaitGroup) {
	// AddUint64原子性的将val的值添加到*addr并返回新值。
	atomic.AddUint64(&nums, 1)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)  //开启协程
		go incr(&wg)
	}
	wg.Wait()
	fmt.Println("nums =", nums)
	// LoadUint64原子性的获取*addr的值。
	fmt.Println("nums =",atomic.LoadUint64(&nums))
}