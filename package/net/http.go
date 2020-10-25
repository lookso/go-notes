package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	log.Println("请求来了", r.URL)
	fmt.Println(w.Write([]byte("hello cmcm")))
}
func main() {
	http.HandleFunc("/index", index)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		fmt.Println(err)
	}
}
