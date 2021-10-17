package main

import (
	"fmt"
	"github.com/orcaman/concurrent-map"
)

func concurrentMap() {
	mp := cmap.New()
	// 存储值，key只能是string类型，value可以是任意类型
	mp.Set("1", "hello")
	mp.Set("2", "golang")
	//val: 如果存在则返回值，不存在则返回nil
	// ok： 表示值如果存在则是true，如果不存在则是false
	val, ok := mp.Get("2")
	if ok {
		fmt.Println(val.(string))
	}
	mp.Remove("1")
	_, ok = mp.Get("1")
	if !ok {
		fmt.Println("data is not exist")
	}
}
func main() {
	concurrentMap()
}
