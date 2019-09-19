package main

import (
	"fmt"
	"os"
)

func main() {
	// 获取环境变量,没有设置为空
	fmt.Println(os.Getenv("GOPROXY"))
	// 获取环境变量,返回值和布尔类型,没有设置返回false
	path, statue := os.LookupEnv("GOPATH")
	if statue {
		fmt.Println(path)
	}
	// 设置环境变量
	os.Setenv("GOGO", "baidu")
	fmt.Println(os.Getenv("GOGO"))
	fmt.Println("current pid", os.Getpid())
}
