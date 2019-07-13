package main

import "fmt"

func main() {
	data := map[string]interface{}{}
	data["address"] = []int{1, 2, 3, 4, 5}
	if address,ok:=data["address"];ok{
		fmt.Println(address)
		if values,ok:=data["address"].([]int);ok{
			for _,v:=range values{
				fmt.Println("[]int元素值:",v)
			}
		}
	}
	fmt.Println(data)
}
