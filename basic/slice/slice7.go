/*
@Time : 2019/5/7 9:11 AM 
@Author : Tenlu
@File : slice
@Software: GoLand
*/
package main

import (
	"fmt"
	"os"
	"reflect"
)

func main()  {
	tt:=[][]interface{}{{"hello"}}
	fmt.Println(tt)
	h,err:=tt[0][0].(string)
	if err==true{
		fmt.Println("是string")
	}
	fmt.Println(h)
	fmt.Println(reflect.TypeOf(h))

	os.Exit(333)

	var a = [...]int{1, 2, 3} // a 是一个数组
	var b = &a                // b 是指向数组的指针
	fmt.Println(a[0], a[1])   // 打印数组的前2个元素
	fmt.Println(b[0], b[1])   // 通过数组指针访问数组元素的方式和数组类似
	for i, v := range b {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}

	var hello=[]string{"h","e","l","l","o"}
	var world=[]string{" ","w","o","r","l","d"}
	hw:=append(hello,world...)

	for _,v:=range hw {
		fmt.Printf("%v",v)
	}
	fmt.Println()

	// myhello
	var my=[]string{"m","y"}
	myHello:=append(my,hello...)
	var myHelloStr string
	for _,m:=range myHello  {
		myHelloStr=myHelloStr+m
	}
	fmt.Printf("myhellostr:%s\n",myHelloStr)

	// 三个点的用法
	setStr()
	//for _,c:=range sliceCopy(){
	//	fmt.Println(c)
	//}
	var strA = "hello"
	strB:=strA
	fmt.Printf("浅拷贝strB:%s\n",strB)
	strB = "world"
	fmt.Printf("浅拷贝strB:%s\n",strB)
	fmt.Printf("浅拷贝strA:%s\n",strA)

	fmt.Printf("深拷贝:%v\n",sliceCopy())
	fmt.Printf("浅拷贝:%v\n",sliceNq())
	
	for i:=0;i<3 ;i++  {
		// 不建议for循环中使用defer
		defer func(i int) {print(i)}(i) //210
		// defer print(i) // 210
		// defer func() {print(i)}() // 333
	}

	var aa=[]interface{}{123,"abc"}
	interfaceSlice(aa...) // 123 abc
	interfaceSlice(aa) // [123 abc]

	fmt.Printf("delsplice:%v\n",delSplice())
	// 向切片中间插入元素
	var sA=[]int{1,2,4,5}
	//sB:=append(sA[:2],[]int{3}...)
	//fmt.Printf("sA:%v\n",sA) // sA:[1 2 3 5]
	// sC:=append(sB,sA[2:]...) //这种方式的是错误的设置方式,sB:[1 2 3 3 5]
	sB:=append(sA[:2],append([]int{3},sA[2:]...)...) // 正确的方式:sB:[1 2 3 4 5]
	fmt.Printf("sB:%v\n",sB)

	//a = append(a, 0)     // 切片扩展1个空间
	//copy(a[i+1:], a[i:]) // a[i:]向后移动1个位置
	//a[i] = x             // 设置新添加的元素
	//
	fmt.Println("----copy实现不创建临时切片来插入元素-----")
	// 比如要向第二个位置插入元素
	var sd=[]int{1,2,4}
	sd = append(sd, 0) // 切片扩展1个空间
	copy(sd[3:], sd[2:])     // a[i:]向后移动1个位置
	sd[2] = 3               // 设置新添加的元素
	fmt.Printf("sd:%v\n",sd)

	fmt.Println("-----切片copy删除元素------")
	var sc =[]int{1,2,3,3,4,5}
	num:=copy(sc[3:],sc[4:]) // 2, 后面的元素copy给前面的元素
	fmt.Printf("copy:%d\n",num)
	sc=sc[:2+num] // // 删除中间1个元素
	fmt.Printf("sd:%v\n",sc) // sd:[1 2 3 4]
	fmt.Println("-------")
	// new与make区别：
	// new只分配内存它并不初始化内存，只是将其置零。new(T)会为T类型的新项目，分配被置零的存储，并且返回它的地址，一个类型为T的值，也即其返回一个指向新分配的类型为T的指针，这个指针指向的内容的值为零（zero value），注意并不是指针为零。比如，对于bool类型，零值为false；int的零值为0；string的零值是空字符串。
	// make用于slice，map，和channel的初始化，返回一个初始化的(而不是置零)，类型为T的值（而不是T）。
	// 之所以有所不同，是因为这三个类型是使用前必须初始化的数据结构。
	// 例如，slice是一个三元描述符，包含一个指向数据（在数组中）的指针，长度，以及容量，在这些项被初始化之前，slice都是nil的。
	// 对于slice，map和channel，make初始化这些内部数据结构，并准备好可用的值。
	//---------------------
	// slice 介绍:https://www.cnblogs.com/junneyang/p/6074786.html

	var x = make([]int,10,20)

	//x=[]int{1,2,3,4,5}
	copy(x,[]int{1,2,3,4,5})
	//len:10,cap:20
	//len:11,cap:20

	//x = []int{1,2,3,4,5}
	//len:5,cap:5
	//len:6,cap:10
	// fmt.Printf("change:%v\n",change(x))
	fmt.Printf("len:%d,cap:%d\n",len(x),cap(x))

	x = change(x)
	fmt.Printf("len:%d,cap:%d\n",len(x),cap(x))
	fmt.Printf("change:x:%d\n",x)
	fmt.Println("----指针数组----")
	myArr := [4]string{"abc", "ABC", "123", "一二三"}
	//查看数组指针的类型和值
	fmt.Printf("%T，%v \n", &myArr, &myArr)
}

func change(x []int) []int{

	return append(x,10)
//	return c
}
func getStr(args ...string) { //可以接受任意个string参数
	for _, v:= range args{
		fmt.Println(v)
	}
}

func setStr(){
	var strss= []string{
		"qwr",
		"234",
		"yui",
		"cvbc",
	}
	getStr(strss...) // 切片被打散传入,迭代
}

func sliceCopy()  []string{
	fmt.Println("---切片深拷贝copy----")
	var sliceA = []string{"hello","world"}
	var sliceB = make([]string,2)
	copy(sliceB,sliceA)
	sliceA=[]string{"alibaba"} // 深拷贝,改变A不会影响B
	return sliceB
}
func sliceNq()[]string  {
	fmt.Println("---切片浅拷贝----")
	var sliceA = []string{"hello","world"}
	var sliceB = sliceA
	sliceA=append(sliceA,sliceB...) // 浅拷贝
	return sliceA
}

func interfaceSlice(a... interface{})  {
	fmt.Println(a...)
}

func delSplice() []int {
	fmt.Println("-----删除切片元素------")
	var a = []int{1, 2, 3}
	a = a[:len(a)-1]   // 删除尾部1个元素
	return a
}


