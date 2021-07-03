package main

import (
	"fmt"
	"time"
)

func AddFunc() {
	var s = []int{1, 2}
	if len(s)>=3 {
	//	fmt.Println(s[:3])
		fmt.Println(s[3:])
	}
}
func main() {
	AddFunc()
	timeLayout := "2006-01-02 15:04:05"
	layout := "2006-01-02"
	location, _ := time.LoadLocation("Local")
	timeStr := time.Now().Format(layout)
	fmt.Println(timeStr + " 00:00:00")

	startTime, err := time.ParseInLocation(timeLayout, timeStr+" 00:00:00", location)
	if err != nil {
		fmt.Println("start err", err)
	}
	startTimestamp := startTime.Unix()
	fmt.Println("start", startTimestamp)
	endTime, err := time.ParseInLocation(timeLayout, timeStr+" 23:59:59", location)
	if err != nil {
		fmt.Println("end err", err)
	}
	endTimestamp := endTime.Unix()
	fmt.Println("end", endTimestamp)
	fmt.Println("------------")
	actionLogList := make([]int, 0, 8)

	defer func() {
		if len(actionLogList) > 0 {
			fmt.Println(len(actionLogList))
		}
	}()
	actionLogList = []int{1, 2, 3, 4, 5, 6, 7}
}
