/*
@Time : 2019/5/9 7:03 PM 
@Author : Tenlu
@File : waitgroup
@Software: GoLand
*/
package main

import (
	"fmt"
	"sync"
)

func main()  {
	run()
	fmt.Println()
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()

	fmt.Println(total.value)
	
}

var total struct{
	sync.Mutex
	value int
}

func worker(wg *sync.WaitGroup)  {
	defer wg.Done()

	for i := 0; i <= 100; i++ {
		total.Lock()
		total.value += i
		total.Unlock()
	}
}

func run() {
	var wg sync.WaitGroup
	for i:=0; i<3; i++ {
		//增加一个计数器
		wg.Add(1)
		go func(wg *sync.WaitGroup,mark int){
			//减去计数器
			defer wg.Done()//等价于 wg.Add(-1)
			fmt.Printf("%d goroutine finish \n",mark)
		}(&wg,i)
	}

	//等待所有go程结束
	wg.Wait()
	fmt.Printf("%s","main exit ....")
}

