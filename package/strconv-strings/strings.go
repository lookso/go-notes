package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type My struct {
	Name string
	Age  int
}

var my = My{Name: "jack", Age: 11}

type Your struct {
	Name string
	Age  int
}

var your = Your{Name: "jack", Age: 11}
var user = Your{Name: "jack", Age: 11}

func main() {

	//if my == your { // 不同结构体实例,无法比较值
	//	fmt.Println("true")
	//}
	if your == user {
		fmt.Println(1, "true")
	}
	if reflect.DeepEqual(my, your) {
		fmt.Println(2, "true")
	}
	yourBytes, _ := json.Marshal(&your)
	myBytes, _ := json.Marshal(&my)

	if bytes.EqualFold(yourBytes, myBytes) {
		fmt.Println(3, "true")
	}

	// 判断两个utf-8编码字符串（将unicode大写、小写、标题三种格式字符视为相同）,是否相同,不区分大小写
	if strings.EqualFold("Go", "go") {
		fmt.Println(4, "true")
	}
	// 判断s是否有前缀字符串prefix。
	var name = "whoisnb"
	if strings.HasPrefix(name, "wh") {
		fmt.Println(5, "true")
	}
	// 判断s是否有前缀字符串prefix。
	if strings.HasSuffix(name, "nb") {
		fmt.Println(6, "true")
	}
	// 字符串中是否包含子串
	if strings.Contains(name, "is") {
		fmt.Println(7, "true")
	}
	// 字符串中指定的重复的子串的个数
	fmt.Println(strings.Count("cheese", "e"))

	//子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	fmt.Println(strings.Index("baidu", "a"))
	var ch byte = 'a' // 只能一个字符
	fmt.Println(strings.IndexByte("baidu", ch))

}
