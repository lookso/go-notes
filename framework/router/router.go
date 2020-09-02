package router

import "github.com/gin-gonic/gin"

// 也可以通过分别在小路由文件中定义结构体实现,但是不够优雅
func All(e *gin.Engine) {
	// 注册路由
	Api(e)
	// 测试路由
	Test(e)
}