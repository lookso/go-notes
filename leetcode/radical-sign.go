/*
@Time : 2019-07-29 22:18 
@Author : Tenlu
@File : radical-sign
@Software: GoLand
*/
package main

import (
	"math"
)

func main() {
	redicalSign(12, 0)
}

// 求根号运算
func redicalSign(num, t int) int {
	var x1 = num
	var x2 = x1 / 2
	for int(math.Abs(float64(x1-x2))) > t {
		x1 = x2
		x2 = (x1 + num/2) / 2
	}
	return x2
}
