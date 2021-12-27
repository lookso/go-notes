package main

import (
	"fmt"
	"time"
)

func main() {
	var ts = "0001-01-01 00:00:00"
	nt, _ := time.ParseInLocation("2006-01-02 15:04:05", ts, time.Local)
	fmt.Println("cast.ToInt(nt.Unix())",nt.Unix())
	fmt.Println("nt.Unix()",time.Unix(nt.Unix(),0).Format("2006-01-02 15:04:05"))

	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	todayDate := time.Now().In(cstSh)
	fmt.Println(todayDate)
}
