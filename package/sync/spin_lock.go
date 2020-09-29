package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// 自旋锁的实现
type SpinLock uint32

func (sl *SpinLock) Lock() {
	for !atomic.CompareAndSwapUint32((*uint32)(sl), 0, 1) { // cas
		runtime.Gosched() // //强制调度器切换,让出时间片，先让别的协议执行，它执行完，再回来执行此协程
	}
}
func (sl *SpinLock) Unlock() {
	atomic.StoreUint32((*uint32)(sl), 0)
}
func NewSpinLock() *SpinLock {
	var lock SpinLock
	return &lock
}
func main() {
	var group sync.WaitGroup
	var sl = NewSpinLock()
	for i := 1; i <= 10; i++ {
		group.Add(1)
		sl.Lock()
		go func(i int) {
			sl.Unlock()
			defer group.Done()
			fmt.Println(i)
		}(i)
	}
	group.Wait()
}