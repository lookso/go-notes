package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		Name="jack"
		d
		e
		_
		f=iota
		h
		
	)
	fmt.Println(c,d,e,f,h)

}
