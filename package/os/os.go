package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	getEnvInfo()
	fmt.Println("-----------file----------")
	doFile()

}
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func getEnvInfo() {
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
	fmt.Println(os.Hostname())
}
func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}
func doFile() {

	err := os.Mkdir("os-dir", 777)
	if err != nil {
		fmt.Println(err)
	}
	var wireteString = "test\n"
	var filename = "os-file.txt"
	var f *os.File
	var err1 error
	/****第一种方式: 使用 io.WriteString 写入文件 ******/
	if checkFileIsExist(filename) { //如果文件存在
		f, err1 = os.OpenFile(filename, os.O_APPEND, 0666) // 打开文件
		// os.Open 只能读取内容
		fmt.Println(err1)
		fmt.Println("文件存在")
	} else {
		f, err1 = os.Create(filename) //创建文件
		fmt.Println("文件不存在")
	}
	check(err1)
	n, err1 := io.WriteString(f, wireteString) //写入文件(字符串)
	check(err1)
	fmt.Printf("写入 %d 个字节n", n)
}
