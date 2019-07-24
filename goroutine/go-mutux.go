/*
@Time : 2019-07-22 23:10 
@Author : Tenlu
@File : mutux
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

var shareCnt int64 = 0

var lck sync.Mutex

func incrShareCnt(group *sync.WaitGroup) {
	for i:=0; i < 10; i++ {
		lck.Lock()
		shareCnt++
		lck.Unlock()
	}
	fmt.Println(shareCnt)
	group.Done()

}

func main()  {
	var group sync.WaitGroup
	group.Add(2)
	for i:=0; i < 2; i++ {
		go incrShareCnt(&group)
	}
	group.Wait()
	
}
