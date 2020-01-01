/*
@Time : 2019-07-27 14:28 
@Author : Tenlu
@File : go-runtime
@Software: GoLand
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() {
	var ch1, ch2 = make(chan bool), make(chan bool)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
		ch1 <- true
	}()
	go func() {
		fmt.Println("hello world")
		ch2 <- true
	}()
	<-ch1
	<-ch2

	funcGoexit()
}

func funcGoexit() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		func() {
			defer println("B.defer")
			runtime.Goexit()
		}()
		println("A")
	}()

	wg.Wait()
}
