package main

import (
	"fmt"
	"sync"
)

func main()  {
	var lock sync.Mutex
	var lock2 sync.Mutex
	a:=1
	lock.Lock()
	lock2.Lock()
	a++
	defer lock.Unlock()
	defer lock2.Unlock()
	fmt.Println(a)
}
