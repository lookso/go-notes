package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cast"
)

type JsonData struct {
	LabelContainer struct {
		Chapters []int `json:"chapters"`
	} `json:"label_container"`
}

func main() {
	// 读取JSON文件
	// file, err := ioutil.ReadFile("/Users/peanut/apps/works/go/src/go-notes/script/label.json")
	// if err != nil {
	// 	fmt.Println("读取文件错误:", err)
	// 	return
	// }

	file, _ := os.Open("/Users/peanut/apps/works/go/src/go-notes/script/label.json")
	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	// 解析JSON数据
	var jsonData []JsonData
	err := json.Unmarshal(byteValue, &jsonData)
	if err != nil {
		fmt.Println("解析JSON错误:", err)
		return
	}
	for _, item := range jsonData {
		fmt.Println(item.LabelContainer.Chapters)
	}

	//将chapters数据写入文件
	outputFile, err := os.Create("53382.txt")
	if err != nil {
		fmt.Println("创建文件错误:", err)
		return
	}
	defer outputFile.Close()

	for _, page := range jsonData {
		// 将每个页面的pageId和chapters数据一行一个地写入文件
		for _, chapter := range page.LabelContainer.Chapters {
			outputFile.WriteString(cast.ToString(chapter) + ",")
		}
	}

	fmt.Println("数据已成功写入id.txt文件")
}
