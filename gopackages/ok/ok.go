/*
@Time : 2019-02-23 11:49 
@Author : Tenlu
@File : ok
@Software: GoLand
*/
package main

import "fmt"

// 逗号ok模式
func main() {
	useMap()
	useChan()
	userType()
}

//1.检测map中是否含有指定key
func useMap() {
	company := map[string]string{"polly": "tencent", "robin": "baidu", "jack": "alibaba", "tenlu": "itech8"}
	if _, ok := company["toms"]; !ok {
		fmt.Println("key toms is not exists")
	}
	for ck, cv := range company {
		fmt.Println(ck + "->" + cv)
	}
}

// 2. 检测chan是否关闭
func useChan() {
	ch := make(chan int, 10)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch) // chan赋值结束,必须关闭通道
	elem := <-ch
	fmt.Println(elem) // 1  // FIFO队列,先发送的chan一定最先被接收

	for cc := range ch {
		fmt.Println(cc)
	}
	// ch <- 8  // 此时如果在发送数据则报错,panic: send on closed channel
	elem2 := <-ch
	elem3 := <-ch
	fmt.Printf("elem2:%d\n", elem2) // 0,因为通道已经关闭
	fmt.Printf("elem3:%d\n", elem3) // 0
	if _, isClosed := <-ch; !isClosed {
		fmt.Println("channel closed")
	}
	// go是没有while循环的,此处类似其他语言的while(true)
	for {
		i, ok := <-ch
		if !ok {
			fmt.Println("channel closed!")
			break
		}
		fmt.Println(i)
	}
}

// 3.检测是否实现了接口类型
type I interface {
	// 有一个方法的接口 I
	Get() Int
}

type Int int

// Int 类型实现了 I 接口
func (i Int) Get() Int {
	return i
}

func userType() {
	var myint Int = 5
	var inter I = myint // 变量赋值给接口
	val, ok := inter.(Int)
	fmt.Printf("%v, %v\n", val, ok) // 输出为：5，true
}