package main

import "fmt"

/*
1.Channel原理 维护三个队列,1.数据读协程队列,2.数据写协程队列,3.数据缓冲循环队列
2.无缓冲channel和有缓冲channel,
3.select-case
4.nil、buffered、closed
5.
*/
func main() {
	var status chan int
	if status == nil {
		fmt.Println("未初始化的chan 值是nil")
	}
	status = make(chan int, 10)
	fmt.Println("init value:", status)
	status <- 1
	fmt.Println("status value:", <-status)

	pingc := make(chan string, 1)
	pongc := make(chan string, 1)


	ping(pingc, "i'm comming")
	pong(pingc, pongc)
	fmt.Println(<-pongc)
}

func ping(pingc chan<- string, msg string) {
	pingc <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}