package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonToMapAndStruct()
	fmt.Println("-------toJson---------")
	mapAndStructToJson()
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
	var men =new(Men)
	toStructErr := json.Unmarshal([]byte(jsonStr), men)
	if toStructErr != nil {
		fmt.Println("JsonToStruct err: ",toStructErr)
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
