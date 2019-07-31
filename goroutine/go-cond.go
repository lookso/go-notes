package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)
// Go并发编程之select、锁和条件变量
// https://www.jianshu.com/p/69e098be6255

func init()  {
	fmt.Println("init1")
}
func init()  {
	fmt.Println("init2")
}
// 使用channl实现消费者生产者模型。
func producer(out chan<- int, cond *sync.Cond) {
	for {
		cond.L.Lock()
		for len(out) == 5 {
			cond.Wait()
		}
		num := rand.Intn(500)
		out <- num
		fmt.Println("----生产数据：", num)
		time.Sleep(time.Millisecond * 100)

		cond.L.Unlock()
		cond.Signal()

	}
}

func consumer(in <-chan int, cond *sync.Cond) {
	for {
		cond.L.Lock()
		if len(in) == 0 {
			cond.Wait()
		}
		num := <-in
		fmt.Println("消费数据：", num)
		time.Sleep(time.Millisecond * 200)
		cond.L.Unlock()
		cond.Signal()
	}

}

func main() {
	rand.Seed(time.Now().UnixNano())
	// 创建带缓冲channel
	ch := make(chan int, 5)

	cond := new(sync.Cond)
	cond.L = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		go producer(ch, cond)
	}
	for i := 0; i < 5; i++ {
		go consumer(ch, cond)
	}

	for true {
		runtime.GC()
	}

//	http.ListenAndServe("127,0,0,1:80",nil)
}
