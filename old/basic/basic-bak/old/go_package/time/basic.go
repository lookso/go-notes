/*
@Time : 2019/3/1 10:00 AM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

// go time时间包
type Weekday int

const (
	Sunday    Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println(Monday)
	year, month, day := time.Now().Date()
	if month != time.August && day == 01 {
		fmt.Println("Happy Go Go day!")
	}
	fmt.Printf("now date:%d-%d-%d\n", year, month, day) // now date:2019-3-1
	fmt.Println(time.Month(Monday))                     // January

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())

	fmt.Println(time.Now().Unix())     // 时间戳,精确到秒
	fmt.Println(time.Now().UnixNano()) // 时间戳,精确到纳秒
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	t1 := time.Now().Year()       //年
	t2 := time.Now().Month()      //月
	t3 := time.Now().Day()        //日
	t4 := time.Now().Hour()       //小时
	t5 := time.Now().Minute()     //分钟
	t6 := time.Now().Second()     //秒
	t7 := time.Now().Nanosecond() //纳秒
	fmt.Printf("北京时间:%v年%v月%v日%v点%v分%v秒%v纳秒\n", t1, t2, t3, t4, t5, t6, t7)

	now := time.Now()
	nowYear, nowMonth, nowDay := now.Date()
	fmt.Printf("今天是:%d-%d-%d\n", nowYear, nowMonth, nowDay)
	fmt.Printf("当前时间戳精确到秒:%v\n", time.Now().Unix())       // 当前时间戳
	fmt.Printf("当前时间戳精确到纳秒:%v\n", time.Now().UnixNano()) // 当前时间戳
	fmt.Printf("格式化当前日期:%v\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v\n", time.Now().Format("2006-01-02"))

	currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local)
	fmt.Printf("%v\n", currentTimeData)

	fmt.Printf("今天是周%d\n", Friday)
	formatTimeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println("将当前时间戳转换为格式化日期:", formatTimeStr)

	nowDate := "2017-02-27 17:30:20"
	p, _ := time.Parse("2006-01-02 15:04:05", nowDate)
	fmt.Println("将指定日期转换为时间戳:", p.Unix())

}
