package main

import (
	"fmt"
	"sync"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err", err)
		}
	}()
	var lock sync.Mutex
	var wg sync.WaitGroup
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("err", err)
			}
		}()
		wg.Add(1)
		defer wg.Done()
		lock.Lock()
		lock.Lock()
		fmt.Println(123)
		lock.Unlock()
	}()
	fmt.Println(345435)
	wg.Wait()

	//var wg sync.WaitGroup
	//go func() {
	//	wg.Add(1)
	//	go func() {
	//		panic("err") // 因为goroutine不确定是哪个先执行,所以也可能出现panic
	//	}()
	//	defer wg.Done()
	//}()
	//fmt.Println(12345)
	//wg.Wait()
	// git 1
}
