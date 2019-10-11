package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var pool sync.Pool
	pool.New = func() interface{} {
		return getTimestamp()
	}
	pool.Put(getFormatNowTime())
	pool.Put(getTimestamp())
	pool.Put("oFashion")

	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
func getFormatNowTime() string {
	return time.Now().Format("2006-01-02 15:03:04")
}
func getTimestamp() int64 {
	return time.Now().Unix()
}
