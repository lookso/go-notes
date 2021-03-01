package main

import "fmt"

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var arr = make([]*User, 0)
	var u = User{
		Name: "jack",
		Age:  11,
	}
	arr = append(arr, &u)
	for _, v := range arr {
		v.Name = "hh"
	}
	fmt.Println(arr[0].Name)

}
