package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	// 原子加减
	// 在原来的基础上加n
	var num int32 = 10
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
	// 假设被操作的值未曾被改变（即与旧值相等），并一旦确定这个假设的真实性就立即进行值替换
	// 如果想安全的并发一些类型的值，我们总是应该优先使用CAS
	if atomic.CompareAndSwapInt32(&addr, old, new) {
		fmt.Println("old", old)
		fmt.Println("new", new)
		fmt.Println("addr", addr)
	}
	// 不比较旧址,只是简单替换
	oldInt := atomic.SwapInt32(&addr, new)
	fmt.Println("oldInt", oldInt)

	fmt.Println("------LoadInt32--------")
	// 载入
	// 如果一个写操作未完成，有一个读操作就已经发生了，这样读操作使很糟糕的。加写锁
	// 为了原子的读取某个值sync/atomic代码包同样为我们提供了一系列的函数。这些函数都以"Load"为前缀，意为载入
	// 有了“原子的”这个形容词就意味着，在这里读取value的值的同时，当前计算机中的任何CPU都不会进行其它的针对此值的读或写操作。
	// 这样的约束是受到底层硬件的支持的
	// 原子读取count变量的内容
	var pv int32 = 2
	addValue(pv)
	fmt.Println("pv", pv)
	// 在原子的存储某个值的过程中，任何cpu都不会进行针对进行同一个值的读或写操作。
	// 如果我们把所有针对此值的写操作都改为原子操作，那么就不会出现针对此值的读操作读操作因被并发的进行而读到修改了一半的情况。
	atomic.StoreInt32(&pv, 4)
	fmt.Println(pv)
	type NewConfig struct {
		Name string `json:"name"`
	}
	var nc = NewConfig{
		Name: "123",
	}
	var config atomic.Value
	config.Store(nc)
	fmt.Println(config.Load())
}

func addValue(delta int32) {
	var value int32 = 2
	for {
		v := atomic.LoadInt32(&value)
		if atomic.CompareAndSwapInt32(&value, v, v+delta) {
			fmt.Println("value", value)
			break
		}
	}
}
