package main

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed p/hello.txt
var s string

//go:embed p/world.txt
//go:embed p/hello.txt
var f embed.FS

//go:embed p/*
var p embed.FS

func main() {
	fmt.Println(s)

	data, _ := f.ReadFile("p/hello.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("p/world.txt")
	fmt.Println(string(data))
	fmt.Println("----p-----")
	data, _ = f.ReadFile("p/world.txt")
	fmt.Println(string(data))
	dirEntries, _ := f.ReadDir("p") // 遍历目录
	for _, de := range dirEntries {
		fmt.Println(de.Name(), de.IsDir())
	}
}
