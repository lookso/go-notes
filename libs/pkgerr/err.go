package main

import (
	E "errors"
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	err := t1()
	fmt.Printf("err :%+v\n", err)
	fmt.Printf("Cause err :%+v\n", errors.Cause(err))
	fmt.Printf("WithStack err :%+v\n", errors.WithStack(err))

	fmt.Println("-------------------")
	err = t5()
	fmt.Printf("err :%+v\n", err)
}

func t1() error {
	return errors.Wrap(t2(), "t1错误")
}
func t2() error {
	return errors.New("基础错误")
}

// 标准库包装错误

func t5() error {
	// 创建一个错误并包装另一个错误
	err := fmt.Errorf("t5错误: %w", t6())
	// 尝试解包错误
	//return err
	return E.Unwrap(err)
}

func t6() error {
	return fmt.Errorf("t6错误:%+w", t7())
}
func t7() error {
	return E.New("t7错误")
}
