package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// 文件相关
// https://studygolang.com/articles/11101?fr=sidebar

func main() {
	p := path.Join("a", "b", "c")
	fmt.Println(p) // a/b/c

	pwd, _ := filepath.Abs("path.go")
	if filepath.IsAbs(pwd) {
		fmt.Println(pwd)
	}
	f, _ := os.Stat("../")
	fmt.Println("args:",filepath.Dir(pwd))
	// Dir返回路径除去最后一个路径元素的部分，即该路径最后一个元素所在的目录。
	// 在使用Split去掉最后一个元素后，会简化路径并去掉末尾的斜杠。
	// 如果路径是空字符串，会返回"."；如果路径由1到多个斜杠后跟0到多个非斜杠字符组成，会返回"/"；其他任何情况下都不会返回以斜杠结尾的路径。

	if f.IsDir() {

		fmt.Println("dir:",path.Dir(pwd))
		dir,_:=os.Getwd()
		fmt.Println("dir:",dir)
	}
	ext:=path.Ext(pwd)
	fmt.Println("文件扩展是:",ext)

	//root := `d:\go_workspace\`
	root,_:=os.Getwd()
	err := filepath.Walk(root, visit)
	fmt.Printf("filepath.Walk() returned %v\n", err)

}

func visit(path string, f os.FileInfo, err error) error {
	fmt.Printf("Visited: %s\n", path)
	return nil
}