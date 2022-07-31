package main

import (
	"fmt"
	"log"
)

// 并发写入 map 导致的致命错误,recover无法捕获错误
func main() {
	m := make(map[int]string)
	for i := 0; i < 10; i++ {
		go func() {
			defer func() {
				if e := recover(); e != nil {
					log.Printf("recover: %v", e)
				}
			}()
			m[i] = "Go 语言编程之旅：一起用 Go 做项目"
		}()
	}
	fmt.Println(1234)
}
