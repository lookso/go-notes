/*
@Time : 2019-02-23 16:40 
@Author : Tenlu
@File : goroutine_test
@Software: GoLand
*/
package main

import (
	"fmt"
	"time"
)
/**
当启动多个goroutine时，如果其中一个goroutine异常了，并且我们并没有对进行异常处理，那么整个程序都会终止，
所以我们在编写程序时候最好每个goroutine所运行的函数都做异常处理，异常处理采用recover

 */
func addele(a []int, i int) {
	defer func() { //匿名函数捕获错误
		err := recover()
		if err != nil {
			fmt.Println("add ele fail")
		}
	}()
	a[i] = i
	fmt.Println(a)
}

func main() {
	Arry := make([]int, 4)
	for i := 0; i < 10; i++ {
		go addele(Arry, i)
	}
	time.Sleep(time.Second * 2)
} //结果add ele fail
