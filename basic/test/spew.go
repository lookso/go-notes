package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
)

type Demo struct {
	a int
	b string
}

func main() {
	arr := [...]*Demo{{100, "Python"}, {200, "Golang"}}
	fmt.Printf("%v\n-------------------哈哈哈哈分割线-----------\n", arr)
	spew.Dump(arr)
}
