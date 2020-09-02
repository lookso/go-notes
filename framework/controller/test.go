package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-notes/framework/libs/structs"
)

func GetRawData(c *gin.Context) {
	rawData, err := c.GetRawData() // 获取原始数据流
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(structs.Data(rawData))
}
