package main

import (
	"fmt"
)

func main() {
	data := map[string]interface{}{}
	data["address"] = []int{1, 2, 3, 4, 5}
	if address, ok := data["address"]; ok {
		fmt.Println(address)
		if values, ok := data["address"].([]int); ok {
			for _, v := range values {
				fmt.Println("[]int元素值:", v)
			}
		}
	}
	fmt.Println(data)
	
	fmt.Println("-------------------------")

	var num32 int32 = 12
	var md = map[string]interface{}{"name": 10}
	checkType(1, num32, md)

}

func checkType(items ...interface{}) {
	for index, v := range items {
		switch v.(type) {
		case int:
			fmt.Printf("第 %d 个参数 is int\n", index)
		case int32:
			fmt.Printf("第 %d 个参数 is int32\n", index)
		case float32:
			fmt.Printf("第 %d 个参数 is float32\n", index)
		case map[string]interface{}:
			fmt.Printf("第 %d 个参数 is map[string]interface{}\n", index)
		}

		if dv, ok := v.(map[string]interface{}); ok {
			if ddv, ook := dv["name"]; ook {
				fmt.Println("name k exists")
				if qv, qok := ddv.(int); qok { // //转成int，带检查
					fmt.Println("v is int:", qv)
				}
			}
		}
	}

	//[1 2 3 4 5]
	//[]int元素值: 1
	//[]int元素值: 2
	//[]int元素值: 3
	//[]int元素值: 4
	//[]int元素值: 5
	//map[address:[1 2 3 4 5]]
	//-------------------------
	//第 0 个参数 is int
	//第 1 个参数 is int32
	//第 2 个参数 is map[string]interface{}
	//name k exists
	//v is int: 10


}
