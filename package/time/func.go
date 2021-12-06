package main

import (
	"fmt"
	"time"
)

func main()  {
	q, _ := time.ParseInLocation("2006-01-02 15:04:05", "2001-01-01 08:05:43", time.Local)
	fmt.Println(q.IsZero(),q.Unix()) // false 978307543
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", "0001-01-01 08:05:43", time.Local)

	fmt.Println(t.IsZero(),t.Unix()) // true -62135596800
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", "0000-00-00 00:00:00", time.Local)
	fmt.Println(t1.IsZero(),t1.Unix()) // true -62135596800
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", "1970-01-01 00:00:00", time.Local)
	fmt.Println(t2.IsZero(),t2.Unix()) // false
}