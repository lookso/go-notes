package main

import (
	"fmt"
	"sync"
)
// 比较一下下面的两段程序,第一段不使用Waitgroup 需要在子协程里往channel里塞数据,主协程取channel数据
var group sync.WaitGroup

func work(in <-chan int, quit <-chan bool) {
	defer group.Done()
	for {
		select {
		case <-quit:
			fmt.Println("work quit")
			return
		case <-in:
			fmt.Println("working")
		}
	}
}
func main() {
	var quit = make(chan bool)
	var in = make(chan int)
	group.Add(1)
	go work(in, quit)
	for i := 0; i < 3; i++ {
		in <- i
	}
	quit <- true
	group.Wait()
	fmt.Println("------------------------")
	mm()
}

func subWork(in chan int, quit chan bool) {
	for i := 0; i < 3; i++ {
		in <- i
	}
	quit <- true
}
func mm() {
	var quit = make(chan bool)
	var in = make(chan int)
	go subWork(in, quit)
	for {
		select {
		case <-quit:
			fmt.Println("work quit")
			return
		case <-in:
			fmt.Println("working")
		}
	}
}
