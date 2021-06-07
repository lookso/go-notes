package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int{2, 4, 1, 6, 3, 1, 10, 8}
	sort.Ints(arr)
	fmt.Println(arr)
	// 从大到小排
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
	fmt.Println(arr)
	// 因为这三种类型也实现了 sort.Interface 接口
	sortReverse()
	sortSliceStable()
	sortSort()

}

// sortTest
type User = struct {
	Name string
	Age  int
}
type Users []User // 通过对age排序实现了sort.Interface接口

var us = []User{
	{Name: "马云", Age: 67},
	{Name: "李嘉诚", Age: 80},
	{Name: "马化腾", Age: 50},
}

func (u Users) Len() int {
	return len(u) // 待排序的元素个数
}
func (u Users) Less(i, j int) bool {
	return u[i].Age > u[j].Age // 第i个和第j个元素比较
}
func (u Users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i] //第i个和第j个元素交换
}

func sortSort() {
	fmt.Println("----------sortSort----------")
	sort.Sort(Users(us))
	fmt.Println("us", us)
}

// sort.Reverse
type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func sortReverse() {
	dd := make(intSlice, 0)
	dd = append(dd, 2, 4, 1, 6, 3, 1, 10, 8)
	sort.Sort(sort.Reverse(dd))
	fmt.Println("sort.Reverse", dd)
}

// sort.SliceStable
var family = []struct {
	Name string
	Age  int
}{
	{"Alice", 23},
	{"David", 2},
	{"Eve", 2},
	{"Bob", 25},
}

func sortSliceStable() {
	fmt.Println("-------sortSliceStable-------")
	// 用 age 排序,年龄相等的元素保持原始顺序
	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})
	fmt.Println(family) // [{David 2} {Eve 2} {Alice 23} {Bob 25}]
}
