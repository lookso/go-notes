package main

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/petermattis/goid"
)

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion int32
}

func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}
func (m *RecursiveMutex) UnLock() {
	gid := goid.Get()
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d):%d", m.owner, gid))
	}
	m.recursion--
	if m.recursion != 0 {
		return
	}
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}

func main() {
	n := 0
	var wg sync.WaitGroup
	lock := new(RecursiveMutex)
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			n++
			defer lock.UnLock()
		}()
	}
	wg.Wait()
	fmt.Printf("n:%d", n)
}
