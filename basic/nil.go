package main

import (
	"fmt"
	"reflect"
)

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

	var cupA interface{}
	cupA="123"
	cupB:=cupA
	fmt.Printf("cupA:%p,cupB:%p\n",&cupA,&cupB)
	fmt.Println("接口反射类型:",reflect.TypeOf(cupA)) // 接口反射类型: string

	yourName:="jackma"
	myName:=yourName
	fmt.Println("yourName point addr:",&yourName," myName point addr:",&myName)
	fmt.Printf("yourName point addr:%p\n",&yourName)

	var yourInfo = make([]string,50)
	//yourInfo=[]string{"name","age","sex"}
	yourInfo[0]="name"

	yourInfoA:=yourInfo[0:2]
	myInfo:=yourInfo // //其实就是浅拷贝,a和b同时指向同一个地址0xc00006c180 通过yourInfo和myInfo都可以改变值，myInfo值变了yourInfo的值也变了
	myInfo[0]="name_jack"
	fmt.Printf("yourInfo value:%v\n",yourInfo)
	fmt.Printf("yourInfo point addr:%p,myInfo point add:%p,yourInfoA point addr:%p\n",yourInfo,myInfo,yourInfoA)

	heInfo:=append(yourInfo,"hieght")
	fmt.Printf("heInfo point addr:%p\n",heInfo)
	var uInfo=make([]string,30)
	copy(uInfo,yourInfo)
	fmt.Printf("uInfo point addr:%p\n",uInfo)


	var f func(int) int
	if f==nil{
		fmt.Println("func nil")
	}
	fmt.Printf("func point addr:%p\n", f)
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


