package main

import (
	"fmt"
)
/*
1.map 声明和初始化,make和new(),make()长度
2.遍历map
3.判断指定元素是否在map中
4.golang中map不是并发安全
5.map 底层结构 hashtable和searchtree
6.map元素只读,不能更改其属性,且不能获取其地址
7.
*/
func main() {
	// map 是 引用类型的:内存用 make方法来分配,不要使用 new,永远用 make 来构造 map
	// 注意 如果你错误的使用 new() 分配了一个引用对象，你会获得一个空引用的指针，相当于声明了一个未初始化的变量并且取了它的地址

	defer func() {
		fmt.Println("recover:",recover())
	}()
	// 声明
	var userInfo map[string]interface{}
	if userInfo == nil {
		fmt.Println("未初始化的声明的map的值是nil")
	}
	// map赋值前必须初始化
	userInfo = make(map[string]interface{})
	userInfo["func"]= func() int { return 10 }
	userInfo["name"] = "jack"
	userInfo["age"] = 11
	userInfo["sex"] = "男"
	userInfo["hight"] = 169

	// 1.遍历map
	for infoK, infoV := range userInfo {
		fmt.Println(infoK, infoV)
	}
	//2. 元素是否在map中
	if name, ok := userInfo["name"]; ok {
		fmt.Println(name)
	}
	fmt.Println("userInfo's Len:",len(userInfo))

	ReadMap()
	// 变量名可以中文
	var 哈哈="hahaha"
	fmt.Println(哈哈)

}
type Person struct {
	Age int
}
func (p *Person) GrowUp() {
	p.Age++
}

func ReadMap()  {
	m:=map[string]Person{
		"xiaoming":Person{Age:18},
	}
	//m["xiaoming"].Age=20 // golang map 元素属性只读,不可修改
	person:=m["xiaoming1"]
	person.Age=20
	fmt.Println(m) // map[xiaoming:{18}]
	fmt.Println(person) // {20} // 覆盖旧值

}


