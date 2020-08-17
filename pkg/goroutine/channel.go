package main

import (
	"fmt"
	"sync"
)

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
	mm()
}

func subWork(in chan int, quit chan bool) {
	for i := 0; i < 3; i++ {
		in <- i
	}
	quit <- true
}
func mm() {
	fmt.Println("-----------")
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
