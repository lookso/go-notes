package router

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/controller"
	"go-notes/framework/middleware"
)

func Router(router *gin.Engine) {
	user := router.Group("/api/v1/user")
	user.POST("/login", controller.Login)

	auth := router.Group("/api/v1/", middleware.JwtAuth())
	{
		auth.GET("/user/info", controller.UserInfo)
	}
}
