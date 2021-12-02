package main

import (
	"fmt"
	"time"
)

func main()  {
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "0001-01-01 08:05:43", time.Local)
	fmt.Println(t.IsZero()) // true
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "0000-00-00 00:00:00", time.Local)
	fmt.Println(t1.IsZero()) // true
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "1970-01-01 00:00:00", time.Local)
	fmt.Println(t2.IsZero()) // false
}
