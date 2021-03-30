package main

func main() {
	var ch = make(chan int, 2)
	ch <- 1
	ch <- 1
	ch <- 1
}
