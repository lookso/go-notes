package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

func TestHandle(c *gin.Context) {
	bt, _ := json.Marshal(c.Request)
	c.JSON(200, gin.H{
		"message": bt,
	})
}
func main() {
	router := gin.Default()
	router.GET("/test", TestHandle)
	router.Run()
}
