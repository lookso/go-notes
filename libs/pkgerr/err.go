package main

import (
	E "errors"
	"fmt"
	"github.com/pkg/errors"
)

func main() {
	err0 := t1()
	//fmt.Println()

	//err := errors.Wrap(err0, "附加信息")
	//err=errors.WithMessage(err0,"nb")
	err:=errors.WithStack(err0)
	//err=errors.Cause(err)
	if err != nil {
		//打印错误需要%+v才能详细输出
		fmt.Printf("err :%+v\n", err)
		return
	}

	fmt.Println("Hello world")
}

func t1() error {
	return E.New("错误")
}