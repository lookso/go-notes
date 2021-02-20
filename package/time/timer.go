package main

import (
	"fmt"
	"time"
)

// 定时器
// 可参考缓存自动失效机制
func newTimer() {
	t := time.NewTimer(2 * time.Second)
	done := make(chan bool)
	go func(tm *time.Timer) {
		defer tm.Stop()
		for {
			select {
			case <-t.C:
				fmt.Println("timer running....")
				t.Reset(2 * time.Second)
			case _, ok := <-done:
				if ok {
					fmt.Println("game again comming....")
					break
				}
			}
		}
	}(t)
	done <- true
	close(done)
	select {}
}
func newTicker() {
	t := time.NewTicker(2 * time.Second)
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	defer t.Stop()
	aft := <-t.C
	fmt.Println(aft.Format("2006-01-02 15:04:05"))
}

func newAfter() {
	t := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3秒
	af := <-t
	fmt.Printf("t=%s", af.Format("2006-01-02 15:04:05"))
	ch := make(chan bool)

	time.AfterFunc(2*time.Second, func() {
		ch <- true
	})
	fmt.Println("ch",<-ch)
}

func main() {
	newAfter()
	newTicker()
	newTimer()
}
