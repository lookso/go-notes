package main

import (
	"fmt"
	"github.com/bluele/gcache"
	"time"
)

// 提供了 LRU, LFU 和 ARC 三种缓存失效策略
func main() {
	gc := gcache.New(1).
		LRU().
		Build()
	gc.Set("key", "okokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokokok")
	value, err := gc.Get("key")
	if err != nil {
		panic(err)
	}
	gc.SetWithExpire("name", "jack", 5*time.Second)
	v2,err:=gc.Get("name")
	fmt.Println("Get:", value)
	fmt.Println("Getv2:", v2)

}
