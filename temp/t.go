package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Info struct {
	Id   int
	Name string
}

func main() {
	// file, err := os.Open("/Users/peanut/Desktop/bce86a31-3ef9-491c-ac9f-83f6f042ba57.apk")
	// if err != nil {
	// 	fmt.Println("Error opening file:", err)
	// 	return
	// }
	// defer file.Close()

	// body, err := io.ReadAll(file)
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// md5 := md5.New()
	// md5.Write(body)
	// md5Str := hex.EncodeToString(md5.Sum(nil))
	// fmt.Println(md5Str)

	arr := make([]Info, 0, 10)
	arr = append(arr, Info{Id: 1, Name: "jack"})
	arr = append(arr, Info{Id: 2, Name: "kan"})
	arr = append(arr, Info{Id: 3, Name: "toms"})
	// rand.Seed(time.Now().UnixNano())
	seed := time.Now().UnixNano()
	src := rand.NewSource(seed)

	// 创建一个新的随机生成器
	r := rand.New(src)
	r.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("arr", arr)
}
