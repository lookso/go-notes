package router

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/controller"
)

func Test(r *gin.Engine)  {
	r.POST("/test/raw_data",controller.GetRawData)
}
