package main

import "fmt"

func main() {
	var ch = make(chan int)
	go func() {
		fmt.Println(<-ch)
	}()
	ch <- 1

	var c = make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}
