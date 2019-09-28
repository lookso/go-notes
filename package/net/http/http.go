package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

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
