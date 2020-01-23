/*
@Time : 2019-07-28 15:15 
@Author : Tenlu
@File : go-rwmutux
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var group sync.WaitGroup
	group.Add(100)
	var rwmutux sync.RWMutex

	sum := 0
	var age = 10
	for i := 1; i <= 100; i++ {
		rwmutux.RLock()
		rwmutux.RLock()
		go func(i int) {
			sum = sum + i
			defer group.Done()
			rwmutux.RUnlock()
		}(i)
		age = 12
	}
	fmt.Println(age) //10
	group.Wait()
	fmt.Println(age) //12

	fmt.Println(sum)
}
