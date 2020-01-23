package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonToMapAndStruct()
	fmt.Println("-------toJson---------")
	mapAndStructToJson()
	fmt.Println("-------jsonToInterface----")
	jsonToInterface()
}
func jsonToMapAndStruct() {
	var str = `{"name":"jack","age":11}`
	var strMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &strMap)
	if err != nil {
		fmt.Println(err)
	}
	if name, ok := strMap["name"]; ok {
		fmt.Println("name", name)
	}

	jsonStr := `{"name":"jack","age":18}`

	type Men struct {
		Name string `json:"name""`
		Age  int    `json:"age"`
	}
	var men = new(Men)
	toStructErr := json.Unmarshal([]byte(jsonStr), men)
	if toStructErr != nil {
		fmt.Println("JsonToStruct err: ", toStructErr)
	}
	fmt.Printf("json to struct:%v\n", *men)

}

func mapAndStructToJson() {
	type User struct {
		Name string `json:"name"` // 记得大写,可导出
		Age  int    `json:"age"`
	}
	var user = User{Name: "jack", Age: 20}
	//var toJsonErr error
	userStr, err := json.Marshal(&user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(userStr))

	var mapStr = map[string]interface{}{"name": "jack", "age": 12}
	str, _ := json.Marshal(mapStr)
	fmt.Println(string(str))

}

func jsonToInterface() {
	var userInfo interface{}
	userInfo = []int{1,2,3,4,5}
	var uu = userInfo.([]int)
	fmt.Println("uu1",uu[1])

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	if err := json.Unmarshal(b, &f); err == nil {
		m := f.(map[string]interface{})
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string", vv)
			case int:
				fmt.Println(k, "is int", vv)
			case float64:
				fmt.Println(k, "is float64", vv)
			case []interface{}:
				fmt.Println(k, "is an array:")
				for i, u := range vv {
					fmt.Println(i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}
}
