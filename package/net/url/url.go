package url

import (
	"fmt"
	"log"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Add("p", "1")
	v.Add("drive", "ios")
	params := v.Encode()
	fmt.Println(params)
	//// url decode
	m, _ := url.ParseQuery(params)
	fmt.Println(m)

	var address = "http://www.baidu.com?name=jack+" + params
	encodeUrl := url.QueryEscape(address)
	fmt.Println(encodeUrl)
	var uncodeUrl, err = url.QueryUnescape(encodeUrl)
	if err != nil {
		errStruct := url.Error{Op: "add", URL: uncodeUrl, Err: err}
		fmt.Println(errStruct.Err)
	}
	fmt.Println(uncodeUrl)

	//drive=ios&p=1
	//map[drive:[ios] p:[1]]
	//http%3A%2F%2Fwww.baidu.com%3Fname%3Djack%2Bdrive%3Dios%26p%3D1
	//http://www.baidu.com?name=jack+drive=ios&p=1

	u, err := url.Parse("bing.com/search?q=dotnet") // 可以不加协议头https或者http
	fmt.Println(u.Scheme) // 将rawurl为一个URL结构体
	if err != nil {
		log.Fatal(err)
	}
	u.Scheme = "https"
	u.Host = "google.com"
	q := u.Query()
	q.Set("q", "golang")
	u.RawQuery = q.Encode()
	fmt.Println(u)

	urlStruct:=url.URL{Scheme:u.Scheme,Host:u.Host}
	fmt.Println(urlStruct.Host)


	// 本函数会假设rawurl是在一个HTTP请求里，因此会假设该参数是一个绝对URL或者绝对路径，并会假设该URL没有#fragment后
	uu,err:=url.ParseRequestURI("https://www.bing.com/search?q=dotnet") // 必须带上协议头
	if err!=nil{
		panic(err)
	}
	fmt.Println(uu.RequestURI())
	fmt.Println(uu.Query().Get("q"))
	p:=url.Values{}
	p.Add("name","jack")
	fmt.Println(p.Encode())

}
