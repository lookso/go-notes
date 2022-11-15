package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error
	//errors.Warn()
	err = errors.New("upload data err")
	// Go1.13版本为fmt.Errorf函数新加了一个%w占位符用来生成一个可以包裹Error的Wrapping Error。
	w := fmt.Errorf("wrap了一个错误:%w", err)
	fmt.Println(w)
	fmt.Println(errors.Unwrap(w))
	if errors.Is(err,w){
		fmt.Println("err equal")
	}
	var perr *os.PathError
	if errors.As(err, &perr) {
		fmt.Println(perr.Path)
	}

	// Go1.13 错误处理:https://www.cnblogs.com/sunsky303/p/11571440.html
}
