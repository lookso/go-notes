/*
@Time : 2019-07-23 16:53
@Author : Tenlu
@File : pprof
@Software: GoLand
*/
package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

/*
泄露的场景不仅限于以下两类，但因channel相关的泄露是最多的。
channel的读或者写：
无缓冲channel的阻塞通常是写操作因为没有读而阻塞
有缓冲的channel因为缓冲区满了，写操作阻塞
期待从channel读数据，结果没有goroutine写
select操作，select里也是channel操作，如果所有case上的操作阻塞，goroutine也无法继续执行。

*/
func main() {
	var group sync.WaitGroup
	var mutux sync.Mutex
	group.Add(10000)
	for i := 1; i < 10000; i++ {
		mutux.Lock()
		go func() {
			mutux.Unlock()
			fmt.Println(i)
			group.Done()
		}()
	}

	var done = make(chan bool)
	go func() {
		for j:=1;j<10000 ;j++  {
			 done<-true
		}
		close(done)
	}()
	go func() {
		for j:=1;j<10000 ;j++  {
			done<-true
		}
		close(done)
	}()
	time.Sleep(time.Second)
	http.ListenAndServe("localhost:6060", nil)
	group.Wait()
}
