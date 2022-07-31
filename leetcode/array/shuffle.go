package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	Shuffle(slice)
	fmt.Println(slice)
	for _, v := range slice {
		fmt.Println(v)
	}
}

func Shuffle(slice []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(slice) > 0 {
		n := len(slice)
		randIndex := r.Intn(n)
		slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
		slice = slice[:n-1]
	}
}

var (
	rr = rand.New(rand.NewSource(time.Now().UnixNano()))
	a1 = [2]int{}
	a2 = [10]int{}
)

func main1() {
	//a := []int{0, 1}
	for i := 0; i < 10000; i++ {
		randslice()
		//fmt.Println(rr.Intn(10))
		//	sliceOutOfOrder(a)
	}
	fmt.Println(a1)
	fmt.Println(a2)
}
func sliceOutOfOrder(in []int) []int {
	l := len(in)
	for i := l - 1; i > 0; i-- {
		r := rr.Intn(i)
		in[r], in[i] = in[i], in[r]
	}
	a1[in[0]] += 1
	return in
}

func randslice() {
	in := rr.Perm(10)
	a2[in[0]] += 1
}
