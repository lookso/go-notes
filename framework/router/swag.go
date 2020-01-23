// +build doc

package router

import (
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func init() {
	swagHandler = ginSwagger.WrapHandler(swaggerFiles.Handler)
}
