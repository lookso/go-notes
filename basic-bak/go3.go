package main

import (
	"fmt"
	"time"
)

func main()  {
	var state =make(chan bool)
	go func() {
		state<-true
		fmt.Println(time.Now().Unix())
	}()
	<-state
	fmt.Println("----main----")
}
