/*
@Time : 2019-08-10 22:21 
@Author : Tenlu
@File : defer2
@Software: GoLand
*/
package main

import "fmt"

var g = 100

func f() (r int) {
	defer func() {
		g = 200
	}()

	fmt.Printf("f: g = %d\n", g)

	return g
}
var gg = 100
func ff() (rr int) {
	rr = gg
	defer func() {
		rr = 200
	}()

	fmt.Printf("ff: rr = %d\n", rr)

	rr = 0
	return rr
}

func main() {
	i := f()
	fmt.Printf("main: i = %d, g = %d\n", i, g)
	fmt.Println("-----")
	ii := ff()
	fmt.Printf("main: ii = %d, gg = %d\n", ii, gg)


}
