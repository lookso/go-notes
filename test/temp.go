package main

import (
	"fmt"
	"time"
)

func main() {
	timeLayout := "2006-01-02 15:04:05"
	layout:="2006-01-02"
	location, _ := time.LoadLocation("Local")
	timeStr := time.Now().Format(layout)
	fmt.Println(timeStr+" 00:00:00")

	startTime, err := time.ParseInLocation(timeLayout, timeStr+" 00:00:00", location)
	if err != nil {
		fmt.Println("start err", err)
	}
	startTimestamp := startTime.Unix()
	fmt.Println("start",startTimestamp)
	endTime, err := time.ParseInLocation(timeLayout, timeStr+" 23:59:59", location)
	if err != nil {
		fmt.Println("end err", err)
	}
	endTimestamp := endTime.Unix()
	fmt.Println("end",endTimestamp)
}
