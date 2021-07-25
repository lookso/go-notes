package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwLock sync.RWMutex
	c := make(chan int, 10)
	for i := 0; i < 10; i++ {
		rwLock.RLock() // 因为读写锁,读之间不互斥,可以同时加把读锁,所以输出的结果也不是有序的,如果换成互斥锁就会是有序输出
		go func(i int) {
			c <- i
			rwLock.RUnlock()
		}(i)
	}
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}
