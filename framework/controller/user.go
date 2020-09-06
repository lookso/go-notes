package controller

import (
	"github.com/gin-gonic/gin"
	"go-notes/framework/libs/jwt"
	"go-notes/framework/libs/structs"
	"go-notes/framework/resource"
	"go-notes/framework/utils"
)

var (
	uuid     string
	username = "peanut"
	password = "123456"
)

func Login(c *gin.Context) {
	var urq resource.LoginRequest
	if err := c.ShouldBind(&urq); err != nil {
		c.JSON(structs.BadRequest("参数错误"))
		return
	}
	uuid = utils.GetUuid()
	tokenString, err := jwt.GenerateToken(uuid, urq.UserName, urq.Password)
	if err != nil {
		c.JSON(structs.BadRequest("鉴权失败"))
		return
	}
	// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTg4NjU0NjEsImlhdCI6MTU5ODg2MTg2MX0.qbsLiy9Z_k3J61kuHyxBZYTCY_ZuD2rVmG6zUsY4FBI
	c.JSON(structs.Data(tokenString))
}

func UserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*jwt.MyClaims)
	if claims.UserName == "" || claims.Uid == "" {
		c.JSON(structs.BadRequest(""))
		return
	}
	c.JSON(structs.Data(resource.UserInfoResponse{
		Id:   claims.Uid,
		Name: claims.UserName,
		Sex:  1,
		Age:  28,
	}))
}
// 更新token
func Refresh(c *gin.Context) {
	c.JSON(structs.Data(nil))
}