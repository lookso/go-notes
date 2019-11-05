/*
@Time : 2019/1/6 2:04 PM
@Author : Tenlu
@File : hello
@Software: GoLand
*/
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// 进程运行时PID是由操作系统随机分配的
	fmt.Printf("当前进程ID:%d,当前父进程ID:%d\n", os.Getpid(), os.Getppid())
	fmt.Printf("-------------os.Args-----------\n")
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	// ./hello jack lucy toms
	fmt.Printf("-------------flag-----------\n")
	wordPtr := flag.String("word", "foo", "a string")
	numbPtr := flag.Int("numb", 42, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")
	flag.Parse()
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	//word: foo
	//numb: 42
	//fork: false
	//svar: bar
	//tail: [jack lucy toms]

	// go run hello.go -word baidu -numb 100 -fork true -svar robin
	//-------------flag-----------
	//word: baidu
	//numb: 100
	//fork: true
	//svar: bar
	//tail: [true -svar robin]

}
