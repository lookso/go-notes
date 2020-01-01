//select是Go中的一个控制结构，类似于用于通信的switch语句。每个case必须是一个通信操作，要么是发送要么是接收。
//select随机执行一个可运行的case。如果没有case可运行，它将阻塞，直到有case可运行。一个默认的子句应该总是可运行的。
//以下描述了 select 语句的语法：
//
//每个case都必须是一个通信
//所有channel表达式都会被求值
//所有被发送的表达式都会被求值
//如果任意某个通信可以进行，它就执行；其他被忽略。
//如果有多个case都可以运行，Select会随机公平地选出一个执行。其他不会执行。
//否则：
//如果有default子句，则执行该语句。
//如果没有default字句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

package main

import (
	"fmt"
	"time"
)

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server2"
}
func server3(ch chan string) {
	ch <- "from server3"
}

func server4(ch chan string) {
	ch <- "from server4"
}

func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
		//default:
		//    fmt.Println("no value received")
	}

	// 程序运行到了select 语句。select 会一直发生阻塞，除非其中有 case  准备就绪。
	// 在上述程序里,server1 协程会在 6 秒之后写入 output1 信道,而 server2  协程在3秒之后就写入了output2 信道。
	// 因此 select 语句会阻塞 3 秒钟,等着 server2  向 output2  信道写入数据。
	// 3 秒钟过后，程序会输出： from server2

	//  如果 select 只含有值为 nil 的信道，也同样会执行默认情况
	var cch chan string
	select {
	case v := <-cch:
		fmt.Println("received value", v)
	default:
		fmt.Println("default case executed")
	}
	//在上面程序中，ch 等于 nil，而我们试图在 select 中读取 ch（第 8 行）。如果没有默认情况，select 会一直阻塞，导致死锁。
	//由于我们在 select内部加入了默认情况，程序会执行它

	// 当 select 由多个 case 准备就绪时，将会随机地选取其中之一去执行。
	// 因此这个 select 语句中的两种情况都准备好执行了。如果你运行这个程序很多次的话，输出会是 from server3 或者 from server4,这会根据随机选取的结果而变化。
	op1 := make(chan string)
	op2 := make(chan string)
	go server3(op1)
	go server4(op2)
	time.Sleep(1 * time.Second)
	select {
	case s3 := <-op1:
		fmt.Println(s3)
	case s4 := <-op2:
		fmt.Println(s4)
	default:
		fmt.Println("no value received")
	}
	// select {}
	// 除非有 case 执行，select 语句就会一直阻塞着。在这里，select 语句没有任何 case，因此它会一直阻塞，导致死锁。该程序会触发 panic,输出如下:

}
