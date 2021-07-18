package main

import (
	"fmt"
	"time"
)

func newTicker() {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	t := time.NewTicker(2 * time.Second)
	defer t.Stop()
	aft := <-t.C
	fmt.Println(aft.Format("2006-01-02 15:04:05"))
	for range t.C {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	}
}
func main() {
	newTicker()
}
