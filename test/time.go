package main

import (
	"fmt"
	"time"
)

func main() {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai") //上海
	todayDate := time.Now().In(cstSh)
	fmt.Println(todayDate)
}
