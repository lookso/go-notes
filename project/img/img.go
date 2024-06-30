package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

type User struct {
	Name string
}

func main() {
	var arr = make([]User, 0, 10)
	arr = append(arr, User{Name: "jack"})
	arr = append(arr, User{Name: "toms"})
	for k, _ := range arr {
		if k == 1 {
			arr[k].Name = "lucy"
		}
	}
	fmt.Println(arr)

	imagePath := "/Users/peanut/Desktop/sy/123.jpg"
	file, _ := os.Open(imagePath)
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	b := img.Bounds()
	width := b.Max.X
	height := b.Max.Y

	fmt.Println("width = ", width)
	fmt.Println("height = ", height)
}
