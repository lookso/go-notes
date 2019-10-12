package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Golang开启http服务的三种方式
// https://www.jianshu.com/p/fe502c586034

func SayHello(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.URL.RawQuery)                // 获取?后面的参数, name=2342343&jack=123&age=123
	var vv, _ = url.ParseQuery(req.URL.RawQuery) // 参数解析
	fmt.Println("age", vv["age"])

	req.ParseForm()
	// ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段。
	fmt.Println("name", req.Form["name"])

	w.Write([]byte("Hello"))
	w.Write([]byte("nb"))
	http.NotFound(w, req) // 路由不存在
}

type httpServer struct {
}

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(r.URL.RequestURI()))
}

func main() {
	// StripPrefix返回一个处理器，该处理器会将请求的URL.Path字段中给定前缀prefix去除后再交由h处理。
	// StripPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found。
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
	time.Sleep(time.Second*20)
	getGet()

	var server httpServer
	http.Handle("/hi", server)
	http.HandleFunc("/hello", SayHello)
	http.ListenAndServe(":9990", nil)
}

func getGet() {
	resp, err := http.Get("https://m2.qiushibaike.com/article/list/text?count=2&page=1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
