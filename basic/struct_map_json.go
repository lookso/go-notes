package main

import (
	"sync"
	"encoding/json"
	"fmt"
	"errors"
)

/*
1.struct to json
2.json to struct
3.map to json
4.json to map
*/

func main() {
	var syncGroup sync.WaitGroup
	syncGroup.Add(1)
	go mapToJson(&syncGroup)
	go jsonToMap(&syncGroup)
	go structToJson(&syncGroup)
	go jsonToStruct(&syncGroup)
	syncGroup.Wait()
	defer func() {
		if errMsg := recover(); errMsg != nil {
			fmt.Println(recover())
			fmt.Println(errMsg)
		}
	}()
}

func mapToJson(syncGroup *sync.WaitGroup) (err error){
	defer func() {
		if errMsg := recover(); errMsg != nil {
			fmt.Println(errMsg)
		}
	}()
	var userInfoMap = map[string]interface{}{"name": "toms", "age": 120, "sex": 1, "height": 180}

	userInfoJson, errMsg := json.Marshal(userInfoMap)
	if errMsg != nil {
		errors.New("map to json err")
	}
	fmt.Println("map to json",string(userInfoJson))
	syncGroup.Done()
	return
}
func jsonToMap(syncGroup *sync.WaitGroup) (err error){
	defer func() {
		if errMsg := recover(); errMsg != nil {
			fmt.Println(errMsg)
		}
	}()
	jsonStr := `{"runtimekey":"123","sharetimekey":"456","today":"2018-12-12 20:06:39","userid":"1000"}`
	mapData := map[string]string{}
	errMsg := json.Unmarshal([]byte(jsonStr), &mapData)
	if errMsg != nil {
		errors.New("json to map err")
	}
	if today, ok := mapData["today"]; ok {
		fmt.Println("date:", today)
	}
	syncGroup.Wait()
	return

}

type UserInfo struct {
	Name   string
	Age    int
	Sex    int
	Height int
}

func structToJson(syncGroup *sync.WaitGroup) (err error){
	defer func() {
		if errMsg := recover(); errMsg != nil {
			fmt.Println(errMsg)
		}
	}()
	userInfo := new(UserInfo)
	userInfo.Name = "Jack"
	userInfo.Age = 120
	userInfo.Sex = 1
	userInfo.Height = 180
	data,errMsg:=json.Marshal(&userInfo)
	if errMsg!=nil{
		errors.New("struct to json err")
	}
	fmt.Println("struct to json:",string(data))
	syncGroup.Wait()
	return

}

func jsonToStruct(syncGroup *sync.WaitGroup) (err error){

	//var userInfo=new(UserInfo)
	var userInfo UserInfo
	jsonStr:=`{"Name":"Jack","Age":120,"Sex":1,"Height":180}`
	errMsg:=json.Unmarshal([]byte(jsonStr),&userInfo)
	if errMsg!=nil{
		errors.New("json to struct err")
	}
	fmt.Println("json to struct:",userInfo.Name)
	syncGroup.Wait()
	return
}