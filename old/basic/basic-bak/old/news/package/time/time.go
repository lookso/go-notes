/*
@Time : 2019/5/25 3:12 PM 
@Author : Tenlu
@File : time
@Software: GoLand
*/
package main

import (
	"time"
	"fmt"
)

func main()  {
	t1 := time.Now().Year()       //年
	t2 := time.Now().Month()      //月
	t3 := time.Now().Day()        //日
	t4 := time.Now().Hour()       //小时
	t5 := time.Now().Minute()     //分钟
	t6 := time.Now().Second()     //秒
	t7 := time.Now().Nanosecond() //纳秒
	fmt.Printf("北京时间:%v年%v月%v日%v点%v分%v秒%v纳秒\n", t1, t2, t3, t4, t5, t6, t7)

	now := time.Now()
	year, month, day := now.Date()
	fmt.Printf("今天是:%d-%d-%d\n", year, month, day)
	fmt.Printf("当前时间戳精确到秒:%v\n", time.Now().Unix())      // 当前时间戳
	fmt.Printf("当前时间戳精确到纳秒:%v\n", time.Now().UnixNano()) // 当前时间戳
	fmt.Printf("格式化当前日期:%v\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v\n", time.Now().Format("2006-01-02"))

	currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local)
	fmt.Printf("%v\n", currentTimeData)

	type Weekday int
	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("今天是周%d\n", Friday)
	formatTimeStr := time.Unix(time.Now().Unix(),0).Format("2006-01-02 15:04:05")
	fmt.Println("将当前时间戳转换为格式化日期:", formatTimeStr)

	nowDate := "2017-02-27 17:30:20"
	p, _ := time.Parse("2006-01-02 15:04:05", nowDate)
	fmt.Println("将指定日期转换为时间戳:", p.Unix())

	fmt.Println("今天是周:",time.Now().Weekday().String())
	time.Sleep(time.Duration(1) * time.Second)
}

