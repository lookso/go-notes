package main

import "fmt"

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
}
