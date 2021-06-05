package main

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"math/rand"
	"sort"
	"strings"
	"time"
)

var stuList = []int{13, 12, 22, 21, 44}

type Names struct {
	Name string
}

func main() {

	timestamp := cast.ToInt(time.Now().Unix())
	if 1622773399-timestamp > 30 {
		fmt.Println("任务执行时间超过当前时间30分钟不可执行")
	}

	fmt.Println("time.Now().Hour()", time.Now().Hour())

	cast.ToInt(time.Now().Unix())

	var m = make(map[int]Names)
	fmt.Println("m", m)
	if len(m) == 0 {
		fmt.Println("okok------------------")
	}
	var t []int
	for _, v := range []int{1, 2, 3, 5} {
		if v == 3 {
			t = append(t, v)
			continue
		}
		fmt.Println(v)
	}
	fmt.Println("t", t)

	str := `[{"name":"a"},{"name":"b"}]`
	var st = make([]Names, 0)
	err := json.Unmarshal([]byte(str), &st)
	fmt.Println(st, len(st))
	fmt.Println("err", err)
	fmt.Println("------------aaaa--------")
	fmt.Println(strings.Replace(fmt.Sprint(stuList), " ", ",", -1))
	fmt.Println("----bbb--")
	idStr := strings.Trim(strings.Replace(fmt.Sprint(stuList), " ", ",", -1), "[]")
	fmt.Println("idStr", idStr)

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(200))
	fmt.Println("-------------")

	var tag bool
	tag = true
	if !tag {
		fmt.Println("123")
	}
	fmt.Println(tag)

	var s = []string{"102691", "102861", "103392", "35479", "36618"}
	ss := cast.ToIntSlice(s)
	sort.Ints(ss)
	fmt.Println(ss)

	//var a = make(map[string]interface{}, 0)
	//a["a"] = 123
	//var aa = a["a"]
	//fmt.Println(aa.(int) - 12)
	//
	//fmt.Println(reflect.TypeOf(a["a"]))
	//time.Now().Unix()
	//str := "中国 牛逼"
	//if Sp(str) {
	//	fmt.Println("ok")
	//} else {
	//	fmt.Println("false")
	//}
}
