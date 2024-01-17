package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func main() {
	fmt.Printf("%f", float32(3/2))

	s := `<div class="question-item-container">我草<img class="question-item-pic" src="https://static-data.tipaipai.com/maliang/jf/label/question/202207/22/dfJdtnRSVl6u84LrhZ1v.jpg?x-oss-process=image/crop,x_231,y_5279,w_1764,h_479"><p class="ocr_text_invisible">3.一个等腰三角形的周长是18cm，腰长是6.5cm，底边上的高是6cm，这个三角形的面积是( $$)cm^{2}。$$ }。</p></div>`
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(s))
	text:=doc.Text()
	fmt.Println(text)
}
