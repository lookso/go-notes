package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://mini-server.ainirobot.com/mini_ifttt/v1/task/list?page_no=0&page_size=100&task_type=1&family_id=f78042c675b1c1bac40ec9ae9892cc52"
	method := "GET"

	client := &http.Client{
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Family-ID", "f78042c675b1c1bac40ec9ae9892cc52")
	req.Header.Add("X-Login-Token", "h6clnnlpk41qqceo9p9abedtps")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Device-Type", "3")
	req.Header.Add("Device-ID", "M01SCNA0100200159994")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))

	a := "hello"
	b := "world"
	c := "golang"
	fmt.Println(strings.Join([]string{a, b, c}, ","))
	
	s := " world"
	buf := bytes.NewBufferString("hello")
	fmt.Println(buf.String()) //hello
	buf.WriteString(s)        //将string写入到buf的尾部
	fmt.Println(buf.String()) //hello world
}
