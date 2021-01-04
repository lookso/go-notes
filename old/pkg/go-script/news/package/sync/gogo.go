package main

import "fmt"
import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	fmt.Println("----run in man----")

	go func() {

		defer wg.Done()
		fmt.Println("----111----")
		go func() {

			defer wg.Done()
			fmt.Println("----222----")
			go func() {

				defer  wg.Done()
				fmt.Println("----333----")
				go func() {

					wg.Done()

					fmt.Println("----444----")

				}()
			}()
		}()
	}()
	fmt.Println("--------main exit-------")
	wg.Wait()
}