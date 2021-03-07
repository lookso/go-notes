package main

import (
	"fmt"
	"sync"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var wg sync.WaitGroup
	go func() {
		wg.Add(1)
		go func() {
			panic("err") // 因为goroutine不确定是哪个先执行,所以也可能出现panic
		}()
		defer wg.Done()
	}()
	fmt.Println(12345)
	wg.Wait()

}
