/*
@Time : 2019/3/27 6:45 PM 
@Author : Tenlu
@File : go_channel
@Software: GoLand
*/
package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
// pong 函数允许通道（pings）来接收数据，另一通道（pongs）来发送数据。

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}
func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}