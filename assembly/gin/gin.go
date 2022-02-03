package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"log"
	"net/http"
)

type Param struct {
	Name string `json:"name"`
	Skip []int  `json:"skip"`
}
type response struct {
	Code    int         `json:"code"`
	Stat    int         `json:"stat"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func Success(v interface{}) interface{} {
	ret := response{Stat: 1, Code: 0, Message: "ok", Data: v}
	return ret
}

func main() {
	var param = Param{
		Skip: make([]int, 0),
	}
	router := gin.Default()
	router.Use(gin.Recovery())

	router.POST("/post", func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		c.String(http.StatusOK, "ok")
	})
	router.GET("/post", LoggerMiddleWare(), func(c *gin.Context) {
		// 获取原始字节
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		c.String(http.StatusOK, string(d)+":热编译123456")
	})
	router.GET("/get/json", func(c *gin.Context) {
		type User struct {
			IDs  []int  `json:"id"`
			Name string `json:"name"`
		}
		type Resopnse struct {
			Users []User `json:"users"`
		}
		//var r = Resopnse{} // { "ids": null }
		var r = Resopnse{
			Users: make([]User, 0), // { "ids": []] }
		}
		c.JSON(http.StatusOK, Success(r))
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
		c.String(http.StatusOK, string(d))
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
