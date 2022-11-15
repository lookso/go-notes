package main

import (
	"fmt"
	"strconv"
	"time"
)

// 定时器
// 可参考缓存自动失效机制
func newTimer() {
	fmt.Println("------------------timer----------------")
	t := time.NewTimer(2 * time.Second)
	count := 0
	done := make(chan bool)
	go func(tm *time.Timer, count int) {
		defer tm.Stop()
		for {
			select {
			case <-t.C:
				count++
				fmt.Println("timer running...." + strconv.Itoa(count))
				t.Reset(2 * time.Second)
			case _, ok := <-done:
				if ok {
					fmt.Println("game again comming....")
					break
				}
			}
		}
	}(t, count)
	done <- true
	close(done)
	select {}
}

func newAfter() {
	fmt.Println("------------------- time.After---------------")
	t := time.After(time.Second * 3)
	fmt.Printf("t type=%T\n", t)
	//阻塞3秒
	af := <-t
	fmt.Printf("t=%s", af.Format("2006-01-02 15:04:05"))
	ch := make(chan bool)

	time.AfterFunc(2*time.Second, func() {
		ch <- true
	})
	fmt.Println("ch", <-ch)
}

func main() {
	newAfter()
	newTimer()
}
