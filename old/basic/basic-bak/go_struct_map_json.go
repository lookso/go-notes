/*
@Time : 2018/12/12 8:02 PM
@Author : Tenlu
@File : struct_map_json
@Software: GoLand
*/
package main

import (
	"encoding/json"
	"fmt"
	"github.com/goinggo/mapstructure"
	"log"
	"time"
)

/**
@json和map互转以及json和struct相互转换
*/
func main() {
	MapToJson()
	JsonToMap()
	StructToJson()
	JsonToStruct()
	MapToStruct()
	HandleJson()

}

func MapToJson() {
	data := map[string]string{
		"userid":       "1000",
		"runtimekey":   "123",
		"sharetimekey": "456",
		"today":        time.Now().Format("2006-01-02 15:04:05"),
	}
	outputData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("mapToJson err: ", err)
	}
	fmt.Printf("%v\n", outputData) // 默认输出的是对应的ascii
	fmt.Printf("map to json:%v\n", string(outputData))
}

func JsonToMap() {
	jsonStr := `{"runtimekey":"123","sharetimekey":"456","today":"2018-12-12 20:06:39","userid":"1000"}`
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMap err: ", err)
	}
	fmt.Printf("json to map:%v\n", mapResult)
}

// struct和json 相互转换
// 同一个包里面不能包含相同的结构体和方法
type Men struct {
	Name string `json:"name_title"`
	Age  int    `json:"age_size"`
}

// 在结构体中引入tag标签,这样匹配的时候json串对应的字段名需要与tag标签中定义的字段名匹配,
// 当然tag中定义的名称不需要首字母大写，且对应的json串中字段名仍然大小写不敏感。
// 此时,结构体中对应的字段名可以不用和匹配的一致,但是首字母必须大写，只有大写才是可对外提供访问的

// 注意json里面的key和struct里面的key要一致,struct中的key的首字母必须大写,而json中大小写都可以。
func StructToJson() {
	p := Men{
		Name: "jackma",
		Age:  18,
	}
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	// struct to json:{"name_title":"jackma","age_size":18}
	fmt.Printf("struct to json:%v\n", string(jsonBytes))
}

func JsonToStruct() {
	jsonStr := `{"name_title":"jackma","age_size":18}`

	var men Men
	err := json.Unmarshal([]byte(jsonStr), &men)
	if err != nil {
		fmt.Println("JsonToStruct err: ", err)
	}
	fmt.Printf("json to struct:%v\n", men)
}

// struct和map相互转换
func MapToStruct() {
	mapInstance := make(map[string]interface{})
	mapInstance["Name"] = "jqw"
	mapInstance["Age"] = 18

	var men Men
	err := mapstructure.Decode(mapInstance, &men)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("map to struct:%v\n", men)
}

func HandleJson() {
	// go语言中使用json
	// https://blog.csdn.net/tiaotiaoyly/article/details/38942311

	// 此处只是演示,实际处理中不建议把结构体存放在函数内部
	fmt.Printf("\n-------handlejson()-------\n")
	type Movie struct {
		Title  string
		Year   int  `json:"released"`        // 对应json中的released字段
		Color  bool `json:"color,omitempty"` // omitempty标识这个字段的值如果是空,这个字段就不显示在json体中了
		Actors []string
		Time   int64 `json:"-"` // 直接忽略字段
	}
	// 结构体成员标签定义tag

	var movies = []Movie{
		{Title: "aaa", Year: 1942, Color: false,
			Actors: []string{"abc", "bcd"}},
		{Title: "bbb", Year: 1942, Color: true,
			Actors: []string{"abc", "bcd"}},
		{Title: "ccc", Year: 1942, Color: true,
			Actors: []string{"abc", "bcd"}},
	}
	//data,err:=json.Marshal(movies)
	data, err := json.MarshalIndent(movies, "", "  ") // MarshalIndent:格式化显示数组
	if err != nil {
		log.Fatal("handle json err", err)
	}
	fmt.Printf("%s\n", data)
}
