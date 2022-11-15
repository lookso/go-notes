package main

import "fmt"

type MyTest interface {
	Add() string
}
func main() {
	var ch = make(chan int, 10)
	var ch2 = make(chan int, 20)
	for i := 1; i <= 10; i++ {
		ch2 <- i
		select {
		case ch <- i:
		case v := <-ch:
			fmt.Println(v) // 随机输出
		}
	}
	<-ch2
	fmt.Println(cap(ch2),len(ch2)) // 20 9 len()返回缓存中还没取走的元素的数量个数

}
