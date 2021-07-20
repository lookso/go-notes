package main

import "fmt"

func main() {
	const (
		a = iota
		b = 10
		c
		Name = "jack"
		d    = 11
		e
		_
		f = iota
		h
	)
	fmt.Println(a, b, c, d, e, f, h)
	// 0 10 10 11 11 7 8
	// 从上面可以更清晰的看出iota实际上是遍历const块的索引，每行中即便多次使用iota，其值也不会递增。

	// git 2
}
