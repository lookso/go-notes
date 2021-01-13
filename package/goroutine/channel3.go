package main

import "fmt"

func main() {
	//创建一个channel
	c := make(chan int)
	cc := make(chan int)
	go func() {
		fmt.Println(123)
		c <- 1
	}()
	go func() {
		fmt.Println(789)
		<-cc
	}()
	<-c
	cc <- 1
	fmt.Println(456)

}
