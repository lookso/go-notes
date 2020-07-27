package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err!=nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.GET("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err!=nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.HEAD("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err!=nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.Run(":8080")
}
