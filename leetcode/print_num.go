/*
@Time : 2019/5/13 10:12 AM
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	//getNum()
	getNum2()
	//getNum3()
	time.Sleep(2 * time.Second)
}

func getNum2() {
	a := make(chan bool, 1)
	b := make(chan bool)
	exit := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			if ok := <-a; ok {
				fmt.Println("a = ", 2*i-1)
				b <- true
			}
		}
	}()
	go func() {
		defer func() {
			close(exit)
		}()
		for i := 1; i <= 10; i++ {
			if ok := <-b; ok {
				fmt.Println("b : ", 2*i)
				a <- true
			}
		}
	}()
	a <- true
	<-exit
}

// 协程交替输出1-20之间的数字
func getNum() {

	//go func() {
	//	for i:=1;i<=10 ;i++  {
	//		if i%2==1 {
	//			fmt.Println(i)
	//		}
	//	}
	//}()
	//go func() {
	//	for i:=1;i<=10 ;i++  {
	//		if i%2==0 {
	//			fmt.Println(i)
	//		}
	//	}
	//}()

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			<-c1
			go func(i int) {
				fmt.Println(2*i - 1)
				c2 <- 1
			}(i)
		}
	}()
	go func() {
		for i := 1; i <= 10; i++ {
			<-c2
			go func(i int) {
				fmt.Println(2 * i)
				c1 <- 1
			}(i)
		}
	}()
	c1 <- 1
}

func getNum3() {
	ch := make(chan int)
	exit := make(chan struct{})

	go func() {
		for i := 1; i <= 20; i++ {
			println("g1:", <-ch)
			i++
			ch <- i
		}
	}()

	go func() {
		defer func() {
			close(ch)
			close(exit)
		}()
		for i := 0; i < 20; i++ {
			i++
			ch <- i
			println("g2:", <-ch)
		}
	}()

	<-exit
}
