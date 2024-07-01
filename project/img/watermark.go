package main

import (
	"image"
	"image/png"
	"os"

	"golang.org/x/image/draw"
)

// 将水印图片铺满整个图片
func overlayWatermark(img image.Image, watermark image.Image, transparency float64) *image.RGBA {
	bounds := img.Bounds()
	newImage := image.NewRGBA(bounds)
	draw.Draw(newImage, bounds, img, image.ZP, draw.Src)

	// 计算水印图片的宽度和高度
	wWidth := watermark.Bounds().Dx()
	wHeight := watermark.Bounds().Dy()

	// 计算水印在原始图片中的起始位置
	startX := 0
	startY := 0

	// 铺满整个图片的循环
	for y := startY; y < bounds.Dy(); y += wHeight {
		if y+wHeight > bounds.Dy() {
			wHeight = bounds.Dy() - y
		}
		for x := startX; x < bounds.Dx(); x += wWidth {
			if x+wWidth > bounds.Dx() {
				wWidth = bounds.Dx() - x
			}
			// 将裁剪后的水印图片画到新图像上
			draw.Draw(
				newImage,
				image.Rect(x, y, x+wWidth, y+wHeight),
				watermark,
				image.Pt(0, 0),
				draw.Over,
			)
		}
	}

	return newImage
}

func main() {
	// 打开原始图片
	file, err := os.Open("/Users/peanut/Desktop/sy/123.jpg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	// 打开水印图片
	watermarkFile, err := os.Open("/Users/peanut/Desktop/sy/tpp.png")
	if err != nil {
		panic(err)
	}
	defer watermarkFile.Close()

	watermark, _, err := image.Decode(watermarkFile)
	if err != nil {
		panic(err)
	}

	// 设置水印的透明度
	transparency := 0.1

	// 生成铺满水印的新图像
	newImage := overlayWatermark(img, watermark, transparency)

	// 创建并保存新的图片文件
	outputFile, err := os.Create("watermarked_image.png")
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// 将新图像以PNG格式保存
	png.Encode(outputFile, newImage)
}
