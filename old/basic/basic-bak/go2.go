package main

import (
	"fmt"
	"sync"
)
type goSync struct {
	group sync.WaitGroup
	lock sync.Mutex
}
func main() {
	var goSync goSync // 想想为啥不可以直接使用类似 goSync{}.group.Add()
	// Slice和map 并发不安全,Chan并发安全
	var qp []int
	var dict =make(map[int]int,10)
	//var nums=make(chan int,10)
	var arr = [10]int{}
	for i := 0; i < 5; i++ {
		goSync.group.Add(1)
		go func(i int) {
			goSync.lock.Lock()
			qp= append(qp , i)
			dict[i] = i
			goSync.lock.Unlock()
			arr[i] = i
			defer goSync.group.Done()
		}(i)
	}
	goSync.group.Wait()
	fmt.Println("slice:", qp)
	fmt.Println("arr:", arr)
	fmt.Println("map:", dict)

//slice: [4 0 1 2 3]
//arr: [0 1 2 3 4 0 0 0 0 0]
//map: map[0:0 1:1 2:2 3:3 4:4]


}


