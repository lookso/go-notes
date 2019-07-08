/*
@Time : 2019/5/11 5:09 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
	"context"
	"time"
)

/**
子协程和主线程同步有哪些处理办法？

1.使用channel机制,每个goroutine传一个channel进去然后往里写数据
再在主线程中读取这些channel,直到全部读到数据了子goroutine也就全部运行完了，
那么主goroutine也就可以结束了。这种模式是子线程去通知主线程结束。

2.使用context中cancel函数，这种模式是主线程去通知子线程结束。

3.sync.WaitGroup模式，Add方法设置等待子goroutine的数量，使用Done方法设置等待子goroutine的数量减1，当等待的数量等于0时，Wait函数返回。

4.time.sleep(),这种方式很太死板

另外 了解channel close


 */


func main()  {

	fmt.Println(123)
	fmt.Println("------goroutine-time.sleep------")
	usesleep()

	fmt.Println("------goroutine-channel-a------")
	channelA()
	fmt.Println("------goroutine-channel-b------")
	channelB()


	var c3 = make(chan int)
	go func() {
		fmt.Println("------goroutine-channel-c------")
		fmt.Print("hello world\n")
		c3<-1
	}()
	<-c3


	fmt.Println("------goroutine-context-cancel-----")
	contextCancel()
	fmt.Println("-----goroutine-waitGroup-------")
	waitGroupFunc()
	fmt.Println("-----goroutine-mutex-------")
	mutexGoFunc()


}
func usesleep()  {
	go printsleep()
	time.Sleep(time.Millisecond * 100)
}
func printsleep() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}
func channelA()  {
	var once sync.Once
	done := make(chan bool)
	for i:=0;i<10 ;i++  {
		go func(i int) {
			fmt.Println(i)
			done <- true
		}(i)
		once.Do(onlyone)  // 只输出一条hello,只会被加载过一次
	}
	for i:=0;i<10;i++ {
		<- done
	}
}
func onlyone()  {
	fmt.Println("hello")
}

func channelB()  {
	var chanTest = make(chan int)
	var chanMain = make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			chanTest <- i
			fmt.Println("product write data", i)
		}
		 close(chanTest) // 关闭写channel
		 // channel 应仅由发送方执行，而不应由接收方执行，并且在收到最后发送的值后具有关闭通道的效果。
	}()
	go func() {
		for v := range chanTest {
			fmt.Println("\t consume read data", v)
		}
		chanMain <- 666
	}()
	go func() {
		for v := range chanTest {
			fmt.Println("\t\t also consume read data", v)
		}
		chanMain <- 666
	}()
	<-chanMain
	<-chanMain // 主goroutine读取channel

}



func contextCancel()  {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	intChan := printCount(ctx)
	for n := range intChan {
		if n == 9 {
			break
		}
		fmt.Println(n)
	}
}
func printCount(ctx context.Context) chan int {
	dst := make(chan int)
	n := 0
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func waitGroupFunc()  {
	var wg sync.WaitGroup
	for i:=0;i<10 ;i++  {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("waitgroup:%d\n",i)
		}(i)
	}
	//等待所有go程结束
	wg.Wait()
	fmt.Printf("%s\n","main exit ....")
}

func mutexGoFunc()  {

	fmt.Println("---mutex---")
	var mu sync.Mutex
	mu.Lock()
	go func() {
		println("hello world-互斥锁实现协程同步")
		mu.Unlock()
	}()
	mu.Lock()
	mu.Unlock()
}
