package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

type StuInfoRequest struct {
	StuID string `json:"stu_id"`
}

var done = make(chan string, 20)
var wg sync.WaitGroup

func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("请求来了", r.URL)
	fmt.Println(w.Write([]byte("hello cmcm")))
}
func Demo(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Query().Get("name"))

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("read body err, %v\n", err)
		return
	}
	println("json:", string(body))

	var stu StuInfoRequest
	if err = json.Unmarshal(body, &stu); err != nil {
		fmt.Printf("Unmarshal err, %v\n", err)
		return
	}
	stuIDs := strings.Split(stu.StuID, ",")
	for _, v := range stuIDs {
		done <- v
	}
	//todo
	wg.Add(1)
	go func() {
		handle(done)
		defer wg.Done()
	}()
	wg.Wait()

	var build strings.Builder
	build.WriteString("this is demo")
	build.WriteString(stu.StuID)
	res := build.String()
	fmt.Println(w.Write([]byte(res)))
}
func Chan(w http.ResponseWriter, r *http.Request) {
	taskChan := make(chan int)
	consumer := func() {
		for task := range taskChan {
			_ = task
		}
	}
	producer := func() {
		for i := 0; i < 10; i++ {
			taskChan <- i
		}
		close(taskChan) // 此处如果没有关闭chan,将发生协程泄露
	}
	go consumer()
	go producer()
	fmt.Println(w.Write([]byte("goroutine-num:" + strconv.Itoa(runtime.NumGoroutine())))) // 打印当前协程数量
}

func handle(done chan string) {
	for {
		time.Sleep(time.Second * 2)
		select {
		case e := <-done:
			fmt.Printf("Get element from done: %s\n", e)
		default:
			fmt.Println("default")
			return
		}
	}
}
func main() {
	http.HandleFunc("/index", Index)
	http.HandleFunc("/demo", Demo)
	http.HandleFunc("/chan", Chan)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
