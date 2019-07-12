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
	defer func() {
		//if err:=recover();err!=nil{
		//	fmt.Println(err)
		//}
		errMsg:=recover()
		fmt.Println(errMsg)
	}()
	var status chan int
	if status == nil {
		fmt.Println("未初始化的chan 值是nil")
	}
	status = make(chan int,3)

	fmt.Println("init value:", status)
	status <- 1
	status <- 2
	status <- 3
	status <- 4
	// 当超过容量值就回抛出错误
	close(status)

	for {
		value, ok := <-status
		if !ok {
			fmt.Println("status chan closed!")
			break
		}
		fmt.Println("status value:",value)
	}

	pingc := make(chan string, 1)
	pongc := make(chan string, 1)

	go ping(pingc, "i'm comming")
	go pong(pingc, pongc)
	fmt.Println(<-pongc)

	ch1:=make(chan int,10)
	ch2:=make(chan int,10)

	done:=make(chan bool)

	go sendChan(ch1,done)
	go reveChan(ch1,ch2,done)

	<-done
	<-done
}

func ping(pingc chan<- string, msg string) {
	pingc <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func sendChan(ch1 chan int,done chan bool)  {
	fmt.Println("------sendChan--------")

	for i:=1;i<=10;i++{
		ch1<-i
	}
	close(ch1)
	done<-true
}

func reveChan(ch1 chan int,ch2 chan int,done chan bool)  {
	fmt.Println("------receChan--------")
	for c:=range ch1 {
		ch2<-c
		fmt.Println("ch1:",c)
	}
	if _,ok:=<-ch1;!ok{
		fmt.Println("ch1 is closed")
	}
	done<-true
}

