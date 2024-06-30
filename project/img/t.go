package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype"

	// "codechina.csdn.net/diandianxiyu/freetype"
	imgtype "github.com/shamsher31/goimgtype"
)

func main() {
	var imgdiandianxiyu_geek string = `/Users/peanut/Desktop/sy/123.jpg`
	imgb, err := os.Open(imgdiandianxiyu_geek)
	if err != nil {
		panic(err)
	}
	defer imgb.Close()

	datatype, err2 := imgtype.Get(imgdiandianxiyu_geek)
	var imgtype string = ""
	if err2 != nil {
		imgtype = ""
	} else {
		// 根据文件类型执行响应的操作
		switch datatype {
		case `image/jpeg`:
			imgtype = "jpeg"
		case `image/png`:
			imgtype = "png"
		}
	}
	if imgtype == "" {
		panic("暂不支持文件类型")
	}

	var imgbinfo image.Image
	if imgtype == "jpeg" {
		imgbinfo, _ = jpeg.Decode(imgb)
	} else {
		imgbinfo, _ = png.Decode(imgb)
	}

	//读取水印图片
	watermark, err := os.Open("/Users/peanut/Desktop/sy/tpp.png")
	if err != nil {
		panic(err)
	}
	defer watermark.Close()
	imgwatermark, err := png.Decode(watermark)
	if err != nil {
		panic(err)
	}

	offset := image.Pt(imgbinfo.Bounds().Dx()-imgwatermark.Bounds().Dx()-19, imgbinfo.Bounds().Dy()-imgwatermark.Bounds().Dy()-19)
	b := imgbinfo.Bounds()
	m := image.NewNRGBA(b) //按原图生成新图

	//文字水印
	fontBytes, err := ioutil.ReadFile("Arial Unicode.ttf") //读取字体文件
	if err != nil {
		log.Println(err)
	}
	font, err := freetype.ParseFont(fontBytes)
	if err != nil {
		log.Println(err)
	}

	f := freetype.NewContext()
	f.SetDPI(72)      //设置DPI
	f.SetFont(font)   //设置字体
	f.SetFontSize(24) //设置字号
	f.SetClip(imgbinfo.Bounds())
	f.SetDst(m)
	f.SetSrc(image.NewUniform(color.RGBA{R: 255, G: 0, B: 0, A: 255})) //设置颜色

	//新图写入原图和背景图
	draw.Draw(m, b, imgbinfo, image.ZP, draw.Src)
	draw.Draw(m, imgwatermark.Bounds().Add(offset), imgwatermark, image.ZP, draw.Over)

	pt := freetype.Pt(10, 50)
	_, err = f.DrawString("blog.csdn.net/diandianxiyu_geek 邻座的小雨同学", pt)

	//输出图像
	imgw, _ := os.Create("new.jpg")
	jpeg.Encode(imgw, m, &jpeg.Options{100})

}
