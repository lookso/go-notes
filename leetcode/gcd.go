/*
@Time : 2019-07-26 07:18 
@Author : Tenlu
@File : gcd
@Software: GoLand
*/
package main

import "fmt"

// 求出最大公约数和最小公倍数
func main() {
	numA, numB := 1280,720
	fmt.Println("最大公约数:",gcd(numA, numB))
	fmt.Println("最小公倍数:",minGys(numA, numB))
}
func gcd(numA, numB int) int {
	if numB == 0 {
		return numA
	} else {
		return gcd(numB, numA%numB)
	}
}

func minGys(numA,numB int) int {
	return numA*numB/gcd(numA,numB)
}