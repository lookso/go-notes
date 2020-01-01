/*
@Time : 2019/3/14 7:34 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import "fmt"
import "math/rand"
import "time"

var (
	rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	a1 = [2]int{}
	a2 = [10]int{}
)

func main() {
	//a := []int{0, 1}
	for i := 0; i < 10000; i++ {
		randslice()
		//fmt.Println(rr.Intn(10))
	//	sliceOutOfOrder(a)
	}
	fmt.Println(a1)
	fmt.Println(a2)
}
//func sliceOutOfOrder(in []int) []int {
//	l := len(in)
//	for i := l - 1; i > 0; i-- {
//		r := rr.Intn(i)
//		in[r], in[i] = in[i], in[r]
//	}
//	a1[in[0]] += 1
//	return in
//}

func randslice() {
	in := rr.Perm(10)
	a2[in[0]] += 1
}


