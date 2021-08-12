package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

var SleepTime = 3000

func main() {
	var a = []int{1}
	var b = append(a, 1)
	fmt.Println(a, b)

	classIDs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	if len(classIDs) > 5 {
		classIDs = classIDs[:5]
	}
	fmt.Println("classIDs", classIDs)
	fmt.Println(cast.ToStringSlice(12))

	var phone = "15010659038"
	sub := phone[len(phone)-4:]
	fmt.Println("sub", sub)

	time.Sleep(time.Second * 3)
	var ch = make(chan bool)
	go mt(ch)
	<-ch

	fmt.Println(99999)
}
func mt(ch chan bool) {

	fmt.Println(12312312321)
	fmt.Println(ch)
	for i := 0; i < 4; i++ {
		if i == 3 {
			continue
		}
		time.Sleep(time.Second * 4)
		fmt.Println("i", i)
	}
	fmt.Println("xs-123")
	ch <- true
}
