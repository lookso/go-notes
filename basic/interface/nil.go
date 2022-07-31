package main

import "fmt"

type Data interface {
	Show()
}
type User struct{}

func (u *User) Show() {}
func main() {
	var d Data = &User{}
	if d == nil { //  interface 只是iface中的data 为nil而已。 但是iface struct{}本身并不为nil,所以不相等
		fmt.Println("ok")
	}
}
