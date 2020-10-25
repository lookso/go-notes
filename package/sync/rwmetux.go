package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var Password = secret{password: "myPassword"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

// 通过rwmutex读
func rwMutexShow(c *secret) string {
	c.RWM.RLock()
	fmt.Println("show with rwmutex", time.Now().Second())
	defer c.RWM.RUnlock()
	return c.password
}

func tf() {
	// 定义一个稍后用于覆盖(重写)的函数
	var show = func(c *secret) string { return "" }
	show = rwMutexShow
	fmt.Println(show(&secret{password: "123"}))
}

var count int
var rw sync.RWMutex

func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 5; i++ {
		go read(i, ch)
	}
	for i := 0; i < 5; i++ {
		go write(i, ch)
	}
	for i := 0; i < 10; i++ {
		<-ch
	}
}
func read(n int, ch chan struct{}) {
	rw.RLock()
	fmt.Printf("goroutine %d 进入读操作...\n", n)
	v := count
	fmt.Printf("goroutine %d 读取结束，值为：%d\n", n, v)
	rw.RUnlock()
	ch <- struct{}{}
}
func write(n int, ch chan struct{}) {
	rw.Lock()
	fmt.Printf("goroutine %d 进入写操作...\n", n)
	v := rand.Intn(1000)
	count = v
	fmt.Printf("goroutine %d 写入结束，新值为：%d\n", n, v)
	rw.Unlock()
	ch <- struct{}{}
}
