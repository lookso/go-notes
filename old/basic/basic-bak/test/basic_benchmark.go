/*
@Time : 2019-07-27 21:49 
@Author : Tenlu
@File : basic_benchmark
@Software: GoLand
*/
package main

import "testing"

func BenchmarkAdd(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
