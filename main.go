package go_notes

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t := time.NewTimer(2 * time.Second)
	count := 0
	done := make(chan bool)
	go func(tm *time.Timer, count int) {
		defer tm.Stop()
		for {
			select {
			case <-t.C:
				count++
				fmt.Println("timer running...." + strconv.Itoa(count))
				t.Reset(2 * time.Second)
			case _, ok := <-done:
				if ok {
					fmt.Println("game again ComMing....")
					break
				}
			}
		}
	}(t, count)
	done <- true
	close(done)
	select {}
}
