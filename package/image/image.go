package main

import (
	"bytes"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	//fmt.Println(user)
	path := user.HomeDir+"/test/file/123.png"
	fmt.Println(path)
	templateData, err := ioutil.ReadFile(path)
	fmt.Println("123 err", err)
	_,f,e:=image.Decode(bytes.NewReader(templateData))
	fmt.Println(f,e)
	_, err = png.Decode(bytes.NewReader(templateData))
	fmt.Println("456",err)

	fh, err := os.Open(path)

	if err != nil {
		fmt.Println("ERR 1: " + err.Error())
		os.Exit(1)
	}
	if fh==nil{
		fmt.Println(13123)
	}else{
		fmt.Println(66666)
	}

	_, formatName, err := image.Decode(fh)
	if err != nil {
		panic(err)
	}
	fmt.Println(formatName)

	_, format, err := image.DecodeConfig(fh)
	if (format != "jpeg" && format != "gif" && format != "png") || err != nil {
		fmt.Println("ERR 2: " + err.Error())
		os.Exit(1)
	}
	fmt.Println("format", format)
	_, err = png.Decode(fh)
	if format == "png" {
		_, err = png.Decode(fh)
	} else if format == "jpg" {
		_, err = jpeg.Decode(fh)
	} else if format == "gif" {
		_, err = gif.Decode(fh)
	} else if format == "jpeg" {
		_, err = jpeg.Decode(fh)
	} else {
		fmt.Println("ERR 3: " + err.Error())
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("ERR 4: " + err.Error())
		os.Exit(1)
	}

	fmt.Println("Read OK.")

	os.Exit(1)
}
