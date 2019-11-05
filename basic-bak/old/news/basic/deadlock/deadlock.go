/*
@Time : 2019/5/18 4:53 PM 
@Author : Tenlu
@File : deadlock
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	// go 语言里的 channel 只有推送和读取同时存在时才不会发生死锁，只读或只推都会发生死锁。


	// 第一种:因为给 ch 推送消息而没有读取，这时会阻塞当前协程，如果是 main，那么编译就报错 deadlock
	/*
	ch:=make(chan int)
	ch<-1
	fmt.Println(" ----- here ----- ")
	*/

	//fatal error: all goroutines are asleep - deadlock!
	//
	//	goroutine 1 [chan send]:
	//main.main()
	///Users/itech8/data/app/my_app/go/src/go_script/news/basic/deadlock/deadlock.go:13 +0x54
	//exit status 2

	// 第二种:因为读取 ch1 时 ch1 又在读取 ch2，而 ch2 并没有推送消息，故，死锁。解决办法就是放开注释那行代码就OK
	/*
	var ch1=make(chan int)
	var ch2=make(chan int)


	fmt.Println("111")
	go func(s string) {
		fmt.Println(s)
		ch1<-<-ch2
	}("deadlock2")
 	// ch2<-1
	fmt.Println("222")
	<-ch1
	*/

	//deadlock2
	//fatal error: all goroutines are asleep - deadlock!
	//
	//	goroutine 1 [chan receive]:
	//main.main()
	///Users/itech8/data/app/my_app/go/src/go_script/news/basic/deadlock/deadlock.go:31 +0xb5
	//
	//goroutine 5 [chan receive]:
	//main.main.func1(0xc000062060, 0xc0000620c0, 0x10c27cd, 0x9)
	///Users/itech8/data/app/my_app/go/src/go_script/news/basic/deadlock/deadlock.go:29 +0xaf
	//created by main.main
	///Users/itech8/data/app/my_app/go/src/go_script/news/basic/deadlock/deadlock.go:27 +0x9e
	//exit status 2


	// 第三种是因为读 quit 时被 c <- 1 这行代码阻塞了，c 并没有被读取故宿主被阻塞发生死锁。解决办法就是放开注释行代码就OK
	/*
	c,quit:=make(chan int),make(chan int)

	go func() {
		c<-1
		quit<-2
	}()
	// <-c
	n:=<-quit
	fmt.Println(n)
	*/

	//fatal error: all goroutines are asleep - deadlock!
	//
	//	goroutine 1 [chan receive]:
	//main.main()
	///Users/itech8/data/app/my_app/go/src/go_script/news/basic/deadlock/deadlock.go:58 +0xaf
	//
	//goroutine 18 [chan send]:
	//main.main.func1(0xc00008c000, 0xc00008c060)
	///Users/itech8/data/app/my_app/go/src/go_script/news/basic/deadlock/deadlock.go:54 +0x37
	//created by main.main
	///Users/itech8/data/app/my_app/go/src/go_script/news/basic/deadlock/deadlock.go:53 +0x8e
	//exit status 2

	// 第4种，chan cap 容量不足
	var ch=make(chan int,1) // cap变大点
	ch<-10
	ch<-12

	close(ch)
	for i:=0;i<2;i++  {
		fmt.Println(<-ch)
	}

}

