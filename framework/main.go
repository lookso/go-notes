package main

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/router"
	"net/http"
)

func main() {
	engine := gin.New()
	router.All(engine)

	http.ListenAndServe(":6666", nil)
}
