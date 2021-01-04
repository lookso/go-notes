package main

import "fmt"

func main() {
	const (
		a = iota
		b = 10
		c
		Name = "jack"
		d
		e
		_
		f = iota
		h
	)
	fmt.Println(a, b, c, d, e, f, h)
	// 0 10 10 jack jack 7 8
}
