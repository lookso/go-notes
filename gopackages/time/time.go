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
	MyFormat()
	// 计时器
	limiter := time.NewTicker(time.Second * 2)
	defer limiter.Stop()
	for {
		select {
		case <-limiter.C:
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		}
	}

}

// go time时间包
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

func MyFormat() {

	fmt.Printf("今天是%d年,第%d天\n", time.Now().Year(), time.Now().YearDay())

	//一年中的第几周
	year, week := time.Now().ISOWeek()
	fmt.Println("一年中的第几周: ", year, " | ", week)

	//当前是周几
	fmt.Println("当前是周几: ", time.Now().Weekday().String())

	hBefor, _ := time.ParseDuration("-24h")
	fmt.Println("24h前:", time.Now().Add(hBefor).Format("2006-01-02 15:04:05"))

	hAfter, _ := time.ParseDuration("24h")
	fmt.Println("24h后:", time.Now().Add(hAfter).Format("2006-01-02 15:04:05"))

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
	fmt.Printf("当前时间戳精确到秒:%v\n", time.Now().Unix())      // 当前时间戳
	fmt.Printf("当前时间戳精确到纳秒:%v\n", time.Now().UnixNano()) // 当前时间戳
	fmt.Printf("格式化当前日期:%v\n", time.Now().Format("2006-01-02 15:04:05"))
	fmt.Printf("%v\n", time.Now().Format("2006-01-02"))

	currentTimeData := time.Date(t1, t2, t3, t4, t5, t6, t7, time.Local)
	fmt.Printf("%v\n", currentTimeData)

	formatTimeStr := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	fmt.Println("将当前时间戳转换为格式化日期:", formatTimeStr)

	nowDate := "2017-02-27 17:30:20"
	p, _ := time.Parse("2006-01-02 15:04:05", nowDate)
	fmt.Println("将指定日期转换为时间戳:", p.Unix())
}

// time.tick()和time.sleep() 区别:https://blog.csdn.net/Star_CSU/article/details/86650684
