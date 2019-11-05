/*
@Time : 2019-11-05 19:44 
@Author : Tenlu
@File : mutex
@Software: GoLand
*/
package main

import (
	"fmt"
)

var nums = make([]int, 0)
var exit = make(chan bool)
var next = make(chan bool)

func main() {
	go toFirst(next)
	go toNext(next, exit)

	<-exit
	fmt.Println(nums)

}

func toFirst(next chan bool) {
	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	next <- true
}

func toNext(next, exit chan bool) {
	<-next // 不要放错位置
	for i := 11; i < 20; i++ {
		nums = append(nums, i)
	}
	exit <- true   // 塞放在逻辑下面
}
