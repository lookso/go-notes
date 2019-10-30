package main

import "fmt"

func main() {
	var ch = make(chan int)
	var do = make(chan bool)
	go push(ch)
	go pull(ch, do)
	<-do
}
func push(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func pull(ch chan int, do chan bool) {
	for c := range ch {
		fmt.Println("c",c)
	}
	for j := range ch {
		fmt.Println("j",j) // 已经消费完,不会再有输出
	}
	//for {
	//	i, isClose := <-ch
	//	if !isClose {
	//		fmt.Println("channel closed!")
	//		break
	//	}
	//	fmt.Println(i)
	//}
	do <- true
}
