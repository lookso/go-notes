package main

import (
	"sync"
	"fmt"
	"time"
)

var locker = new(sync.Mutex)
var cond = sync.NewCond(locker)


func test(x int) {
	cond.L.Lock() //获取锁
	fmt.Println("aaa: ", x)
	// Wait添加一个计数，也就是添加一个阻塞的goroutine
	cond.Wait()//等待通知  暂时阻塞
	fmt.Println("bbb: ", x)
	time.Sleep(time.Second * 2)
	cond.L.Unlock()//释放锁
}


func main() {
	condFunA()
	condFuncB()
}
func condFunA()  {
	for i := 0; i < 5; i++ {
		go test(i)
	}
	fmt.Println("start all")
	time.Sleep(time.Second * 1)
	fmt.Println("broadcast")
	cond.Signal()// 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	// Signal解除一个goroutine的阻塞，计数减一。
	cond.Signal()// 3秒之后 下发一个通知给已经获取锁的goroutine
	time.Sleep(time.Second * 1)
	// Broadcast唤醒所有等待cond的线程。调用者在调用本方法时，建议（但并非必须）保持c.L的锁定。
	cond.Broadcast()//3秒之后 下发广播给所有等待的goroutine
	time.Sleep(time.Second * 10)
	fmt.Println("finish all")
}

func condFuncB()  {
	// 注意：在调用Signal，Broadcast之前，应确保目标进入Wait阻塞状态。
	// 可以看出来，每执行一次Signal就会执行一个goroutine。
	// 如果想让所有的goroutine执行，那么将所有的Signal换成一个Broadcast方法可以。
	/*
	Cond在Locker的基础上增加的一个消息通知的功能。但是它只能按照顺序去使一个goroutine解除阻塞。

	Cond有三个方法：Wait，Signal，Broadcast。
	Wait添加一个计数，也就是添加一个阻塞的goroutine。
	Signal解除一个goroutine的阻塞，计数减一。
	Broadcast接触所有wait goroutine的阻塞。

	那外部传入的Locker，是对wait，Signal，Broadcast进行保护。防止发送信号的时候，不会有新的goroutine进入wait。
	在wait逻辑完成前，不会有新的事件发生。
	*/


	wait := sync.WaitGroup{}
	locker := new(sync.Mutex)
	cond := sync.NewCond(locker)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wait.Done()
			wait.Add(1)
			cond.L.Lock()
			fmt.Println("Waiting start...")
			cond.Wait()
			fmt.Println("Waiting end...")
			cond.L.Unlock()

			fmt.Println("Goroutine run. Number:", i)
		}(i)
	}

	time.Sleep(2e9)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(2e9)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	time.Sleep(2e9)
	cond.L.Lock()
	cond.Signal()
	cond.L.Unlock()

	wait.Wait()
}