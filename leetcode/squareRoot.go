/*
@Time : 2019-10-28 21:42 
@Author : Tenlu
@File : square
@Software: GoLand
*/
package main

import (
	"fmt"
)

func main() {
	fmt.Println(squareRoot(9))
}

func getabs(x float64) float64 {

	if x < 0 {
		return -x
	}

	if x == 0 {
		return 0
	}
	return x
}

// 牛顿迭代法
func squareRoot(x float64) float64 {
	z := 1.0

	if x < 0 {
		return 0
	} else if x == 0 {
		return 0
	} else {
		for getabs(z*z-x) > 0.000000000001 {
			z = (z + x/z) / 2
		}
		return z
	}
}
