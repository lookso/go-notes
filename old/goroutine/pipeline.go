/*
@Time : 2019/3/1 10:09 PM 
@Author : Tenlu
@File : pipeline
@Software: GoLand
*/
package main

import (
	"fmt"
)

func main() {
	father := make(chan int, 10)
	son := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			father <- i
		}
		close(father)
	}()
	if _, ok := <-father; ok {
		fmt.Printf("father chan closed" + "\n")
	}
	go func() {
		for f := range father {
			son <- f * f
		}
		close(son)
	}()
	if _, ok := <-son; ok {
		fmt.Printf("son chan closed" + "\n")
	}
	for s := range son {
		fmt.Println(s)
	}
}
