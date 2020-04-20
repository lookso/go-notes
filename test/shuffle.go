package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type Config struct {
	Title string `json:"title"`
	Coin  string `json:"coin"`
}

func main() {
	title := string([]rune("ab世界之大,无奇不有")[:3])
	fmt.Println("title",title)
	
	jsonStr := `[{"title":"111","coin":"0"},{"title":"222","coin":"10"}]`
	var config []*Config
	if err := json.Unmarshal([]byte(jsonStr), &config); err == nil {
		for _, v := range config {
			fmt.Println(v.Coin)
		}
	} else {
		fmt.Println("bbb")
	}
	
	slice := []interface{}{"a", "b", "c", "d", "e", "f"}
	Shuffle(slice)
	fmt.Println(slice)
	//for _, v := range slice {
	//	fmt.Println(v.(string))
	//}
}

func Shuffle(slice []interface{}) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(6, func(i, j int) {
		
	})
	//for len(slice) > 0 {
	//	n := len(slice)
	//	randIndex := r.Intn(n)
	//
	//	slice[n-1], slice[randIndex] = slice[randIndex], slice[n-1]
	//	slice = slice[:n-1]
	//}
}