package main

import (
	"fmt"
)

// 协程交替输出1A2B3C4D...26Z

func main() {
	fmt.Println('a', 'A') // 97  65
	fmt.Println('b', 'B') // 98  66
	fmt.Println('z', 'Z') // 122 90

	fmt.Println("\n-------------")
	methodA()
}
func methodA() {
	var (
		c = make(chan int)
		d = make(chan int)
		e = make(chan int)
	)
	go func() {
		for i := 1; i <= 26; i++ {
			c <- 1
			fmt.Print(i)
			d <- 1
		}
	}()
	go func() {
		for i := 'A'; i <= 'Z'; i++ {
			<-c
			<-d
			fmt.Printf("%c|", i)
		}
		e <- 1
	}()
	<-e
}
