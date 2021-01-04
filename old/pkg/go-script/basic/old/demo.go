package main

import (
	"flag"
	"fmt"
)

// 接受命令行参数
// itech8deMacBook-Pro :: gin_blog/app/script ‹master*› » go run flag.go -name="jack"
// hello,jack!
var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "the greeting object")
}
func main() {
	flag.Parse()
	hello(name)
	changename()
	demoA()
	demoB()
}
func hello(name string) {
	fmt.Printf("hello,%s!\n", name)
}

var block = "package"

func changename() {
	block := "function"
	// 程序实体的访问权限有三种：包级私有的、模块级私有的和公开的。
	{
		block := "inner"
		fmt.Printf("The block is %s.\n", block)
	}
	fmt.Printf("The block is %s.\n", block)

}

//上面方法块运行结果
//The block is inner.
//The block is function.
var container = []string{"zero", "one", "two"}

func demoA() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	fmt.Printf("The element is %q.\n", container[1])
}
func demoB() {
	// 请记住，一对不包裹任何东西的花括号，除了可以代表空的代码块之外,
	// 还可以用于表示不包含任何内容的数据结构（或者说数据类型）。

	//用关键字type声明自定义的各种类型
	type MyString = string
	var name MyString
	name = "jack"
	fmt.Printf("name is %s", name)

	s1 := make([]int, 5)
	fmt.Printf("lsThe length of s1: %d\n", len(s1))
	fmt.Printf("The capacity of s1: %d\n", cap(s1))
	fmt.Printf("The value of s1: %d\n", s1)
	s2 := make([]int, 5, 8)
	fmt.Printf("The length of s2: %d\n", len(s2))
	fmt.Printf("The capacity of s2: %d\n", cap(s2))
	fmt.Printf("The value of s2: %d\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("%d\n", s3[1])
}
