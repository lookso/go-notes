/*
@Time : 2019/3/24 10:33 AM
@Author : Tenlu
@File : go_slice
@Software: GoLand
*/
package main

import "fmt"

func main() {

	s := make([]string, 3,3)
	fmt.Println("emp:", s)
	fmt.Printf("len:%d,cap:%d:\n",len(s),cap(s))
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Printf("len:%d,cap:%d:\n",len(s),cap(s))

	// Slice 也可以被 copy。这里我们创建一个空的和 s 有相同长度的 slice c，并且将 s 复制给 c。

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)
	// Slice 支持通过 slice[low:high] 语法进行“切片”操作。例如，这里得到一个包含元素 s[2], s[3],s[4] 的 slice。

	l := s[2:5]
	fmt.Println("sl1:", l)
	//这个 slice 从 s[0] 到（但是包含）s[5]。

	l = s[:5]
	fmt.Println("sl2:", l)
	// 这个 slice 从（包含）s[2] 到 slice 的后一个值。

	l = s[2:]
	fmt.Println("sl3:", l)
	// 我们可以在一行代码中声明并初始化一个 slice 变量。

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	// Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同。

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

	var st []int
	printSlice(st)

	// append works on nil slices.
	st = append(st, 0)
	printSlice(st)

	// The slice grows as needed.
	st = append(st, 1)
	printSlice(st)

	// We can add more than one element at a time.
	st = append(st, 2, 3, 4)
	printSlice(st)

	goSlice()

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


func goSlice(){
	// 第一种声明方式
	start :=make([]int,0,0)
	start=append(start,100)
	start=append(start,200)
	start=append(start,300)
	fmt.Printf("start:%v,len:%d,cap:%d\n",start,len(start),cap(start))
	// start:[100 200，300],len:3,cap:3
	// 切片的长度可以在运行时修改,最小为 0 最大为相关数组的长度:切片是一个长度可变的数组。
	// 0 <= len(s) <= cap(s)

	fmt.Printf("%v\n",start[:len(start)-1])

	fmt.Printf("%v\n",start)

	// 第二种声明方式
	var language []string
	// 一个切片在未初始化之前默认为 nil，长度为 0
	fmt.Printf("language:%v,len:%d,cap:%d\n",language,len(language),cap(language))
	if language==nil {
		fmt.Printf("Slice language is empty\n")
	}
	language = append(language, "php","python","go","java")

	for k,v:=range language  {
		fmt.Printf("%d->%s\n",k,v)
	}
	iCan:=language[0:3]
	fmt.Printf("%v\n",iCan)

	fmt.Printf("language1:%v\n",language[3:])

	index := 2
	language = append(language[:index], language[index+1:]...)
	fmt.Printf("删除切片指定索引的元素:%v\n",language)

	index = 1
	temp := append([]string{}, language[index:]...)
	language = append(language[:index], "lua")
	fmt.Printf("language:%v\n",language)
	language = append(language, temp...)
	fmt.Printf("在切片中间插入元素:%s\n", language)

	language = language[:len(language)-1]
	fmt.Printf("删除切片最后一个元素:%s\n", language)

	lang:=make([]string,len(language),(cap(language))*2)
	// var lang []string
	var copied = 0
	copied=copy(lang,language)
	fmt.Printf("切片复制:lang:%v,copied:%d\n",lang,copied)
}

