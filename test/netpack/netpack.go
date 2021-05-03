package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	var info interface{}=123
	fmt.Println(cast.ToInt("123"))
	fmt.Println(cast.ToString(info))
}
