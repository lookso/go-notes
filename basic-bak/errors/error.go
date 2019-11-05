/*
@Time : 2019/2/20 11:21 AM 
@Author : Tenlu
@File : error
@Software: GoLand
*/
package main

import (
	"fmt"
	"errors"
)

type I interface {
	// 有一个方法的接口 I
	Get() Int
}

type Int int

// Int 类型实现了 I 接口
func (i Int) Get() Int {
	return i
}

func echo(request string) (reponse string, err error) {
	if request == "" {
		err = errors.New("empty request")
		return
	}
	reponse = fmt.Sprintf("echo: %s", request)
	return
}
func main() {
	
	for _, r := range []string{"hello", "world", ""} {
		req, err := echo(r)
		if err != nil {
			fmt.Printf("when is %s err:%s\n", req, err)
			continue
		}
		fmt.Println(req)
	}
	// 逗号ok模式检测判断map中是否含有指定key
	x := map[string]string{"one": "2", "two": "", "three": "3"}
	if _, ok := x["four"]; !ok { // !ok不存在
		fmt.Println("key four is not exists")
	}

	var myint Int = 5
	var inter I = myint // 变量赋值给接口
	val, ok := inter.(Int)
	fmt.Printf("%v, %v\n", val, ok) // 输出为：5，true

	fmt.Println("Enter function main.")
	// 无论函数结束执行的原因是啥,其中defer函数调用都会在他即将结束执行的那一刻执行,即便导致他执行结束的原因是一个panic也会这样
	// 联用defer语句和recover函数恢复一个已经发生的panic

	// panic的作用就是抛出一条错误信息，从它的参数类型可以看到它可以抛出任意类型的错误信息。在函数执行过程中的某处调用了panic，则立即抛出一个错误信息。
	// 同时函数的正常执行流程终止，但是该函数中panic之前定义的defer语句将被依次执行。之后该goroutine立即停止执行。

	// recover()用于将panic的信息捕捉。recover必须定义在panic之前的defer语句中。
	// 在这种情况下,当panic被触发时,该goroutine不会简单的终止，而是会执行在它之前定义的defer语句

	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// 引发 panic。
	panic(errors.New("something wrong"))
	fmt.Println("Exit function main.")
}

