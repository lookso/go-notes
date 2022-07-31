/*
@Time : 2019/3/1 2:03 PM 
@Author : Tenlu
@File : basic
@Software: GoLand
*/
package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

//  // Len方法返回集合中的元素个数
func (a ByAge) Len() int { return len(a) }

// Swap方法交换索引i和j的两个元素
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less方法报告索引i的元素是否比索引j的元素小
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
	people := []Person{
		{"ob", 31},
		{"aohn", 42},
		{"bichael", 17},
		{"aenny", 26},
	}
	fmt.Println(people)
	sort.Sort(ByAge(people)) // 如果要对复杂对象排序的话，该排序必须包含Len,Swap,Less这三个方法。
	fmt.Println(people)
	// Output:
	//[ob: 31 aohn: 42 bichael: 17 aenny: 26]
	//[bichael: 17 aenny: 26 ob: 31 aohn: 42]

	// 字符串排序
	names := []string{"Hello", "World", "private", "folders", "Users", "workspace"}
	sort.Strings(names)
	fmt.Println(names) // [Hello Users World folders private workspace]
	for _, value := range names {
		fmt.Println(value)
	}
	nums := []int{4, 3, 2, 7, 10, 6}
	sort.Ints(nums)
	fmt.Println(nums)                            // [2 3 4 6 7 10]
	fmt.Println(sort.IntSlice{1, 2, 3, 4}.Len()) // 4
}
