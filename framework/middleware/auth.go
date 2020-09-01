package middleware

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/libs/jwt"
	"go-notes/framework/libs/structs"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			c.JSON(structs.Forbidden(""))
			c.Abort()
			return
		}
		claims, err := jwt.ParseToken(token)
		if err != nil {
			c.JSON(structs.UnAuthorized("Token is not valid", nil))
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
