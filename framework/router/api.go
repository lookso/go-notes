package router

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/controller/images"
)

func api(e *gin.Engine)  {
	api := e.Group("/api")
	api.GET("/img", images.Imgs)
}
