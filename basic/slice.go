package main

import "fmt"

func main() {
	fmt.Println("----------slice----------")
	var arr = make([]int, 10, 10)
	fmt.Printf("arr-before:%v,len:%d,cap:%d,point:%p\n", arr, len(arr), cap(arr), &arr)
	// arr:[0 0 0 0 0 0 0 0 0 0],len:10,cap:10,point:0xc000090000
	//a := arr
	//fmt.Printf("a:%v,%p\n", a, &a)
	arr = []int{1, 2, 3}
	fmt.Printf("arr-after:%v,len:%d,cap:%d,point:%p\n", arr, len(arr), cap(arr), &arr)
	// arr:[1 2 3],len:3,cap:3,point:0xc000090000

	setArr(arr)
	fmt.Println("update after slice", arr) //update after slice [2 2 3]

	// go 只有值传递,但是发现没有?是不是觉得很奇怪,为啥第二个打印返回的结果不是[3 2 3 10]？
	// 原因是切片扩容的时候会导致slice重新分配内存地址
	// append 函数重新创建新的底层数组时,容量会是现有元素的两倍(前提是元素个数小于1000)
	// 如果元素个数超过1000，那么容量会以 1.25 倍来增长
	st()
}
func setArr(s []int) {
	fmt.Println("----setArr----")
	s[0]++
	s = append(s, 10)                                                      // 此处容量不足,底层新建数组, 重新会分配内存地址
	fmt.Printf("slice:%v,len:%d,cap:%d,point:%p\n", s, len(s), cap(s), &s) //  slice:[2 2 3 10],len:4,cap:6,point:0xc00000c0e0
	s[0]++                                                                 // 通过下标设置会改变内存地址
	fmt.Printf("set slice++%v,point:%p\n", s, &s)                          //set slice [3 2 3 10]
	fv()
}

func st() {
	var s = make([]int, 10, 10)
	printSlice(s)
	s = []int{1, 2, 3}
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d point:%p,s:%v\n", len(s), cap(s), &s, s)
}

func fv() {
	var a = []int{1, 2, 3}
	fmt.Printf("fv:%v,%p,%p,%p\n", a, a, &a[0], &a[1])
	a[0] = 10
	fmt.Printf("fv:%v,%p,%p,%p\n", a, a, &a[0], &a[1])
	//fv:[1 2 3],0xc000018380,0xc000018380,0xc000018388
	//对于Slice类型的数据，返回的一直是指向第一个元素的地址，所以我们通过fmt.Printf()中%p来打印Slice的地址,
	//其实打印的结果是内部存储数组元素的首地址
}
