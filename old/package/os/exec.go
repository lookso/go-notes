/*
@Time : 2019-10-02 21:22 
@Author : Tenlu
@File : exec
@Software: GoLand
*/
package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main(){
	fmt.Println(exec.ErrNotFound)
	// 在环境变量PATH指定的目录中搜索可执行文件，如file中有斜杠，则只在当前目录搜索。
	// 返回完整路径或者相对于当前目录的一个相对路径
	path, err := exec.LookPath("go")
	if err != nil {
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n", path)
}
