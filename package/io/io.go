package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var ErrNoProgress = errors.New("multiple Read calls return no data or error")
	fmt.Println(ErrNoProgress)
	fmt.Println(io.ErrShortBuffer)
	str := strings.NewReader("this is what")
	bb, err := ioutil.ReadAll(str)
	if err != nil {
		log.Println("read fail")
	}
	fmt.Println(string(bb))


	res,_:=http.Get("https://www.baidu.com")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	fileData,_:=ioutil.ReadFile("../math/math.go")
	fmt.Println(len(fileData))
	fmt.Println(string(fileData))

	if ioutil.WriteFile("./io-file.txt",[]byte("test-test"),os.FileMode(777));err==nil{
		fd,_:=ioutil.ReadFile("./io-file.txt")
		fmt.Println(string(fd))
	}

}
