package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	n, m := 5, 10
	redPack(n, m)
}
func redPack(n, m int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	var rmc int
	var mr = make(map[int]int, 0)
	for i := 1; i <= n; i++ {
		rm := r.Intn(m)
		fmt.Println("rmmm", rm)
		rmc = rm + rmc
		var rt int
		if rmc > m {
			fmt.Println("i", i)
			for j := 1; j < i; j++ {
				rt = m - mr[j]
			}
			fmt.Println("rt", rt)

			rm = 0
		}
		mr[i] = rm
		fmt.Printf("i:%d money:%d\n", i, rm)
	}
}
