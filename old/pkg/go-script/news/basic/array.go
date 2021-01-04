/*
@Time : 2019/5/15 3:08 PM 
@Author : Tenlu
@File : array
@Software: GoLand
*/
package main

import "fmt"

func main(){
	pointArr()
}

func pointArr(){
	const MAX = 3
	a := []int{10,100,200}
	var i int
	var ptr [MAX]*int

	for  i = 0; i < MAX; i++ {
		ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
	}

	for  i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
	}
	// 打印指针数组元素的内存地址
	fmt.Println(ptr)
}