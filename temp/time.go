package main

import (
	"fmt"
	"github.com/spf13/cast"
	"strings"
	"time"
)

func timeDate(t time.Time, sec, nsec int) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), sec, nsec, time.Local)
}

//3:26,3:27,3:28,3:29,3:30,3:31
//
//3:28-3:30
//
//3:26-3:28
//
//3:18
//
//3:13-3:18
//
//3:18

func main() {
	counselorIds:="1"
	fmt.Println(cast.ToIntSlice(strings.Split(counselorIds, ",")))

	queryMap := map[int]int{2100051137: 1, 58492: 1}
	m := map[int]int{100: 11212, 58492: 1}
	for k := range queryMap {
		fmt.Println(k)
		if _, ok := m[k]; ok {
			fmt.Println(k)
		}
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	current := time.Now()
	endTime := time.Now().Add(time.Minute * 1)
	//execStartTime := common.TimeDate(current, 0, 0).Unix()
	// 提高执行正确率,往后多查5分钟,避免时间误差
	minBefore := current.Add(-time.Minute * 5)
	//execStartTime := common.TimeDate(minBefore, 0, 0).Unix()
	//execEndTime := timeDate(endTime, 0, 0).Unix()

	tenMinStartTime := timeDate(minBefore, 0, 0).Format("2006-01-02 15:04:05")
	tenMinEndTime := timeDate(endTime, 0, 0).Format("2006-01-02 15:04:05")
	fmt.Println(tenMinStartTime)
	fmt.Println(tenMinEndTime)

}
