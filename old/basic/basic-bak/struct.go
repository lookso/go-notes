/*
@Time : 2019-08-27 09:30
@Author : Tenlu
@File : struct
@Software: GoLand
*/
/*
1.匿名结构体
2.结构体比较,map,struct,
3.
*/
package main

import (
	"fmt"
	"reflect"
)

type People struct{}
type Teacher struct {
	People
}

func main() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}
	type UserInfo struct {
		age  int
		name string
	}
	type UserBase struct {
		age  int
		name string
	}
	u1 := UserInfo{age: 10, name: "jack"}
	u2 := UserInfo{age: 10, name: "jack"}
	u3 := UserBase{age: 10, name: "jack"}
	if u1 == u2 {
		fmt.Println("u1==u2")
	}
	// == 表示 判断2个变量或对象实例是否指向同一个内存空间,equals()表示 判断2个变量或对象实例所指向的内存空间的值是否相同。
	// fmt.Println(u1==u3)
	if reflect.DeepEqual(u1, u3) {
	}
	var myinfo = []int{1, 2, 3}
	var yourinfo = []int{1, 2, 3}
	if reflect.DeepEqual(myinfo, yourinfo) {
		fmt.Println("ok")
	}
	sm1 := struct {
		age int
		m   map[string]string // 不可比较
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	//if sm1 == sm2 {
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}
	a1 := []int{1, 2, 3}
	a2 := []int{1, 2, 3}
	if reflect.DeepEqual(a1, a2) {
		fmt.Println("okok")
	}

	t := Teacher{}
	t.ShowA()

	fmt.Println(live())
	if live() == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
	}

}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}
func (p *People) ShowB() {
	fmt.Println("showB")
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

type MyPeople interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() MyPeople {
	var stu *Student
	return stu
}


// go run -gcflags -m struct.go
