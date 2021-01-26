/*
@Time : 2019/5/13 9:38 AM 
@Author : Tenlu
@File : testLock
@Software: GoLand
*/
package main

import (
	"sync"
	"time"
	"fmt"
)

var rw sync.RWMutex
var testNum int

func A(){
	rw.RLock()
	defer rw.RUnlock()
	time.Sleep(1* time.Second)
}

func main()  {
	go A()
	time.Sleep(3* time.Second) // 如果主协程的sleep小于子协程的sleep时间必然死锁
	rw.Lock()

	testNum++
	defer rw.Unlock()
	fmt.Println(testNum)
}
