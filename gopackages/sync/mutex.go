package main

import (
	"fmt"
	"sync"
	"time"
)

// go mutex是互斥锁，只有Lock和Unlock两个方法，在这两个方法之间的代码不能被多个goroutins同时调用到。
// 在同一个协程中加锁后，不能再继续对其加锁，否则会panic。只有在解锁之后才能再次加锁。

var num = 0
func increment(wg *sync.WaitGroup,lock *sync.Mutex) {
	lock.Lock()  //通过加互斥锁处理,现在我们可以对上述程序加上锁，每次只能由一个线程来操作num的值
	num = num + 1
	lock.Unlock()
	wg.Done()
}

func mutexFunc()  {
	var w sync.WaitGroup
	var lock sync.Mutex
	// 调用1000个协程来进行num=num+1
	// 如果不加互斥锁,每次运行都没有达到预期的效果,运行几次的输出分别为
	//num = 971
	//num = 944
	//num = 959
	// 因为多个并发的协程试图访问 num 的值,这时就会发生竞态条件。
	//
	for i := 0; i < 1000; i++ {
		w.Add(1)        //开启协程
		go increment(&w,&lock)
	}
	w.Wait()
	fmt.Println("num =", num)
}

var count = 0
func increment2(wg *sync.WaitGroup, b chan bool) {
	b <- true
	count = count + 1
	<-b
	wg.Done()
}

func channelMutexFunc()  {
	var w sync.WaitGroup
	ch := make(chan bool,1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()
	fmt.Println("count =", count)
}

func mutexTest()  {
	wa := sync.WaitGroup{}

	var mu sync.Mutex
	fmt.Println("加锁0")
	mu.Lock()

	fmt.Printf("上锁中0\t")
	for i := 1; i < 4; i++ {
		wa.Add(1)
		go func(i int) {
			fmt.Printf("加锁%d\t", i)
			mu.Lock()
			fmt.Printf("上锁中%d\t", i)
			time.Sleep(time.Second * 1)
			mu.Unlock()
			fmt.Printf("解锁%d\t\n", i)
			wa.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	mu.Unlock()
	fmt.Println("\n解锁0")
	wa.Wait()
}
func main() {
	fmt.Println("----sync.mutex 互斥锁-----")
	mutexFunc() // 互斥锁
	fmt.Println("----通过channel自定义实现互斥锁-----")
	channelMutexFunc()
	fmt.Println("----lock-test-----")
	mutexTest()
}