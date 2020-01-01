/*
@Time : 2019-07-22 16:10 
@Author : Tenlu
@File : waitgroup
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

func main() {
	var syncGroup sync.WaitGroup
	var family = []string{"father", "mather", "brother", "sister", "wife"}
	count := len(family)
	syncGroup.Add(count)
	for i := 0; i <= count-1; i++ {
		go func() {
			fmt.Println("i:", i)
			syncGroup.Done()
		}()
	}
	syncGroup.Wait()
}
