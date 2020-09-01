package main

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/router"
)

func main() {
	engine := gin.New()
	router.Router(engine)
	engine.Run(":8090")
}
