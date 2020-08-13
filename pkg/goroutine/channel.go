package main

import "fmt"

var ch8 = make(chan int, 6)

func mm() {
	for i := 0; i < 10; i++ {
		ch8 <- 8 * i
	}
	close(ch8)
}
func main() {

	go mm()
	for {
		for data := range ch8 {
			fmt.Print(data, "\t\n")
		}
		break
	}
}
