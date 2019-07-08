/*
@Time : 2019/5/18 10:40 AM 
@Author : Tenlu
@File : producersAndConsumers
@Software: GoLand
*/
/**生产者和消费者模式 */
package main

import (
	"fmt"
	"sync"
)
// 向chan 发送数据,只能向chan里写数据
func producters(wg *sync.WaitGroup,factor int,c chan<- int)  {
	for i:=1; i<10;i++  {
		c<- factor * i
	}
	defer wg.Done()
}
// 接受参数,
// 只能取channel中的数据,千万不要写成chan<-
// 不然会报错:
// ./producersAndConsumers.go:24:9:
// invalid operation: range ch (receive from send-only type chan<- int)
func consumers(wg *sync.WaitGroup,ch <-chan int)  {
	for v:=range ch  {
		fmt.Println("reev:",v)
	}
	defer wg.Done()
}
func main()  {
	var wg sync.WaitGroup

	ch:=make(chan int,10)
	wg.Add(3)
	go producters(&wg,2,ch)
	go producters(&wg,3,ch)
	go consumers(&wg,ch)

	wg.Wait()
	fmt.Println("---main exit-----")

	// Ctrl+C 退出
	//sig := make(chan os.Signal, 1)
	//signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//fmt.Printf("quit (%v)\n", <-sig)
}

