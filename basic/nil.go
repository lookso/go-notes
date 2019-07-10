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

	yourName:="jackma"
	myName:=yourName
	fmt.Println("yourName point add:",&yourName," myName point add:",&myName)
	fmt.Printf("yourName point add:%p\n",&yourName)

	var yourInfo = make([]string,50)
	//yourInfo=[]string{"name","age","sex"}
	yourInfo[0]="name"

	yourInfoA:=yourInfo[0:2]
	myInfo:=yourInfo // //其实就是浅拷贝,a和b同时指向同一个地址0xc00006c180 通过yourInfo和myInfo都可以改变值，myInfo值变了yourInfo的值也变了
	myInfo[0]="name_jack"
	fmt.Printf("yourInfo value:%v\n",yourInfo)
	fmt.Printf("yourInfo point add:%p,myInfo point add:%p,yourInfoA point add:%p\n",yourInfo,myInfo,yourInfoA)

	heInfo:=append(yourInfo,"hieght")
	fmt.Printf("heInfo point add:%p\n",heInfo)
	var uInfo=make([]string,30)
	copy(uInfo,yourInfo)
	fmt.Printf("uInfo point add:%p\n",uInfo)
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
