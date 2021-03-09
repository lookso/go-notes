package main

import (
	"fmt"
)

func main() {
	var ch = make(chan int, 10)
	go product(ch)
	consumer(ch)
	ct()
}

func product(send chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Alice puts product, ID is : %d \n", i)
		send <- i
	}
	defer close(send)
}

func consumer(recv chan int) {
	//for v := range recv {
	//	fmt.Printf("Bob gets product, ID is : %d \n", v)
	//}
	ok := true
	var p int
	for ok {
		if p, ok = <-recv; ok {
			fmt.Printf("Bob gets product, ID is : %d \n", p)
		}
	}
}

func ct() {
	fmt.Println("------ct-------")
	var ch = make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch)
	}

}
