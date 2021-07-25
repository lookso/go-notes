package main

import (
	"fmt"
	"sync"
)

func main() {
	// 互斥锁金和channel+waitgroup都可以保证顺序输出
	var mutex sync.Mutex
	var group sync.WaitGroup
	mutex.Lock()
	fmt.Println("main locked")
	for i := 1; i <= 10; i++ {
		group.Add(1)
		go func(wg *sync.WaitGroup, k int) {
			mutex.Lock()
			fmt.Println("sub locked")
			fmt.Println(k)
			mutex.Unlock()
			fmt.Println("sub unlocked")
			defer wg.Done()
		}(&group, i) // 此处需要使用指针,确保是同一个waitGroup
	}
	mutex.Unlock()
	fmt.Println("main unlocked")
	group.Wait()
	fmt.Println("--------------")
	channel()
}

func channel()  {
	c := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-c
	}
}