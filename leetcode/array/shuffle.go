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
