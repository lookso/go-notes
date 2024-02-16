package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func main() {
	//resp, err := http.Get("https://www.tiktokglobalshopping.com/cn/markets/prices/bitcoin-btc")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer resp.Body.Close()
	//
	//if resp.StatusCode != 200 {
	//	log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	//}
	//
	//doc, err := goquery.NewDocumentFromReader(resp.Body)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//price := doc.Find(".main-container .price").First().Text()
	//fmt.Printf("BTC价格为： %s\n", price)


	// 加载页面内容
	resp, err := http.Get("https://www.tiktokglobalshopping.com/cn/markets/prices/bitcoin-btc")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}
	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	// 选择器选取信息
	priceElement := doc.Find(".price .price-last").First()

	// 提取信息
	price := priceElement.Text()

	// 打印
	fmt.Println("The current price of BTC is:", price)


}