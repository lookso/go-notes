package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Param struct {
	Name string `json:"name"`
}

func main() {
	var param = Param{Name: "jack"}
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.GET("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.GET("/get/json", func(c *gin.Context) {
		// 获取原始字节
		err := c.ShouldBind(&param)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(param.Name)
		c.String(200, "ok")
	})
	router.HEAD("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(200, "ok")
	})
	router.Run(":8080")
}
