package main

import (
	"log"
	"sync"
)

var done bool // 定义全局变量done

func read(name string, c *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done()
	c.L.Lock()
	for !done {
		// 阻塞当前协程：调用c.Wait()会使当前协程等待在sync.Cond上，
		// 直到其他协程调用sync.Cond的Signal或Broadcast方法。
		c.Wait()
	}
	log.Println(name, "starts reading")
	c.L.Unlock()
}

func write(name string, c *sync.Cond, wg *sync.WaitGroup) {
	defer wg.Done() // 确保写操作完成后减少等待组的计数
	log.Println(name, "starts writing")
	// 模拟写操作耗时
	// time.Sleep(time.Second) // 假设这里是实际的写操作
	c.L.Lock()
	done = true
	c.L.Unlock()
	log.Println(name, "wakes all")
	c.Broadcast() // 唤醒所有的协程,不要白白等待了
	// c.Signal() // 唤醒一个等待的协程
}

func main() {
	var wg sync.WaitGroup
	cond := sync.NewCond(&sync.Mutex{})
	done = false // 确保done状态正确初始化

	wg.Add(4) // 为三个读协程和一个写协程添加等待计数
	go read("reader1", cond, &wg)
	go read("reader2", cond, &wg)
	go read("reader3", cond, &wg)
	go write("writer", cond, &wg) // 修改为协程执行

	wg.Wait() // 等待所有读写协程完成
}
