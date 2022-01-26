package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

func LoggerMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// PUT或者POST方法时打印body
		if c.Request.Method == http.MethodPut || c.Request.Method == http.MethodPost {
			data, err := ioutil.ReadAll(c.Request.Body)
			if err != nil {
				glog.Errorf("read request body failed,err = %s.", err)
				return
			}
			glog.Infof("request body = %s.", string(data))
		}
		data, err := c.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		glog.Infof("data: %v\n", string(data))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // // 重新赋值,关键点
		c.Next()
	}
}
