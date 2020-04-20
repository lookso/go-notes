package main

import (
	"fmt"
	"log"
	"time"
)

type Resp struct {
	data  int
	error error
}

func main() {
	handleMsg()
}

func handleMsg() {
	resp := make(chan Resp)
	stop := make(chan struct{})
	go func() {
		t := time.Tick(time.Second)
		index := 0
		for {
			select {
			case <-t:
				resp <- Resp{
					data: index,
				}
				index++
			case <-stop:
				resp <- Resp{
					error: fmt.Errorf("time tick stop error"),
				}
			}
		}
	}()
	for {
		select {
		case val := <-resp:
			if val.error != nil {
				log.Fatal(val.error)
			}
			if val.data == 5 {
				stop <- struct{}{}

			}
			fmt.Println("index", val.data)

		}
	}
}