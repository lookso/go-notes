/*
@Time : 2019-07-23 21:04 
@Author : Tenlu
@File : order
@Software: GoLand
*/
package main

import (
	"fmt"
	"sort"
)
// 保证:go map 顺序输出
func main() {
	/* 声明一个字符串切片，存储map的key值 */
	var Name []string
	Name = append(Name, "Bob", "Andy", "Clark", "David", "Ella")

	/* 声明索引类型为字符串的map */
	var Person = make(map[string]string)
	Person["Bob"] = "B"
	Person["Andy"] = "A"
	Person["Clark"] = "C"
	Person["David"] = "D"
	Person["Ella"] = "E"

	fmt.Println("未排序输出:")
	for key, value := range Person {
		fmt.Println(key, ":", value)
	}
	/* 对slice数组进行排序，然后就可以根据key值顺序读取map */
	sort.Strings(Name)
	fmt.Println("排序输出:")
	for _, Key := range Name {
		/* 按顺序从MAP中取值输出 */
		if Value, ok := Person[Key]; ok {
			fmt.Println(Key, ":", Value)
		}
	}
}
