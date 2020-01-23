/*
@Time : 2019/2/26 5:38 PM 
@Author : Tenlu
@File : once
@Software: GoLand
*/
package main

import (
	"sync"
	"fmt"
	"time"
)

func main() {
	fmt.Println("--------Once()--------")
	goOnce()
	fmt.Println("--------Mutex()-------")
	goMutex()

}

func goOnce() {
	// Once是只执行一次动作的对象
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	for i := 0; i < 10; i++ {
		// 如果once.Do(f)被多次调用，只有第一次调用会执行f，即使f每次调用Do 提供的f值不同。
		// 需要给每个要执行仅一次的函数都建立一个Once类型的实例。
		// Do用于必须刚好运行一次的初始化。因为f是没有参数的，因此可能需要使用闭包来提供给Do方法调用
		// 因为只有f返回后Do方法才会返回，f若引起了Do的调用，会导致死锁。
		once.Do(onceBody) // func (o *Once) Do(f func()) Do方法当且仅当第一次被调用时才执行函数f。
	}
}

// 互斥锁
//1.Mutex 为互斥锁，Lock() 加锁，Unlock() 解锁
//2.在一个 goroutine 获得 Mutex 后,其他 goroutine 只能等到这个 goroutine 释放该 Mutex
//3.使用 Lock() 加锁后，不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
//4.在 Lock() 之前使用 Unlock() 会导致 panic 异常
//5.已经锁定的 Mutex 并不与特定的 goroutine 相关联，这样可以利用一个 goroutine 对其加锁，再利用其他 goroutine 对其解锁
//6.在同一个 goroutine 中的 Mutex 解锁之前再次进行加锁，会导致死锁
//7.适用于读写不确定，并且只有一个读或者写的场景

func goMutex() {
	var mutex sync.Mutex // 字段含义为零值为解锁状态
	fmt.Println("Lock the lock")
	mutex.Lock()
	fmt.Println("The lock is locked")
	channels := make([]chan int, 4)
	for i := 0; i < 4; i++ {
		channels[i] = make(chan int)
		go func(i int, c chan int) {
			fmt.Println("Not lock: ", i)
			mutex.Lock()
			fmt.Println("Locked: ", i)
			time.Sleep(time.Second)
			fmt.Println("Unlock the lock: ", i)
			mutex.Unlock()
			c <- i
		}(i, channels[i])
	}
	time.Sleep(time.Second)
	fmt.Println("Unlock the lock")
	mutex.Unlock()
	time.Sleep(time.Second)

	for _, c := range channels {
		<-c
	}
}


