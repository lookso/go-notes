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
		fmt.Println(1/0)
		defer wg.Done()
	}()
	fmt.Println(12345)
	wg.Wait()

}
