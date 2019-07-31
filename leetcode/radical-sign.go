/*
@Time : 2019-07-29 22:18 
@Author : Tenlu
@File : radical-sign
@Software: GoLand
*/
package main

func main() {
	redicalSign(12)
}

// 求根号运算
func redicalSign(num int) int {
	for i := 1; i < num; i++ {
		for j := 1; j < num; j++ {
			if i*j==12{
				return i
			}
		}
	}
	return 0
}
