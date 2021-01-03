/*
@Time : 2019-02-21 21:56 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"math/rand"
)

func read(ch1 chan int, ch2 chan bool) {
	for {
		v, ok := <-ch1
		if ok {
			fmt.Printf("read a int is %d\n", v)
		} else {
			ch2 <- true
		}
	}

}

func write(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func SendInt(ch chan<- int) {
	ch <- rand.Intn(1000)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan bool)

	go write(ch1)
	go read(ch1, ch2)

	<-ch2

	// FIFO队列,先发送的chan一定最先被接收
	ch3 := make(chan int, 3)
	ch3 <- 2
	ch3 <- 1
	ch3 <- 3
	elem1 := <-ch3
	elem2 := <-ch3

	fmt.Printf("The first element received from channel ch1: %v\n",
		elem1)
	fmt.Printf("The two element received from channel ch1: %v\n",
		elem2)

	var c = make(chan int, 5)

	for i := 0; i < 5; i++ {
		c <- i
	}
	for elem := range c {
		fmt.Println(elem)
	}
	close(c)


	print("----------chan start-----------\n")
	ch := make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	close(ch)
	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("channel closed!")
			break
		}
		fmt.Println(i)
	}
	print("----------chan end-----------\n")
	

}
