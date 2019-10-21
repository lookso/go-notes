/*
@Time : 2019-10-15 10:16 
@Author : Tenlu
@File : map
@Software: GoLand
*/
package main

import (
	"fmt"
	"strconv"
	"sync"
)
type MapLock struct {
	m    map[string]string
	lock sync.Mutex    // 不加锁会导致 fatal error: concurrent map read and map write
}

func (ml *MapLock) Set(key, value string) {
	ml.lock.Lock()
	ml.m[key] = value
	defer ml.lock.Unlock()
}
func (ml *MapLock) Get(key string) string {
	ml.lock.Lock()
	value := ml.m[key]
	defer ml.lock.Unlock()
	return value
}
func main() {
	var mainCh = make(chan bool)
	var goCh = make(chan int, 100)
	var mt MapLock
	mt.m = make(map[string]string)
	go func() {
		for i := 0; i < 20; i++ {
			goCh <- i
			mt.Set(strconv.Itoa(i), strconv.Itoa(i))
		}
	}()

	go func() {
		for q := 0; q < 20; q++ {
			<-goCh
			fmt.Println(mt.Get(strconv.Itoa(q)))
		}
		mainCh <- true
	}()
	<-mainCh
}