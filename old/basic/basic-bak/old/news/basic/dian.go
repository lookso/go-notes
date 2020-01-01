package main

import "fmt"

// golang 三个点...的用法

func main()  {
	// 设置数组长度
	var  arr=[...]string{"name","sex","age","address","height"}
	fmt.Println(len(arr))

	sliceData:=make([]string,len(arr))
	sliceData = arr[0:4]
	fmt.Println(sliceData)
	test(sliceData...)// 迭代切片元素,不能迭代数组元素

}
// 可以接受任意个string参数
func test(args ...string)  {
	for _, v:= range args{
		fmt.Println(v)
	}
}

