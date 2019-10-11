package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	fmt.Println("北京时间:", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println("当前纳秒时间戳:", time.Now().UnixNano())
	// 时间戳转换为格式化时间
	tamp, _ := strconv.ParseInt("1570287850", 10, 64)
	fmt.Println("时间戳转格式化", time.Unix(tamp, 0).Format("2006-01-02 15:04:05"))
	// 格式化时间转换为时间戳
	// 从字符串转为时间戳，第一个参数是格式，第二个是要转换的时间字符串
	tm, _ := time.Parse("2006-01-02 15:04:05", "2019-10-05 15:04:10")
	fmt.Println("tm", tm) // tm 2019-10-05 15:04:10 +0000 UTC
	timeStamp := tm.Unix()
	fmt.Println("字符串/格式化转换为时间戳", timeStamp)

	fmt.Println(time.Hour*2, time.Minute*2, time.Second*2)
	// 输出:2h0m0s 2m0s 2s 将时间段表示为float64类型的小时数,分钟数和秒数

	fmt.Println("现在是", time.Now().Hour(), "点")

	var tt = new(time.Time)
	format := tt.Format("2006-01-02 15:04:05")
	fmt.Println(format)
	start := time.Now()
	ch := make(chan int, 2)
	go func() {
		time.Sleep(4 * time.Second)
		ch <- 1
	}()

	select {
	case <-ch:
		fmt.Println("current value:", <-ch)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
	fmt.Println(time.Since(start))

	fmt.Println("当前协程数目:",runtime.NumGoroutine())
}
