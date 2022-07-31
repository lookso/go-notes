package main

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	zkOt "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zkHttp "github.com/openzipkin/zipkin-go/reporter/http"
	"log"
	"time"
)

// 第一步: 开一个全局变量
var zkTracer opentracing.Tracer

func main() {
	// 第二步: 初始化 tracer
	{
		reporter := zkHttp.NewReporter("http://localhost:9411/api/v2/spans")
		defer reporter.Close()
		endpoint, err := zipkin.NewEndpoint("main", "localhost:80")
		if err != nil {
			log.Fatalf("unable to create local endpoint: %+v\n", err)
		}
		nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
		if err != nil {
			log.Fatalf("unable to create tracer: %+v\n", err)
		}
		zkTracer = zkOt.Wrap(nativeTracer)
		opentracing.SetGlobalTracer(zkTracer)
	}

	r := gin.Default()
	// 第三步: 添加一个 middleWare, 为每一个请求添加span
	r.Use(func(c *gin.Context) {
		span := zkTracer.StartSpan(c.FullPath())
		defer span.Finish()
		c.Next()
	})
	r.GET("/user/info",
		func(c *gin.Context) {
			time.Sleep(500 * time.Millisecond)
			c.JSON(200, gin.H{"code": 200, "msg": "OK"})
		})
	r.GET("/app/detail",
		func(c *gin.Context) {
			time.Sleep(500 * time.Millisecond)
			c.JSON(200, gin.H{"code": 200, "msg": "OK"})
		})
	r.GET("/order/list",
		func(c *gin.Context) {
			time.Sleep(500 * time.Millisecond)
			c.JSON(200, gin.H{"code": 200, "msg": "OK"})
		})
	r.Run(":8086")
}
