/*
@Time : 2019-07-22 08:50
@Author : Tenlu
@File : goroutine
@Software: GoLand
*/
package main

import "fmt"

/*
1. golang 协程同步的几种方式
*/
func main() {
	var done = make(chan bool)
	go goChan(done)
	if data, ok := <-done; ok {
		fmt.Println(data)
	}
	// 管道
	var chanList = make(chan int)
	var mainChanList = make(chan int)
	go send(chanList)
	go recv(chanList,mainChanList)

	for  {
		mainData,ok:=<-mainChanList
		if !ok{
			break
		}
		fmt.Println("主协程数据输出:",mainData)
	}
}
func goChan(done chan bool) {
	done <- true
	fmt.Println(123)
}

func send(chanList chan int) {
	for i := 1; i < 10; i++ {
		chanList <- i
	}
	close(chanList)  // 不close,使用range 获取chan数据,chanList就会发生阻塞
}
func recv(chanList chan int,mainChanList chan int) {
	for ch:=range chanList  {
	 	fmt.Println("子协程数据输出:",ch)
	 	mainChanList<-1
	}
	close(mainChanList)
	// 或者使用下面方式接收数据
	//for j:=1;j<10;j++{
	//	fmt.Println(<-chanList)
	//}
}
