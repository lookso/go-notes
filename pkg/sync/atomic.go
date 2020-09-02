package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var num int32
	nv:=atomic.AddInt32(&num, 1)
	fmt.Println(nv)
}
