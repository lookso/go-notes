package main

import "fmt"

func main() {
	var theMap map[string]string
	if theMap == nil {
		fmt.Println("map未初始化前其值是nil")
	}
	var theChannel chan int
	if theChannel == nil {
		fmt.Println("channel未初始化前其值是nil")
	}
	var theInterface interface{}
	if theInterface == nil {
		fmt.Println("interface未初始化前其值是nil")
	}
	var theSlice []int
	if theSlice == nil {
		fmt.Println("slice未初始化前其值是nil")
	}
	if theSlice == nil {
		fmt.Println("ok")
	}
	//if theSlice == theMap {
	//	fmt.Println("ok") // 两个类型属性值不一样
	//}
	fmt.Println(*newInt())

	var str="name"
	pointStr:=&str
	fmt.Println("string point:",pointStr)
}
func newInt() *int {
	intPoint:=new(int)
	fmt.Println("类型指针地址:",intPoint)

	a := 3
	fmt.Println(a)
	var ip *int /* 声明指针变量 */
	ip=&a
	var mapData=map[string]string{"name":"jack"}
	if name,ok:=mapData["name"];ok{
		fmt.Println("name:",name)
	}
	fmt.Println(ip)
	return &a
}
