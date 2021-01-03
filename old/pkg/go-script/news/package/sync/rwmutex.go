/*
@Time : 2019/5/12 7:10 PM 
@Author : Tenlu
@File : rwmutex
@Software: GoLand
*/
package main

import (
	"sync"
	"fmt"
	"time"
)

//RWMutex是读写互斥锁。该锁可以被同时多个读取者持有或唯一个写入者持有。
//RWMutex可以创建为其他结构体的字段；零值为解锁状态。
//RWMutex类型的锁也和线程无关，可以由不同的线程加读取锁/写入和解读取锁/写入锁。

func main()  {
	fmt.Println("---------mutex-------")
	mutexRun()
	fmt.Println("---------rwmutex-------")
	rwMutexrun()
}
func mutexRun() {
	var rwnum = 0
	var m sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)        //开启协程
		go func(m *sync.Mutex,wg *sync.WaitGroup) {
			m.Lock()  //通过加互斥锁处理,现在我们可以对上述程序加上锁，每次只能由一个线程来操作num的值
			rwnum = rwnum + 1
			m.Unlock()
			wg.Done()
		}(&m,&wg)
	}
	wg.Wait()
	fmt.Println(rwnum) // 1000
}
// 计数器
type counter struct {
	num uint // 计数
	mu sync.RWMutex //读写锁
}
// 获取num值的操作，加读锁
func (c *counter) number() uint {
	// RLock方法将rw锁定为读取状态，禁止其他线程写入，但不禁止读取
	c.mu.RLock()
	// Runlock方法解除rw的读取锁状态，如果m未加读取锁会导致运行时错误。
	defer c.mu.RUnlock()
	return c.num
}

// 修改num值的操作，加写锁
func (c *counter) add (increment uint) uint {
	// Lock方法将rw锁定为写入状态，禁止其他线程读取或者写入。
	c.mu.Lock()
	// Unlock方法解除rw的写入锁状态，如果m未加写入锁会导致运行时错误。
	defer c.mu.Unlock()
	c.num += increment
	return c.num
}

func rwMutexrun()  {
	c := counter{}
	done := make(chan struct{})

	// 增加计数器
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		for i := 0; i < 10; i++ {
			time.Sleep(time.Millisecond * 500)
			c.add(1)
		}
	}()
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		for j := 0; j < 20; j++ {
			time.Sleep(time.Millisecond * 200)
			fmt.Printf("[%d-%02d] 读数: %d\n", 1, j, c.number())
		}
	}()
	go func() {
		defer func() {
			done <- struct{}{}
		}()
		for k := 0; k < 20; k++ {
			time.Sleep(time.Millisecond * 300)
			fmt.Printf("[%d-%02d] 读数: %d\n", 2, k, c.number())
		}
	}()
	<- done
	<- done
	<- done
}