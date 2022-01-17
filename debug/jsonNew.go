package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var dataStr = `{"code":100,"data":{"id":1473951}}` // 1473954384332480512  // 1.473951e+06]

type BasicResp struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type MyData struct {
	ID int `json:"id"`
}

func main() {
	var m=map[int]string{1:"a",2:"b"}
	if i,ok:=m[1];ok{
		fmt.Println(i)
	}

	var resp BasicResp
	json.Unmarshal([]byte(dataStr),&resp)
	fmt.Println("1111",resp)
	d := json.NewDecoder(bytes.NewReader([]byte(dataStr)))
	d.UseNumber()
	d.Decode(&resp)
	fmt.Printf("resp:%v\n", resp)

	bf := bytes.Buffer{}
	err := json.NewEncoder(&bf).Encode(&resp.Data)
	b := bf.Bytes()

	fmt.Printf("err:%v,b:%v\n", err, string(b)) // err:<nil>,b:{"id":1473954384332480500}

	// 业务层只关心data内容
	var md MyData
	err = json.Unmarshal(b, &md)
	fmt.Printf("err:%+v\n", err)
	fmt.Println("md", md) // md {1473954384332480500}
}
