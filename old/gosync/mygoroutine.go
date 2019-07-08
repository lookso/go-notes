/*
@Time : 2019-02-23 18:20
@Author : Tenlu
@File : mygoroutine
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

func Productor(mychan chan int, data int, wait *sync.WaitGroup) {
	mychan <- data
	fmt.Println("product data：", data)
	wait.Done()
}
func Consumer(mychan chan int, wait *sync.WaitGroup) {
	a := <-mychan
	fmt.Println("consumer data：", a)
	wait.Done()
}
func main() {

	datachan := make(chan int, 100) //通讯数据管道
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go Productor(datachan, i, &wg) //生产数据
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(datachan, &wg) //消费数据
		wg.Add(1)
	}
	wg.Wait()
}

// 结果
//consumer data： 1
//product data： 9
//consumer data： 4
//consumer data： 8
//consumer data： 9
//product data： 0
//consumer data： 0
//product data： 6
//consumer data： 6
//product data： 5
//consumer data： 5
//product data： 7
//product data： 1
//consumer data： 7
//product data： 3
//consumer data： 3
//product data： 8
//consumer data： 2
//product data： 2
//product data： 4
