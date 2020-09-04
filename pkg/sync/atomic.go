package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num int32
	nv := atomic.AddInt32(&num, -1)
	fmt.Println(nv)
	var old int32
	var new int32
	var addr int32
	addr = 1
	old = 1
	new = 2
	// CAS 比较交换,乐观锁版本号的概念
	// 如果addr和old相同,就用new代替addr
	if atomic.CompareAndSwapInt32(&addr, old, new) {
		fmt.Println("old", old)
		fmt.Println("new", new)
		fmt.Println("addr", addr)
	}
}
