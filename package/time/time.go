package main

import (
	"fmt"
	"time"
)

func main() {
	layout := "2006-01-02 15:04:05"
	// 时间戳转格式化
	afterDayTimestamp := time.Now().Unix() + int64(10*86400)
	dt := time.Unix(afterDayTimestamp, 0).Format(layout)
	fmt.Println(dt)
	// 格式化转时间戳
	ty, err := time.ParseInLocation(layout, dt, time.Local)
	if err != nil {
		fmt.Println("ParseInLocation err", err)
	}
	fmt.Println(ty.Unix())
	// 计时器
	limiter := time.Tick(time.Second * 5)
	for {
		select {
		case <-limiter:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}
}

// time.tick()和time.sleep() 区别:https://blog.csdn.net/Star_CSU/article/details/86650684
