/*
@Time : 2019-07-24 17:24
@Author : Tenlu
@File : single
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct {
	Name string
}

var ins *singleton
var insTwo *singleton
func NewIns(name string) *singleton {
	var once sync.Once
	var lock sync.Mutex
	// 存在线程安全问题，高并发时有可能创建多个对象
	if ins != nil {
		lock.Lock()
		return ins
	}
	once.Do(func() {
		ins = &singleton{
			Name: name,
		}
	})
	return ins
}
func NewInsTwo(name string) *singleton {
	var once sync.Once
	if insTwo != nil {
		return insTwo
	}
	once.Do(func() {
		insTwo = &singleton{
			Name: name,
		}
	})
	return insTwo
}
//模拟sync.once
type OnlyOne struct {
	m    sync.Mutex
	done uint32
}

func (o *OnlyOne) Once(f func()) {
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}
	o.m.Lock()
	defer o.m.Unlock()
	if o.done == 0 {
		atomic.StoreUint32(&o.done, 1)
	}
	f()
}
// 饿汉模式
// 构建一个结构体，用来实例化单例
type example2 struct {
	name string
}

// 声明一个私有变量，作为单例
var instance2 *example2

// init函数将在包初始化时执行，实例化单例
func init() {
	instance2 = new(example2)
	instance2.name = "初始化单例模式"
}

func GetInstance2() *example2 {
	return instance2
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			singleA := NewIns("jack")
			singleB := NewIns("toms")
			fmt.Println(singleA.Name)
			fmt.Println(singleB.Name)

			single1 := NewInsTwo("two")
			single2 := NewInsTwo("one")
			fmt.Println(single1.Name)
			fmt.Println(single2.Name)
			defer wg.Done()
		}()
		s := GetInstance2()
		fmt.Println(s.name)
	}
	wg.Wait()
}
