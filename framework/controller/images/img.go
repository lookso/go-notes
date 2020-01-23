package images

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Imgs(c *gin.Context)  {
	c.JSON(http.StatusOK, struct {
		Name string
		Url string
	}{
		Name:"广告背景图",
		Url:"https://oimageb2.ydstatic.com/image?id=-6407019294546986424&product=adpublish&w=520&h=347",
	})
}
