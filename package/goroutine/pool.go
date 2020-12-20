package main

import (
	"runtime"
	"sync"
	"time"
)

type Pool struct {
	queue chan int
	wg    *sync.WaitGroup
}
func NewPool(size int) *Pool {
	if size <= 0 {
		size = 1
	}
	return &Pool{
		queue: make(chan int, size),
		wg:    new(sync.WaitGroup),
	}
}
func (p *Pool) Add(delta int) {
	for i := 0; i < delta; i++ {
		p.queue <- 1
	}
	p.wg.Add(delta)
}

func (p *Pool) Done() {
	<-p.queue
	p.wg.Done()
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
// 实现协程池
func main() {
	pool := NewPool(10)
	println("main before NumGoroutine ", runtime.NumGoroutine())
	for k := 0; k < 50; k++ {
		pool.Add(1)
		go func() {
			time.Sleep(time.Second)
			println("sub NumGoroutine ", runtime.NumGoroutine())
			pool.Done()
		}()
	}
	pool.Wait()
	defer println("main defer NumGoroutine ", runtime.NumGoroutine())
}

//Go语言-并发模式-goroutine池实例（work）
//https://www.cnblogs.com/limaosheng/p/11070819.html