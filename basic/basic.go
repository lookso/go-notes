package main

import (
	"fmt"
	"flag"
)

/*
1.for循环
2.匿名函数

*/
//  go run basic.go --name jack
var name string
func init() {
	flag.StringVar(&name, "name", "everyone", "the greeting object")
}
func main()  {
	flag.Parse()
	hello(name)
	sum := 0
	// 1.类似while循环
	for {
		sum ++
		if sum > 10 {
			break
		} else {
			fmt.Println(sum)
		}
	}

	// 2.匿名函数
	defer func() {
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		fmt.Printf("高斯函数之和:%d\n", sum)
	}()

	nextInt := intSeq()
	fmt.Println("nextInt:",nextInt())
}

// 返回一个函数类型值
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func hello(name string) {
	fmt.Printf("hello,%s!\n", name)
}
