package main

import (
	"fmt"
	"github.com/yuin/gopher-lua"
	"os"
)

func main() {
	L := lua.NewState()
	defer L.Close()
	cwd, _ := os.Getwd()
	fmt.Println(cwd)
	if err := L.DoFile(cwd+"/libs/go-lua/hello.lua"); err != nil {
		panic(err)
	}
	lv := L.Get(-1)
	fmt.Println(lv)
}
