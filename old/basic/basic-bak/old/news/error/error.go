/*
@Time : 2019/5/8 12:08 PM 
@Author : Tenlu
@File : error
@Software: GoLand
*/
package main

import "fmt"

func main()  {
	fmt.Println(getNum(10))
}
func getNum(v int)  int{
	defer func(v int){ v++ } (v)
	return v
}

