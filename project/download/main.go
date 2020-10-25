package main

import (
	"fmt"
	dl "go-notes/project/download/download"
)
// 协程池
func main() {
	end := make(chan bool)
	u := dl.Urls{
		Chs: make(chan int, 5), // 默认同时下载5个
		Ans: make(chan bool),
	}
	// 初始化url
	go u.InitUrl(end)
	if ok := <-end; ok {
		// 分发的下载线程
		go func() {
			for _, v := range u.Urls {
				u.Chs <- 1 // 限制线程数 （每次下载缓存加1， 直到加满阻塞）
				u.Wg.Add(1)
				go u.Work(v)
				fmt.Println(v)
			}
			u.Wg.Wait()  // 等待所有分发出去的线程结束
			close(u.Ans) // 否则range 会报错哦
		}()

		// 静静的等待每个下载完成
		for _ = range u.Ans {
			//	fmt.Println(status)
		}
	}
}

// Golang通过Goroutine+Channel指定同时下载的数量
// https://my.oschina.net/90design/blog/1607131

// https://studygolang.com/articles/12242?fr=sidebar
