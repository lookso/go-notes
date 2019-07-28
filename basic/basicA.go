/*
@Time : 2019-07-26 21:08 
@Author : Tenlu
@File : basicA
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
	"unsafe"
)

type Test struct {
	name string
	age  int
}

func main() {
	a := "abc"
	aa := []rune(a[0:])

	fmt.Println("a", &a)
	fmt.Println("aa", aa)

	fmt.Println("a arr", a[0:1])
	for index, data := range a {
		fmt.Printf("%d:%c\n", index, data)
	}
	s1 := []byte(a)
	fmt.Println("s1 arr", s1)
	for index, data := range s1 {
		fmt.Printf("byte:%d:%c\n", index, data)
	}
	b := "abc你好"
	s2 := []rune(b)
	fmt.Println("s2 arr", s2)
	for index, data := range s2 {
		fmt.Printf("rune:%d:%c\n", index, data)
	}
	fmt.Println("a len", len([]rune(a))) // 5
	fmt.Println("a len", len(a))         // 9
	fmt.Println(len([]rune(a)))
	_ = len(a)
	str := "abc"
	fmt.Println(unsafe.Sizeof(str))
	const c = unsafe.Sizeof(a)
	fmt.Println(c)

	if a != "" {

	}
	fmt.Println("-------------")
	at := []int{0, 1, 2}
	fmt.Printf("%p\n", &at)
	for i, v := range at {
		if i == 0 {
			at[1], at[2] = 999, 999
			fmt.Printf("%p\n", &at)
			fmt.Println(at)
		}
		fmt.Println(v)
		at[i] = int(v) + int(100)
	}
	fmt.Println(at)
	fmt.Printf("%p\n", &at)
	// index、value 都是从复制品中取出。 // 在修改前，我们先修改原数组。
	// 确认修改有效，输出 [0, 999, 999]。
	// 使⽤用复制品中取出的 value 修改原数组。
	// 输出 [100, 101, 102]。
	fmt.Println("-------------")
	fmt.Println("unsafe pointer", unsafe.Pointer(&str))

	fmt.Println("test struct size", unsafe.Sizeof(Test{}))

	fmt.Println("--------")
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(f)
	}
	defer f.Close()

	type UserInfo struct {
		name string
	}
	var user UserInfo
	fmt.Println(user.name)

	var tt = map[UserInfo]string{UserInfo{name: "jack"}: "bj"}
	fmt.Println(tt)

	ttt := []UserInfo{{name: "jack"}, {name: "toms"}}
	fmt.Println(ttt)

	ff:=signFunc
	fmt.Println(ff(11))
}


func signFunc(age int) int  {
	return age
}
