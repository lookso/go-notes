package main

import (
	"fmt"
	"time"
)
// 计算时间差
func main() {
	now := time.Now()
	fmt.Printf("since time:%+v\n", time.Since(now))
	after := time.Now()
	fmt.Printf("sub time:%+v\n", after.Sub(now))
}
