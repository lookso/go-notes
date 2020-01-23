/*
@Time : 2019-11-03 21:16 
@Author : Tenlu
@File : nil
@Software: GoLand
*/
package main

import "fmt"

func main() {
	var a []int
	if a == nil {
		fmt.Println("nil")
	}
	var p *int
	if p == nil {
		fmt.Println("nil")
	}
	//if p == a {    // 报错
	// ./nil.go:20:7: invalid operation: p == a (mismatched types *int and []int)
	//}
}
