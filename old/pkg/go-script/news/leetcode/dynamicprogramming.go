/*
@Time : 2019/5/15 4:34 PM 
@Author : Tenlu
@File : dynamicprogramming
@Software: GoLand
*/
package main

import (
	"os"
	"fmt"
)
// 写一个函数 输出 '只能用数字 1 2 5 10 求和等于10' 的所有组合
var rewards = []int{1,2,5,10}
func getSum(total int,res []int){
	if total == 0{
		fmt.Println(res)
		return
	}else if total<0 {
		return
	}
	i:=0
	for i<len(rewards){
		resTmp:=res
		resTmp = append(resTmp,rewards[i])
		getSum(total-rewards[i],resTmp)
		i++
	}
}
func main(){
	getSum(10,[]int{})
	os.Exit(0)
}

