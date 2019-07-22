/*
@Time : 2019-07-22 16:38 
@Author : Tenlu
@File : goorder
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

/*
1.保证go协程顺序输出
*/

//func main() {
//	var wg = &sync.WaitGroup{}
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	wg.Add(3)
//	go sayA(wg, ch2, ch1)
//	go sayB(wg, ch1, ch2)
//	wg.Wait()
//}
//func sayA(wg *sync.WaitGroup, ch2 chan int, ch1 chan int) {
//	defer wg.Done()
//	for i := 1; i <= 10; i++ {
//		ch2 <- 2*i - 1
//		fmt.Println(<-ch1)
//	}
//}
//func sayB(wg *sync.WaitGroup, ch1 chan int, ch2 chan int) {
//	defer wg.Done()
//	for i := 1; i <= 10; i++ {
//		fmt.Println(<-ch2)  // 接收数据协程-类似协程同步的主协程
//		ch1 <- 2 * i
//	}
//}

func main() {
	var str = "abcdefgh"
	var wg = &sync.WaitGroup{}
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	count:=len(str)
	wg.Add(2)

	go func(str string) {
		for j := 0; j <= count-1; j++ {
			if j%2==0 {
				ch2 <- string(str[j])
				fmt.Println(<-ch1)
			}
		}
		wg.Done()
	}(str)
	go func(str string) {
		for i := 0; i <= count-1; i++ {
			if i%2!=0 {
				fmt.Println(<-ch2)
				ch1 <- string(str[i])
			}
		}
		wg.Done()
	}(str)
	wg.Wait()
}

//func main() {
//	var str = "abcdefgh"
//	go func(str string) {
//		for i := 11; i <= 20; i++ {
//			fmt.Println(i)
//		}
//	}(str)
//	go func(str string) {
//		for j := 1; j <= 10; j++ {
//			fmt.Println(j)
//		}
//		time.Sleep(time.Second*6)
//	}(str)
//	time.Sleep(time.Second*10)
//}