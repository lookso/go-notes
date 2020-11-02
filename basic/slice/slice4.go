package main

import "fmt"

func main() {
	fmt.Println("----------slice----------")
	var arr = make([]int, 10, 10)
	fmt.Printf("arr-before:%v,len:%d,cap:%d,point:%p\n", arr, len(arr), cap(arr), &arr)
	// arr:[0 0 0 0 0 0 0 0 0 0],len:10,cap:10,point:0xc000090000
	newArr := arr // 值copy,创建一个新数组
	fmt.Printf("newArr:%v,%p\n", newArr, newArr)
	//newArr:[0 0 0 0 0 0 0 0 0 0],0xc000020050
	arr = []int{1, 2, 3}
	fmt.Printf("arr-after:%v,len:%d,cap:%d,point:%p\n", arr, len(arr), cap(arr), &arr)
	// arr:[1 2 3],len:3,cap:3,point:0xc000090000

	setArr(arr)
	fmt.Printf("update after slice %v,%p\n", arr, arr) //update after slice [2 2 3]
	// go 只有值传递,但是发现没有?是不是觉得很奇怪,为啥update after slice打印返回的结果不是[2 2 3 10]？
	// 原因是切片扩容的时候会导致slice重新分配内存地址
	// append 函数重新创建新的底层数组时,容量会是现有元素的两倍(前提是元素个数小于1000)

	makeAndNew()
	copyArr()
}
func setArr(s []int) {
	fmt.Println("----setArr----")
	s[0]++
	s = append(s, 10)                                                      // 此处容量不足,底层新建数组, 重新会分配内存地址
	fmt.Printf("slice:%v,len:%d,cap:%d,point:%p\n", s, len(s), cap(s), &s) //  slice:[2 2 3 10],len:4,cap:6,point:0xc00000c0e0
}

func makeAndNew() {
	var s = make([]int, 10, 10)
	fmt.Printf("s:%v\n", s)
	type Demo struct {
		Name string `json:"name"`
	}
	fmt.Printf("demo: %v\n", new(Demo))
}
func copyArr() {
	fmt.Println("copyArr")
	var s = []int{1, 2, 3, 4}
	fmt.Printf("%v%p\n",s,s) //[1 2 3 4]0xc000018380
	ns := s
	ns[0]=10
	fmt.Printf("%v%p\n",s,s) //[10 2 3 4]0xc000018380
	var a = make([]int, 3)
	copy(a, s)
	a[0]=11
	fmt.Printf("%v%p\n",a,a) // [11 2 3]0xc0000183a0
}
