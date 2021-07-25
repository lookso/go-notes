package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
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
		names := []string{"lena", "austin", "foo"}

		// 将会输出:   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
	router.POST("/post/json", func(c *gin.Context) {
		// 获取原始字节
		err := c.ShouldBind(&param)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(param.Name)
		c.JSON(http.StatusOK, param)
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
	router.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// The specific definition of protobuf is written in the testdata/protoexample file.
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// Note that data becomes binary data in the response
		// Will output protoexample.Test protobuf serialized data
		c.ProtoBuf(http.StatusOK, data)
	})

	router.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		// 访问 http://localhost:8000/jsonp?callback=call
		// 将会输出:   call({foo:"bar"})
		c.JSONP(http.StatusOK, data)
	})

	router.Run(":8000")
}
