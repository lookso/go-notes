package math

import "fmt"

func mySqrt(x float64) float64 {
	res := x
	//牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
		fmt.Println(res)
	}
	return res
}
