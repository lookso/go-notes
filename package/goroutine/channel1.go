package main

import (
	"fmt"
	"sync"
)

func main() {
	var a = 100
	var p *int = &a
	fmt.Println(*p)
	var mt = make(map[string]string, 0)
	mt["name"] = "jacks"
	if value, ok := mt["name"]; ok {
		fmt.Println(value)
	}

	var ch = make(chan int)
	var mainCh = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go product(&wg, ch)
	go consumption(&wg, ch, mainCh)
	for m:=range mainCh {
		fmt.Println(m)
	}
	wg.Wait()
}

func product(wg *sync.WaitGroup, ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()
}

func consumption(wg *sync.WaitGroup, ch chan int, mainCh chan int) {
	for c := range ch {
		mainCh <- c * 2
	}
	close(mainCh)
	wg.Done()
}
