package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
}

func parseSvgDimension(dim string) (int, error) {
	// 移除可能存在的单位
	dim = strings.TrimRight(dim, "px")
	dim = strings.TrimRight(dim, "pt")
	dim = strings.TrimRight(dim, "pc")
	dim = strings.TrimRight(dim, "mm")
	dim = strings.TrimRight(dim, "cm")
	dim = strings.TrimRight(dim, "in")
	dim = strings.TrimRight(dim, "em")
	dim = strings.TrimRight(dim, "ex")
	dim = strings.TrimRight(dim, "%")

	// 将字符串转换为浮点数
	f, err := strconv.ParseFloat(dim, 64)
	if err != nil {
		return 0, err
	}

	// 将浮点数转换为整数
	return int(f), nil
}
func getSVGSize(svgContent []byte) (width, height int, err error) {
	var doc Svg
	err = xml.Unmarshal(svgContent, &doc)
	if err != nil {
		return 0, 0, err
	}

	width, err = parseSvgDimension(doc.Width)
	if err != nil {
		return 0, 0, err
	}

	height, err = parseSvgDimension(doc.Height)
	if err != nil {
		return 0, 0, err
	}

	return width, height, nil
}

func getSVGAspectRatio(svgFilePath string) (float64, error) {
	// svgContent, err := ioutil.ReadFile(svgFilePath)
	// if err != nil {
	// 	return 0, err
	// }
	svgContent, err := os.ReadFile(svgFilePath)
	if err != nil {
		return 0, err
	}

	width, height, err := getSVGSize(svgContent)
	if err != nil {
		return 0, err
	}
	aspectRatio := float64(width) / float64(height)
	return aspectRatio, nil
}

func main() {
	svgFilePath := "/Users/peanut/Desktop/1234.svg"
	aspectRatio, err := getSVGAspectRatio(svgFilePath)
	if err != nil {
		fmt.Println("Error getting SVG aspect ratio:", err)
		return
	}

	fmt.Printf("The aspect ratio of the SVG is: %f\n", aspectRatio)
}
