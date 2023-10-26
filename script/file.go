package main

import (
	"fmt"
	"github.com/spf13/cast"
)

func main() {
	call:="181"
	fmt.Println(cast.ToInt(string(call[2])),len(call))
	return
}
