/*
@Time : 2019-07-23 08:32 
@Author : Tenlu
@File : goorder-mutux
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)
func main() {
	var mutux = &sync.Mutex{}
	var group = &sync.WaitGroup{}
	group.Add(2)
	mutux.Lock()
	//go func() {
	//	for i := 1; i <= 20; i++ {
	//		fmt.Println(i)
	//	}
	//	mutux.Unlock()
	//	group.Done()
	//}()

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(i)
		}
		mutux.Unlock()
		mutux.Lock()
		group.Done()
	}()
	go func() {
		for i := 11; i <= 20; i++ {
			fmt.Println(i)
		}
		mutux.Unlock()
		group.Done()
	}()
	group.Wait()

}
