/*
@Time : 2019/5/11 10:54 AM
@Author : Tenlu
@File : once
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	testMap:=make(map[int]int)
	testMap[1]=1
	v:=testMap[1]
	fmt.Println(v)

	var once sync.Once
	for i := 0; i < 5; i++ {
		go func(i int) {
			fun1 := func() {
				fmt.Printf("i:=%d\n", i)
			}
			once.Do(fun1)
		}(i)
	}
	// 为了防止主goroutine直接运行完了，啥都看不到
	time.Sleep(50 * time.Millisecond)
}