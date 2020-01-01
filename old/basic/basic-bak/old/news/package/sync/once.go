/*
@Time : 2019/5/11 10:54 AM 
@Author : Tenlu
@File : once
@Software: GoLand
*/
package main

import (
	"sync"
	"fmt"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	done := make(chan bool)
	for i:=0;i<10;i++  {
		once.Do(onceBody) // Only once
		// 如果把once放在协程函数中执行
		go func() {
			once.Do(onceBody)
			done <- true
		}()
	}

	for i := 0; i < 10; i++ {
		<-done
	}
	onceFunc()
}

func onceFunc()  {
	// 整个程序,只会执行onces()方法一次,onced()方法是不会被执行的。
	var once sync.Once
	for i, v := range make([]string, 10) {
		once.Do(onces)
		once.Do(onces) // 第二个onces函数不会被执行
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			once.Do(onced) // 此处函数也不会被执行
			fmt.Println(i)
		}(i)
	}
}
func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}