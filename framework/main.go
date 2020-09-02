package main

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/router"
)

func main() {
	engine := gin.New()
	router.All(engine)

	engine.Run(":8090")
}
