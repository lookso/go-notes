package main

import (
	"fmt"
)

func main() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}

	fmt.Println(v)

	var closeCh = make(chan int)
	var mainCh = make(chan int)
	go func() {
		for i := 0; i < len(closeCh); i++ {
			closeCh <- i
		}
		close(closeCh)
	}()
	go func() {
		for i := 0; i < len(mainCh); i++ {
			mainCh <- i
		}
		close(mainCh)
	}()

	for {
		select {
		case e := <-closeCh:
			fmt.Printf("Get element from closeCh: %d\n", e)
		case e := <-mainCh:
			fmt.Printf("Get element from cmainCh: %d\n", e)
		//case <-time.After(1 * time.Duration(2)):
		//	fmt.Println(errors.New("timeout").Error())
		//	return
		//}
		default:
			fmt.Println("default")
			return
		}
	}
	fmt.Println("main exit.")
}
