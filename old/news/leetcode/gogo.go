/*
@Time : 2019/5/13 10:12 AM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"time"
	"fmt"
)

func main()  {
	//getNum()
	getNum2()
	//getNum3()
	time.Sleep(2*time.Second)
}

// 协程交替输出1-20之间的数字
func getNum()  {

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
				c2<-1
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
	c1<-1
}

func getNum2()  {
	A := make(chan bool, 1)
	B := make(chan bool)
	Exit := make(chan bool)

	go func() {
		for i := 1; i <= 10; i++ {
			if ok := <-A; ok {
				fmt.Println("A = ", 2*i-1)
				B <- true
			}
		}
	}()
	go func() {
		defer func() {
			close(Exit)
		}()
		for i := 1; i <= 10; i++ {
			if ok := <-B; ok {
				fmt.Println("B : ", 2*i)
				A <- true
			}
		}
	}()
	A <- true
	<-Exit
}
func getNum3()  {
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

