package main

import "fmt"

type Lili struct {
	Name string
}

func (Lili *Lili) fmtPointer() {
	fmt.Println("poniter")
}

func (Lili Lili) fmtReference() {
	fmt.Println("reference")
}
func main() {
	var li Lili
	li.fmtPointer()
	li.fmtReference()

	//  cannot call pointer method on Lili literal

	// Lili{}.fmtReference() // ok
	// Lili{}.fmtPointer() // 报错,为啥？
}
