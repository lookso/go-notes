package main

import (
	"fmt"
	"strconv"

	"github.com/samber/lo"
)

type Person struct {
	Name string
	Age  int
}

// 类似的标准库 go 1.22 slices 和maps 的函数式操作

func main() {
	//将切片中的每个元素应用一个函数，并返回一个新的切片：
	result := lo.Map([]int{1, 2, 3}, func(item int, index int) int {
		fmt.Println(item, index)
		return item * 2
	})
	fmt.Println("map", result)
	// result: [2, 4, 6]
	// 过滤切片，只保留满足条件的元素：
	r1 := lo.Filter([]int{1, 2, 3, 4}, func(x int, _ int) bool {
		return x%2 == 0
	})
	fmt.Println("filter", r1)
	// r1: [2, 4]
	// 将切片中的元素组合起来，生成一个单一的结果
	r2 := lo.Reduce([]int{1, 2, 3, 4}, func(acc int, x int, _ int) int {
		return acc + x
	}, 0)
	fmt.Println("Reduce", r2)
	// 检查切片或映射是否包含一个元素
	fmt.Println("Contains", lo.Contains([]int{1, 2, 3}, 2))
	// 返回一个去重后的切片
	fmt.Println("Uniq", lo.Uniq([]int{1, 2, 2, 3}))
	// 将切片分割成多个大小为 n 的块
	fmt.Println("Chunk", lo.Chunk([]int{1, 2, 3, 4, 5}, 2))
	// 根据给定的函数对切片进行分组：
	grouped := lo.GroupBy([]int{1, 2, 3, 4}, func(x int) string {
		if x%2 == 0 {
			return "even"
		}
		return "odd"
	})
	// grouped: map[even:[2 4] odd:[1 3]]
	fmt.Println("grouped", grouped)

	fmt.Println("Min", lo.Min([]int{1, 2, 2, 3})) // Min 1

	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Cathy", 29},
	}
	// 找最小的值
	minPerson := lo.MinBy(people, func(a, b Person) bool {
		return a.Age < b.Age
	})
	// {Bob 25}
	fmt.Println("minPerson", minPerson)

	//-----map------------------
	originalMap := map[string]int{"a": 1, "b": 2, "c": 3}

	// MapKeys 示例
	mappedKeys := lo.MapKeys(originalMap, func(v int, k string) string {
		return k + "_new"
	})
	// MapKeys map[a_new:1 b_new:2 c_new:3]
	fmt.Println("MapKeys", mappedKeys)

	// MapValues 示例
	mappedValues := lo.MapValues(originalMap, func(v int, k string) int {
		return v * 2
	})
	// MapValues map[a:2 b:4 c:6]
	fmt.Println("MapValues", mappedValues)

	// FilterMap 示例
	list := []int64{1, 2, 3, 4}

	filterMap := lo.FilterMap(list, func(nbr int64, index int) (string, bool) {
		return strconv.FormatInt(nbr*2, 10), nbr%2 == 0
	})
	// filterMap:[4 8]
	fmt.Printf("filterMap:%v\n", filterMap)

	// Reduce 示例
	reduce := lo.Reduce(list, func(agg int64, item int64, index int) int64 {
		return agg + item
	}, 0)
	//reduce:10
	fmt.Printf("reduce:%v\n", reduce)

	// Keys 示例
	keys := lo.Keys(originalMap)
	// Keys [a b c]
	fmt.Println("Keys", keys)

	// Values 示例
	values := lo.Values(originalMap)
	// Values [1 2 3]
	fmt.Println("Values", values)

}
