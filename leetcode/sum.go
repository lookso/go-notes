/*
@Time : 2019-08-14 19:36 
@Author : Tenlu
@File : sum
@Software: GoLand
*/
package main

import "fmt"

// 1-2+3-4+5-6+7-8+9-10....m
func main() {
	sum := 0
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			sum += i
		} else {
			sum -= i
		}
	}
	fmt.Println(sum)
}
