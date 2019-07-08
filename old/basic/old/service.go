/*
@Time : 2018/12/14 9:22 AM
@Author : Tenlu
@File : service
@Software: GoLand
*/
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"

	"log"
	"net/http"
	"strconv"
)

//定义一个slice切片变量,复合声明
var palette = []color.Color{color.White, color.RGBA{0, 255, 68, 255}, color.RGBA{26, 0, 255, 255}}

//声明一组常量
const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8888", nil))
	fmt.Println("hello")
}

// http://127.0.0.1:8888/?cycles=5

//处理http请求
func handler(w http.ResponseWriter, r *http.Request) {
	//if语句中嵌套表达式
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		if k == "cycles" {
			//函数有多个返回值时，multiple-value strconv.Atoi() in single-value context ，不需要的参数要用_接收
			cycles, _ := strconv.ParseFloat(v[0], 64)
			//fmt.Fprintf(w,"%q:%d",k,cycles);
			//调用gif图形函数
			if cycles < 50 {
				lissajous(w, cycles)
			}
		} else {
			fmt.Fprintf(w, "URL PATH:%q <br/>", r.URL.Path)
			fmt.Fprintf(w, "%q:%q", k, v)
		}
	}
}

//gif图形函数
func lissajous(out http.ResponseWriter, cycles float64) {
	const (
		//cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors

}
