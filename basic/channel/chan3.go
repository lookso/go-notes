package main

import "fmt"

var chan1 = make(chan bool, 1)
var chan2 = make(chan bool)
var index = make(chan bool, 1)

func func1() {
	for i := 1; i < 26; i += 2 {
		// 2, 取走chan1里的元素,继续运行,打印两个数字
		<-chan1
		fmt.Print(i)
		fmt.Print(i + 1)
		// 3, 给chan2 放入一个元素,等待取走元素
		chan2 <- true
	}
}

func func2() {
	for i := 'A'; i <= 'Z'; i += 2 {
		// 4, chan2 取出元素,执行打印两个字符 ,
		<-chan2
		fmt.Print(string(i))
		fmt.Print(string(i + 1))
		// 5, chan1 接收一个元素,进入阻塞状态,等待取走元素,进入第2步,2345步一直循环直到打印完
		chan1 <- true
	}
	// 6, 结束循环,index通道接收一个元素,进入阻塞状态,等待取走
	index <- true
}

func main() {
	var c, d = make(chan int), make(chan int)
	c <- 1
	go func() {
		d <- <-c
	}()
	<-d

	fmt.Println("------------------\\n")

	// 1, chan1 里放一个值,进入阻塞状态,等待取走元素
	chan1 <- true
	go func1()
	go func2()
	// 7, index通道取走元素,继续往下执行
	<-index
	// 结果: 12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ
}
